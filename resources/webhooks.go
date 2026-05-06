package resources

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"crypto/subtle"
	"encoding/base64"
	"errors"
	"os"
	"strings"

	"github.com/rewritetoday/golang/api"
)

// WebhookManager provides webhook resource operations.
type WebhookManager struct {
	Base
}

// Webhooks is kept as a compatibility alias for WebhookManager.
type Webhooks = WebhookManager

type VerifyWebhookOptions struct {
	Secret  string
	Payload string
	Headers map[string]string
}

func (r *WebhookManager) Verify(options VerifyWebhookOptions) (bool, error) {
	secret := options.Secret
	if secret == "" {
		secret = os.Getenv("REWRITE_WEBHOOK_SECRET")
	}
	if secret == "" {
		return false, errors.New("Rewrite could not find a webhook secret to verify.")
	}

	value := secret
	if trimmed, ok := strings.CutPrefix(secret, "whsec_"); ok {
		value = trimmed
	}
	_, hasRWSecretPrefix := strings.CutPrefix(secret, "rw_whsec_")

	var key []byte
	switch {
	case hasRWSecretPrefix:
		key = []byte(secret)
	default:
		decoded, err := base64.StdEncoding.DecodeString(value)
		if err == nil {
			key = decoded
		}
	}

	svixID := headerValue(options.Headers, "svix-id")
	svixTimestamp := headerValue(options.Headers, "svix-timestamp")
	signature := webhookSignature(headerValue(options.Headers, "svix-signature"))

	if len(key) == 0 || svixID == "" || svixTimestamp == "" || signature == "" {
		return false, nil
	}

	mac := hmac.New(sha256.New, key)
	mac.Write([]byte(svixID))
	mac.Write([]byte("."))
	mac.Write([]byte(svixTimestamp))
	mac.Write([]byte("."))
	mac.Write([]byte(options.Payload))
	expected := base64.StdEncoding.EncodeToString(mac.Sum(nil))

	if len(signature) != len(expected) {
		return false, nil
	}

	return subtle.ConstantTimeCompare([]byte(signature), []byte(expected)) == 1, nil
}

func (r *WebhookManager) Create(ctx context.Context, body api.RESTPostCreateWebhookBody) (api.RESTPostCreateWebhookData, error) {
	var out api.RESTPostCreateWebhookData
	err := r.Rest.Post(ctx, api.Routes.Webhooks.Create(), body, &out, nil)
	return out, err
}

func (r *WebhookManager) Update(ctx context.Context, id string, body api.RESTPatchUpdateWebhookBody) (api.RESTPatchUpdateWebhookData, error) {
	var out api.RESTPatchUpdateWebhookData
	err := r.Rest.Patch(ctx, api.Routes.Webhooks.Update(id), body, &out, nil)
	return out, err
}

func (r *WebhookManager) Delete(ctx context.Context, id string) (api.RESTDeleteWebhookData, error) {
	var out api.RESTDeleteWebhookData
	err := r.Rest.Delete(ctx, api.Routes.Webhooks.Delete(id), &out, nil)
	return out, err
}

func (r *WebhookManager) Sweep(ctx context.Context, body api.RESTDeleteWebhooksBody) (api.RESTDeleteWebhooksData, error) {
	var out api.RESTDeleteWebhooksData
	err := r.Rest.DeleteWithBody(ctx, api.Routes.Webhooks.Sweep(), body, &out, nil)
	return out, err
}

func (r *WebhookManager) List(ctx context.Context, query *api.RESTGetListWebhooksQueryParams) (api.RESTGetListWebhooksData, error) {
	var out api.RESTGetListWebhooksData
	err := r.Rest.Get(ctx, api.Routes.Webhooks.List(query), &out, nil)
	return out, err
}

func (r *WebhookManager) Get(ctx context.Context, id string) (api.RESTGetWebhookData, error) {
	var out api.RESTGetWebhookData
	err := r.Rest.Get(ctx, api.Routes.Webhooks.Get(id), &out, nil)
	return out, err
}

func headerValue(headers map[string]string, key string) string {
	for name, value := range headers {
		if strings.EqualFold(name, key) {
			return value
		}
	}
	return ""
}

func webhookSignature(value string) string {
	if value == "" {
		return ""
	}

	parts := strings.Split(value, ",")
	if len(parts) < 2 {
		return ""
	}

	return parts[1]
}

package rewrite

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestNewFromOptionsAndTemplatesList(t *testing.T) {
	var authHeader string
	var requestPath string
	var requestQuery string

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader = r.Header.Get("Authorization")
		requestPath = r.URL.Path
		requestQuery = r.URL.RawQuery
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"ok":true,"data":[],"cursor":{"persist":false}}`))
	}))
	defer server.Close()

	client, err := New(RewriteOptions{
		Secret: "rw_test",
		Rest: &RESTOptions{
			BaseURL: server.URL,
			Retry: &RetryOptions{
				Max:   1,
				Delay: func(int) time.Duration { return 0 },
			},
		},
	})
	if err != nil {
		t.Fatalf("unexpected constructor error: %v", err)
	}

	templates, err := client.Templates.List(context.Background(), nil)
	if err != nil {
		t.Fatalf("unexpected list error: %v", err)
	}

	if authHeader != "Bearer rw_test" {
		t.Fatalf("unexpected auth header: %q", authHeader)
	}
	if requestPath != "/v1/templates" {
		t.Fatalf("unexpected path: %q", requestPath)
	}
	if requestQuery != "limit=15" {
		t.Fatalf("unexpected query: %q", requestQuery)
	}
	if !templates.OK || len(templates.Data) != 0 {
		t.Fatalf("unexpected payload: %+v", templates)
	}
	if templates.Cursor == nil || templates.Cursor.Persist {
		t.Fatalf("unexpected cursor: %+v", templates.Cursor)
	}
}

func TestConstructorSecretTypeError(t *testing.T) {
	_, err := New(123)
	if err == nil || err.Error() != "Expected a string for the secret" {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestRoutesMatchPublicAPI(t *testing.T) {
	if route := Routes.Templates.List(nil); route != "/templates?limit=15" {
		t.Fatalf("unexpected templates list route: %s", route)
	}
	with18n := true
	if route := Routes.Templates.List(&RESTGetListTemplatesQueryParams{With18n: &with18n}); route != "/templates?limit=15&with18n=true" {
		t.Fatalf("unexpected templates list route with with18n: %s", route)
	}
	if route := Routes.Messages.List(nil); route != "/messages?limit=15" {
		t.Fatalf("unexpected messages list route: %s", route)
	}
	if route := Routes.Messages.Send(); route != "/messages" {
		t.Fatalf("unexpected message send route: %s", route)
	}
	if route := Routes.Messages.Get("abc"); route != "/messages/:abc" {
		t.Fatalf("unexpected message get route: %s", route)
	}
	if route := Routes.Messages.Cancel("abc"); route != "/messages/abc/cancel" {
		t.Fatalf("unexpected message cancel route: %s", route)
	}
	if route := Routes.Webhooks.Logs("abc", nil); route != "/webhooks/abc/logs?limit=15" {
		t.Fatalf("unexpected webhook logs route: %s", route)
	}
	if route := Routes.Webhooks.Get("abc"); route != "/webhooks/abc" {
		t.Fatalf("unexpected webhook route: %s", route)
	}
}

func TestTemplatesCreateDoesNotSendHiddenFields(t *testing.T) {
	var payload map[string]any

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/v1/templates" {
			t.Fatalf("unexpected path: %s", r.URL.Path)
		}
		if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
			t.Fatalf("decode body: %v", err)
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"ok":true,"data":{"id":"1","createdAt":"2026-02-19T20:01:09.000Z"}}`))
	}))
	defer server.Close()

	client, err := New(RewriteOptions{Secret: "rw", Rest: &RESTOptions{BaseURL: server.URL}})
	if err != nil {
		t.Fatalf("unexpected constructor error: %v", err)
	}

	_, err = client.Templates.Create(context.Background(), RESTPostCreateTemplateBody{
		Name:        "welcome",
		Description: "Welcome message",
		Content:     "Hello {{name}}",
		I18N: map[CountryCode]string{
			"br": "Ola {{name}}",
		},
		Variables: []APITemplateVariable{
			{Name: "name", Fallback: "customer"},
		},
	})
	if err != nil {
		t.Fatalf("unexpected create error: %v", err)
	}

	if _, ok := payload["project"]; ok {
		t.Fatalf("did not expect project in body: %#v", payload["project"])
	}
	if payload["name"] != "welcome" {
		t.Fatalf("unexpected name in body: %#v", payload["name"])
	}
	if payload["description"] != "Welcome message" {
		t.Fatalf("unexpected description in body: %#v", payload["description"])
	}
}

func TestTemplatesListSupportsNodeAndLegacyWith18nFields(t *testing.T) {
	var requestQuery string

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestQuery = r.URL.RawQuery
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"ok":true,"data":[],"cursor":{"persist":false}}`))
	}))
	defer server.Close()

	client, err := New(RewriteOptions{Secret: "rw", Rest: &RESTOptions{BaseURL: server.URL}})
	if err != nil {
		t.Fatalf("unexpected constructor error: %v", err)
	}

	with18n := true
	if _, err := client.Templates.List(context.Background(), &RESTGetListTemplatesQueryParams{
		With18n: &with18n,
	}); err != nil {
		t.Fatalf("unexpected list error: %v", err)
	}
	if requestQuery != "limit=15&with18n=true" {
		t.Fatalf("unexpected node query: %q", requestQuery)
	}

	if _, err := client.Templates.List(context.Background(), &RESTGetListTemplatesQueryParams{
		WithI18N: &with18n,
	}); err != nil {
		t.Fatalf("unexpected legacy list error: %v", err)
	}
	if requestQuery != "limit=15&with18n=true" {
		t.Fatalf("unexpected legacy query: %q", requestQuery)
	}
}

func TestTemplatesUpdateUsesNodeContract(t *testing.T) {
	var requestPath string
	var payload map[string]any

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestPath = r.URL.Path
		if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
			t.Fatalf("decode body: %v", err)
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"ok":true,"data":{"id":"1","createdAt":"2026-02-19T20:01:09.000Z"}}`))
	}))
	defer server.Close()

	client, err := New(RewriteOptions{Secret: "rw", Rest: &RESTOptions{BaseURL: server.URL}})
	if err != nil {
		t.Fatalf("unexpected constructor error: %v", err)
	}

	response, err := client.Templates.Update(context.Background(), "tmpl_123", RESTPostCreateTemplateBody{
		Name:    "welcome",
		Content: "Hello {{name}}",
		Variables: []APITemplateVariable{
			{Name: "name", Fallback: "friend"},
		},
	})
	if err != nil {
		t.Fatalf("unexpected update error: %v", err)
	}

	if requestPath != "/v1/templates/tmpl_123" {
		t.Fatalf("unexpected path: %q", requestPath)
	}
	if payload["name"] != "welcome" {
		t.Fatalf("unexpected payload: %#v", payload)
	}
	if response.Data.ID != "1" {
		t.Fatalf("unexpected response: %+v", response)
	}
}

func TestMessagesSendSendsIdempotencyKey(t *testing.T) {
	var idempotencyKey string
	var requestPath string

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		idempotencyKey = r.Header.Get("Idempotency-Key")
		requestPath = r.URL.Path
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"ok":true,"data":{"id":"1","createdAt":"2026-02-19T20:01:09.000Z","analysis":{"characters":5,"encoding":"gsm7","segments":{"concat":153,"count":1,"reason":"fits","single":160}}}}`))
	}))
	defer server.Close()

	client, err := New(RewriteOptions{Secret: "rw", Rest: &RESTOptions{BaseURL: server.URL}})
	if err != nil {
		t.Fatalf("unexpected constructor error: %v", err)
	}

	response, err := client.Messages.Send(context.Background(), SendMessageOptions{
		IdempotencyKey: "msg-1",
		RESTPostSendMessageBody: RESTPostSendMessageBody{
			To:      "+5511999999999",
			Content: "hello",
		},
	})
	if err != nil {
		t.Fatalf("unexpected send error: %v", err)
	}

	if requestPath != "/v1/messages" {
		t.Fatalf("unexpected path: %q", requestPath)
	}
	if idempotencyKey != "msg-1" {
		t.Fatalf("unexpected idempotency key: %q", idempotencyKey)
	}
	if response.Data.ID != "1" {
		t.Fatalf("unexpected response: %+v", response)
	}
}

func TestMessagesCancelUsesMessageIDInRoute(t *testing.T) {
	var requestPath string
	var requestBody string

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestPath = r.URL.Path
		body, _ := io.ReadAll(r.Body)
		requestBody = string(body)
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"ok":true,"data":null}`))
	}))
	defer server.Close()

	client, err := New(RewriteOptions{Secret: "rw", Rest: &RESTOptions{BaseURL: server.URL}})
	if err != nil {
		t.Fatalf("unexpected constructor error: %v", err)
	}

	if _, err := client.Messages.Cancel(context.Background(), "msg_123"); err != nil {
		t.Fatalf("unexpected cancel error: %v", err)
	}

	if requestPath != "/v1/messages/msg_123/cancel" {
		t.Fatalf("unexpected path: %q", requestPath)
	}
	if requestBody != "" {
		t.Fatalf("unexpected body: %q", requestBody)
	}
}

func TestMessagesGetUsesNodeRouteBuilder(t *testing.T) {
	var requestPath string

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestPath = r.URL.Path
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"ok":true,"data":{"id":"1","createdAt":"2026-02-19T20:01:09.000Z","to":"+5511999999999","type":"SMS","tags":[],"status":"SENT","country":"br","content":"hello","encoding":"GMS7","isPayAsYouGo":false}}`))
	}))
	defer server.Close()

	client, err := New(RewriteOptions{Secret: "rw", Rest: &RESTOptions{BaseURL: server.URL}})
	if err != nil {
		t.Fatalf("unexpected constructor error: %v", err)
	}

	if _, err := client.Messages.Get(context.Background(), "msg_123"); err != nil {
		t.Fatalf("unexpected get error: %v", err)
	}

	if requestPath != "/v1/messages/:msg_123" {
		t.Fatalf("unexpected path: %q", requestPath)
	}
}

func TestOTPVerifyUsesMessageIDInPath(t *testing.T) {
	var requestPath string

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestPath = r.URL.Path
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"ok":true,"data":{"id":"1","valid":true,"verifiedAt":"2026-02-19T20:01:09.000Z"}}`))
	}))
	defer server.Close()

	client, err := New(RewriteOptions{Secret: "rw", Rest: &RESTOptions{BaseURL: server.URL}})
	if err != nil {
		t.Fatalf("unexpected constructor error: %v", err)
	}

	_, err = client.OTP.Verify(context.Background(), VerifyOTPOptions{
		ID: "1",
		RESTPostVerifyOTPCodeBody: RESTPostVerifyOTPCodeBody{
			To:   "+5511999999999",
			Code: "123456",
		},
	})
	if err != nil {
		t.Fatalf("unexpected verify error: %v", err)
	}

	if requestPath != "/v1/otp/1/verify" {
		t.Fatalf("unexpected path: %q", requestPath)
	}
}

func TestAPIKeysDeleteReturnsHTTPError(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(`{"ok":false,"error":{"code":"INVALID","message":"invalid api key id"}}`))
	}))
	defer server.Close()

	client, err := New(RewriteOptions{Secret: "rw", Rest: &RESTOptions{BaseURL: server.URL}})
	if err != nil {
		t.Fatalf("unexpected constructor error: %v", err)
	}

	_, err = client.APIKeys.Delete(context.Background(), "abc")
	if err == nil {
		t.Fatal("expected error")
	}

	var httpErr *HTTPError
	if !errors.As(err, &httpErr) {
		t.Fatalf("expected HTTPError, got %T", err)
	}
	if httpErr.Status != http.StatusBadRequest {
		t.Fatalf("unexpected status: %d", httpErr.Status)
	}
	if httpErr.Message != "invalid api key id" {
		t.Fatalf("unexpected message: %s", httpErr.Message)
	}
}

func TestWebhookVerifyReturnsTrueForValidSignature(t *testing.T) {
	client, err := New("rw_test")
	if err != nil {
		t.Fatalf("unexpected constructor error: %v", err)
	}

	payload := `{"type":"message.sent"}`
	mac := hmac.New(sha256.New, []byte("secret"))
	mac.Write([]byte("msg_test.1710000000." + payload))
	signature := base64.StdEncoding.EncodeToString(mac.Sum(nil))

	valid, err := client.Webhooks.Verify(VerifyWebhookOptions{
		Secret:  "whsec_c2VjcmV0",
		Payload: payload,
		Headers: map[string]string{
			"svix-id":        "msg_test",
			"svix-timestamp": "1710000000",
			"svix-signature": "v1," + signature,
		},
	})
	if err != nil {
		t.Fatalf("unexpected verify error: %v", err)
	}
	if !valid {
		t.Fatal("expected valid signature")
	}
}

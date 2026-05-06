package rewrite

import (
	"errors"

	"github.com/rewritetoday/golang/resources"
	"github.com/rewritetoday/golang/rest"
)

// Client is the main SDK client for the Rewrite API.
type Client struct {
	// Rest exposes the low-level REST client.
	Rest *rest.Client

	// Secret is the resolved API secret used for authentication.
	secret string

	// APIKeys exposes public API key operations.
	APIKeys *resources.APIKeyManager

	// Contacts exposes contact operations.
	Contacts *resources.ContactManager

	// Deliveries exposes webhook delivery operations.
	Deliveries *resources.DeliveryManager

	// Logs exposes request log operations.
	Logs *resources.LogManager

	// Messages exposes message operations.
	Messages *resources.MessageManager

	// OTP exposes OTP operations.
	OTP *resources.OTPManager

	// Segments exposes segment operations.
	Segments *resources.SegmentManager

	// Tags exposes reusable tag operations.
	Tags *resources.TagManager

	// Templates exposes template operations.
	Templates *resources.TemplateManager

	// WebhookLogs is kept as a compatibility alias to Deliveries.
	WebhookLogs *resources.DeliveryManager

	// Webhooks exposes webhook operations.
	Webhooks *resources.WebhookManager
}

// Rewrite is an alias to Client for naming parity with the Node SDK.
type Rewrite = Client

// RewriteOptions configures New/NewRewrite.
type RewriteOptions struct {
	// Secret is the Rewrite API key.
	Secret string
	// Rest customizes the low-level REST client options.
	Rest *rest.Options
}

// New creates a new Rewrite client instance.
//
// Accepted options:
//   - string (API secret)
//   - RewriteOptions
//   - *RewriteOptions
func New(options any) (*Client, error) {
	resolved, err := resolveOptions(options)

	if err != nil {
		return nil, err
	}

	restOptions := rest.Options{}

	if resolved.Rest != nil {
		restOptions = *resolved.Rest
	}

	restOptions.Auth = resolved.Secret

	restClient, err := rest.New(restOptions)

	if err != nil {
		return nil, err
	}

	client := &Client{
		Rest:       restClient,
		secret:     resolved.Secret,
		APIKeys:    &resources.APIKeyManager{Base: resources.Base{Rest: restClient}},
		Contacts:   &resources.ContactManager{Base: resources.Base{Rest: restClient}},
		Deliveries: &resources.DeliveryManager{Base: resources.Base{Rest: restClient}},
		Logs:       &resources.LogManager{Base: resources.Base{Rest: restClient}},
		Messages:   &resources.MessageManager{Base: resources.Base{Rest: restClient}},
		OTP:        &resources.OTPManager{Base: resources.Base{Rest: restClient}},
		Segments:   &resources.SegmentManager{Base: resources.Base{Rest: restClient}},
		Tags:       &resources.TagManager{Base: resources.Base{Rest: restClient}},
		Templates:  &resources.TemplateManager{Base: resources.Base{Rest: restClient}},
		Webhooks:   &resources.WebhookManager{Base: resources.Base{Rest: restClient}},
	}
	client.WebhookLogs = client.Deliveries

	return client, nil
}

// NewRewrite is an alias for New.
func NewRewrite(options any) (*Rewrite, error) {
	return New(options)
}

func resolveOptions(options any) (RewriteOptions, error) {
	switch v := options.(type) {
	case string:
		return RewriteOptions{Secret: v}, nil
	case RewriteOptions:
		return v, nil
	case *RewriteOptions:
		if v == nil {
			return RewriteOptions{}, errors.New("Expected a string for the secret")
		}
		return *v, nil
	default:
		return RewriteOptions{}, errors.New("Expected a string for the secret")
	}
}

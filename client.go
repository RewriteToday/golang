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

	// APIKeys exposes API key operations.
	APIKeys *resources.APIKeys
	// Projects exposes project operations.
	Projects *resources.Projects
	// Templates exposes template operations.
	Templates *resources.Templates
	// Webhooks exposes webhook operations.
	Webhooks *resources.Webhooks
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
		Rest:      restClient,
		secret:    resolved.Secret,
		APIKeys:   &resources.APIKeys{Base: resources.Base{Rest: restClient}},
		Projects:  &resources.Projects{Base: resources.Base{Rest: restClient}},
		Templates: &resources.Templates{Base: resources.Base{Rest: restClient}},
		Webhooks:  &resources.Webhooks{Base: resources.Base{Rest: restClient}},
	}

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

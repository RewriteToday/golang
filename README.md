<div align="center">

# Rewrite Go SDK

[`github.com/rewritetoday/golang`](https://pkg.go.dev/github.com/rewritetoday/golang), the official Go SDK for the Rewrite API.

It wraps authentication, typed REST calls, and resource helpers on top of the SDK REST and API layers.

<img src="https://cdn.rewritetoday.com/assets/banners/go-sdk.png" width="100%" alt="Rewrite Banner"/>

## Installation

Use your preferred Go workflow:

</div>

```bash
go get github.com/rewritetoday/golang
```

<div align="center">

## Quick Start

</div>

```go
package main

import (
	"context"
	"fmt"
	"log"

	rewrite "github.com/rewritetoday/golang"
)

func main() {
	client, err := rewrite.New("rw_abc")

	if err != nil {
		log.Fatal(err)
	}

	hooks, err := client.Webhooks.List(
		context.Background(),
		"123456789012345678",
		&rewrite.RESTGetListWebhooksQueryParams{Limit: 10},
	)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%+v\n", hooks)
}
```

<div align="center">

## Create The Client

You can pass the API key directly or use the full options object.

</div>

```go
package main

import (
	"time"

	rewrite "github.com/rewritetoday/golang"
)

func buildClient() (*rewrite.Client, error) {
	return rewrite.New(rewrite.RewriteOptions{
		Secret: "rw_abc",
		Rest: &rewrite.RESTOptions{
			Timeout: 10 * time.Second,
			Headers: map[string]string{
				"x-trace-id": "my-service",
			},
			Retry: &rewrite.RetryOptions{
				Max: 3,
				Delay: func(attempt int) time.Duration {
					return time.Duration(attempt) * 250 * time.Millisecond
				},
			},
		},
	})
}
```

<div align="center">

### Templates

</div>

```go
projectId := "123456789012345678"

created, err := client.Templates.Create(context.Background(), rewrite.CreateTemplateOptions{
	Project: projectId,
	RESTPostCreateTemplateBody: rewrite.RESTPostCreateTemplateBody{
		Name:    "welcome_sms",
		Content: "Hi {{name}}, welcome to {{company}}.",
		Variables: []rewrite.APITemplateVariable{
			{Name: "name", Fallback: "customer"},
			{Name: "company", Fallback: "Rewrite"},
		},
	},
})

if err != nil {
	log.Fatal(err)
}

templates, err := client.Templates.List(context.Background(), projectId, &rewrite.RESTGetListTemplatesQueryParams{Limit: 20})

if err != nil {
	log.Fatal(err)
}

fmt.Printf("created=%+v templates=%+v\n", created, templates)
```

<div align="center">

### Webhooks

</div>

```go
projectId := "123456789012345678"

hook, err := client.Webhooks.Create(context.Background(), rewrite.CreateWebhookOptions{
	Project:  projectId,
	Name:     "delivery-events",
	Endpoint: "https://example.com/webhooks/rewrite",
	Events: []rewrite.WebhookEventType{
		rewrite.WebhookEventTypeSMSDelivered,
		rewrite.WebhookEventTypeSMSFailed,
	},
})

if err != nil {
	log.Fatal(err)
}

_, err = client.Webhooks.Update(context.Background(), string(hook.Data.ID), rewrite.UpdateWebhookOptions{
	Project: projectId,
	RESTPatchUpdateWebhookBody: rewrite.RESTPatchUpdateWebhookBody{
		Status: rewrite.WebhookStatusInactive,
	},
})

if err != nil {
	log.Fatal(err)
}

hooks, err := client.Webhooks.List(context.Background(), projectId, &rewrite.RESTGetListWebhooksQueryParams{Limit: 10})

if err != nil {
	log.Fatal(err)
}

fmt.Printf("%+v\n", hooks)
```

<div align="center">

### API Keys

</div>

```go
projectId := "123456789012345678"

key, err := client.APIKeys.Create(context.Background(), rewrite.CreateAPIKeyOptions{
	Project: projectId,
	RESTPostCreateAPIKeyBody: rewrite.RESTPostCreateAPIKeyBody{
		Name: "backend-prod",
		Scopes: []rewrite.APIKeyScope{
			rewrite.APIKeyScopeReadProject,
			rewrite.APIKeyScopeReadTemplates,
		},
	},
})

if err != nil {
	log.Fatal(err)
}

fmt.Printf("%+v\n", key)
```

<div align="center">

## Error Handling

Requests run through the SDK REST client. HTTP failures can return `HTTPError`.

</div>

```go
_, err := client.APIKeys.List(context.Background(), "invalid_id", nil)

if err != nil {
	var httpErr *rewrite.HTTPError

	if errors.As(err, &httpErr) {
		fmt.Println("HTTP Error:", httpErr.Status, httpErr.Method, httpErr.URL)
	}
}
```

<div align="center">

---

Made with 🤍 by the Rewrite team. <br/>
SMS the way it should be.

</div>

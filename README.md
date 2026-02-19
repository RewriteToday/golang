<div align="center">

# Rewrite Go SDK

[`github.com/rewritetoday/golang`](https://pkg.go.dev/github.com/rewritetoday/golang), the official Go SDK for the Rewrite API.

It wraps authentication, typed REST calls, and resource helpers on top of the SDK REST and API layers.

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
	rewriteClient, err := rewrite.New("rw_live_xxx")
	
	if err != nil {
		log.Fatal(err)
	}

	project, err := rewriteClient.Projects.Get(context.Background(), "123456789012345678")
	
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%+v\n", project)
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
		Secret: "rw_live_xxx",
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

### Projects

</div>

```go
project, err := rewriteClient.Projects.Create(context.Background(), rewrite.RESTPostCreateProjectBody{
	Name: "AbacatePay Notifications",
})

if err != nil {
	log.Fatal(err)
}

fmt.Printf("%+v\n", project)
```

<div align="center">

### Templates

</div>

```go
projectID := "123456789012345678"
customerFallback := "customer"
companyFallback := "Rewrite"

created, err := rewriteClient.Templates.Create(context.Background(), rewrite.CreateTemplateOptions{
	Project: projectID,
	RESTPostCreateTemplateBody: rewrite.RESTPostCreateTemplateBody{
		Name: "welcome_sms",
		Variables: []rewrite.APITemplateVariable{
			{Name: "name", Fallback: &customerFallback},
			{Name: "company", Fallback: &companyFallback},
		},
	},
})

if err != nil {
	log.Fatal(err)
}

templates, err := rewriteClient.Templates.List(context.Background(), projectID, &rewrite.RESTGetListTemplatesQueryParams{Limit: 20})

if err != nil {
	log.Fatal(err)
}

fmt.Printf("created=%+v templates=%+v\n", created, templates)
```

<div align="center">

### Webhooks

</div>

```go
projectID := "123456789012345678"

hook, err := rewriteClient.Webhooks.Create(context.Background(), rewrite.CreateWebhookOptions{
	Project:  projectID,
	Name:     "delivery-events",
	Endpoint: "https://example.com/rewrite/webhooks",
	Events: []rewrite.WebhookEventType{
		rewrite.WebhookEventTypeSMSDelivered,
		rewrite.WebhookEventTypeSMSFailed,
	},
})

if err != nil {
	log.Fatal(err)
}

_, err = rewriteClient.Webhooks.Update(context.Background(), string(hook.Data.ID), rewrite.UpdateWebhookOptions{
	Project: projectID,
	RESTPatchUpdateWebhookBody: rewrite.RESTPatchUpdateWebhookBody{
		Status: rewrite.WebhookStatusInactive,
	},
})

if err != nil {
	log.Fatal(err)
}

hooks, err := rewriteClient.Webhooks.List(context.Background(), projectID, &rewrite.RESTGetListWebhooksQueryParams{Limit: 10})

if err != nil {
	log.Fatal(err)
}

fmt.Printf("%+v\n", hooks)
```

<div align="center">

### API Keys

</div>

```go
projectID := "123456789012345678"

key, err := rewriteClient.APIKeys.Create(context.Background(), rewrite.CreateAPIKeyOptions{
	Project: projectID,
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

Requests run through the SDK REST client. HTTP failures can throw `HTTPError`.

</div>

```go
_, err := rewriteClient.Projects.Get(context.Background(), "invalid_id")
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

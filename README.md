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
created, err := client.Templates.Create(context.Background(), rewrite.RESTPostCreateTemplateBody{
	Name:        "welcome_sms",
	Description: rewrite.NewNullableString("Welcome copy"),
	Content:     "Hi {{name}}, welcome to {{company}}.",
	Variables: []rewrite.APITemplateVariable{
		{Name: "name", Fallback: "customer"},
		{Name: "company", Fallback: "Rewrite"},
	},
	Tags: []rewrite.APITemplateTag{
		{Name: "category", Value: "welcome"},
	},
})

if err != nil {
	log.Fatal(err)
}

templates, err := client.Templates.List(context.Background(), &rewrite.RESTGetListTemplatesQueryParams{Limit: 20})

if err != nil {
	log.Fatal(err)
}

fmt.Printf("created=%+v templates=%+v\n", created, templates)
```

<div align="center">

### Webhooks

</div>

```go
hook, err := client.Webhooks.Create(context.Background(), rewrite.RESTPostCreateWebhookBody{
	Name:     "delivery-events",
	Endpoint: "https://example.com/webhooks/rewrite",
	Events: []rewrite.WebhookEventSelection{
		rewrite.WebhookEventTypeMessageDelivered,
		rewrite.WebhookEventTypeMessageFailed,
	},
})

if err != nil {
	log.Fatal(err)
}

_, err = client.Webhooks.Update(context.Background(), string(hook.Data.ID), rewrite.RESTPatchUpdateWebhookBody{
	Status: rewrite.WebhookStatusInactive,
})

if err != nil {
	log.Fatal(err)
}

hooks, err := client.Webhooks.List(context.Background(), &rewrite.RESTGetListWebhooksQueryParams{Limit: 10})

if err != nil {
	log.Fatal(err)
}

fmt.Printf("%+v\n", hooks)
```

<div align="center">

### Messages

</div>

```go
createdMessage, err := client.Messages.Send(context.Background(), rewrite.SendMessageOptions{
	IdempotencyKey: "msg-123",
	RESTPostSendMessageBody: rewrite.RESTPostSendMessageBody{
		To:      "+5511999999999",
		Content: "Hello from Rewrite",
	},
})

if err != nil {
	log.Fatal(err)
}

fmt.Printf("%+v\n", createdMessage)
```

<div align="center">

### OTP

</div>

```go
otp, err := client.OTP.Send(context.Background(), rewrite.SendOTPMessageOptions{
	IdempotencyKey: "otp-123",
	RESTPostSendOTPMessageBody: rewrite.RESTPostSendOTPMessageBody{
		To:        "+5511999999999",
		Prefix:    "Rewrite",
		ExpiresIn: 5,
	},
})

if err != nil {
	log.Fatal(err)
}

verified, err := client.OTP.Verify(context.Background(), rewrite.VerifyOTPOptions{
	ID: rewrite.Snowflake(otp.Data.ID),
	RESTPostVerifyOTPCodeBody: rewrite.RESTPostVerifyOTPCodeBody{
		To:   "+5511999999999",
		Code: "123456",
	},
})

if err != nil {
	log.Fatal(err)
}

fmt.Printf("otp=%+v verified=%+v\n", otp, verified)
```

<div align="center">

### Contacts And Segments

</div>

```go
contact, err := client.Contacts.Create(context.Background(), rewrite.RESTPostCreateContactBody{
	Phone: "+5511999999999",
	Name:  "Ada Lovelace",
	Tags: map[string]any{
		"source": "landing-page",
	},
})

if err != nil {
	log.Fatal(err)
}

segment, err := client.Segments.Create(context.Background(), rewrite.RESTPostCreateSegmentBody{
	Name:        "vip",
	Description: rewrite.NewNullableString("Customers with priority handling"),
})

if err != nil {
	log.Fatal(err)
}

_, err = client.Segments.AttachContact(context.Background(), string(segment.Data.ID), rewrite.RESTPostAttachSegmentContactBody{
	ContactID: contact.Data.ID,
})

if err != nil {
	log.Fatal(err)
}
```

<div align="center">

## Error Handling

Requests run through the SDK REST client. HTTP failures can return `HTTPError`.

</div>

```go
_, err := client.Webhooks.Get(context.Background(), "invalid_id")

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

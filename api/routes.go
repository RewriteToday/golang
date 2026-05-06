package api

import (
	"fmt"
	"net/url"
	"strconv"
	"strings"
)

const (
	// APIBaseURL is the canonical Rewrite API base URL.
	APIBaseURL = "https://api.rewritetoday.com"
)

// Routes exposes helper builders for Rewrite API routes.
var Routes = RouteRegistry{
	APIKeys:     APIKeyRoutes{},
	Contacts:    ContactRoutes{},
	Deliveries:  DeliveryRoutes{},
	Logs:        LogRoutes{},
	Messages:    MessageRoutes{},
	OTP:         OTPRoutes{},
	Segments:    SegmentRoutes{Contacts: SegmentContactRoutes{}},
	Tags:        TagRoutes{},
	Templates:   TemplateRoutes{},
	WebhookLogs: WebhookLogRoutes{},
	Webhooks:    WebhookRoutes{},
}

// RouteRegistry groups route builders by public resource.
type RouteRegistry struct {
	APIKeys     APIKeyRoutes
	Contacts    ContactRoutes
	Deliveries  DeliveryRoutes
	Logs        LogRoutes
	Messages    MessageRoutes
	OTP         OTPRoutes
	Segments    SegmentRoutes
	Tags        TagRoutes
	Templates   TemplateRoutes
	WebhookLogs WebhookLogRoutes
	Webhooks    WebhookRoutes
}

type APIKeyRoutes struct{}
type ContactRoutes struct{}
type DeliveryRoutes struct{}
type LogRoutes struct{}
type MessageRoutes struct{}
type OTPRoutes struct{}
type SegmentRoutes struct {
	Contacts SegmentContactRoutes
}
type SegmentContactRoutes struct{}
type TagRoutes struct{}
type TemplateRoutes struct{}
type WebhookLogRoutes struct{}
type WebhookRoutes struct{}

func (APIKeyRoutes) List(options *RESTGetListAPIKeysQueryParams) string {
	return appendQuery("/api-keys", createCursorQuery((*RESTCursorOptions)(options)))
}

func (APIKeyRoutes) Create() string {
	return "/api-keys"
}

func (APIKeyRoutes) Sweep() string {
	return "/api-keys"
}

func (APIKeyRoutes) Get(id string) string {
	return fmt.Sprintf("/api-keys/%s", id)
}

func (APIKeyRoutes) Update(id string) string {
	return fmt.Sprintf("/api-keys/%s", id)
}

func (APIKeyRoutes) Delete(id string) string {
	return fmt.Sprintf("/api-keys/%s", id)
}

func (APIKeyRoutes) Logs(id string, options *RESTGetListAPIKeyLogsQueryParams) string {
	return appendQuery(fmt.Sprintf("/api-keys/%s/logs", id), createLogsListQuery(options))
}

func (ContactRoutes) List(options *RESTGetListContactsQueryParams) string {
	return appendQuery("/contacts", createCursorQuery((*RESTCursorOptions)(options)))
}

func (ContactRoutes) Create() string {
	return "/contacts"
}

func (ContactRoutes) Sweep() string {
	return "/contacts"
}

func (ContactRoutes) Batch() string {
	return "/contacts/batch"
}

func (ContactRoutes) Get(identifier string) string {
	return fmt.Sprintf("/contacts/%s", identifier)
}

func (ContactRoutes) Update(id string) string {
	return fmt.Sprintf("/contacts/%s", id)
}

func (ContactRoutes) Delete(id string) string {
	return fmt.Sprintf("/contacts/%s", id)
}

func (ContactRoutes) AddTags(id string) string {
	return fmt.Sprintf("/contacts/%s/tags", id)
}

func (ContactRoutes) RemoveTags(id string) string {
	return fmt.Sprintf("/contacts/%s/tags", id)
}

func (DeliveryRoutes) List(options *RESTGetListDeliveriesQueryParams) string {
	return appendQuery("/deliveries", createDeliveriesListQuery(options))
}

func (DeliveryRoutes) Get(id string) string {
	return fmt.Sprintf("/deliveries/%s", id)
}

func (DeliveryRoutes) ByWebhook(id string, options *RESTGetListWebhookDeliveriesQueryParams) string {
	return Routes.Webhooks.Deliveries(id, options)
}

func (LogRoutes) List(options *RESTGetListLogsQueryParams) string {
	return appendQuery("/logs", createLogsListQuery(options))
}

func (LogRoutes) Get(id string) string {
	return fmt.Sprintf("/logs/%s", id)
}

func (MessageRoutes) Send() string {
	return "/messages"
}

func (r MessageRoutes) Create() string {
	return r.Send()
}

func (MessageRoutes) Batch() string {
	return "/messages/batch"
}

func (MessageRoutes) Cancel(id string) string {
	return fmt.Sprintf("/messages/%s/cancel", id)
}

func (MessageRoutes) Get(id string) string {
	return fmt.Sprintf("/messages/%s", id)
}

func (MessageRoutes) List(options *RESTGetListMessagesQueryParams) string {
	return appendQuery("/messages", createMessagesListQuery(options))
}

func (OTPRoutes) Send() string {
	return "/otp"
}

func (r OTPRoutes) Create() string {
	return r.Send()
}

func (OTPRoutes) Verify(id string) string {
	return fmt.Sprintf("/otp/%s/verify", id)
}

func (SegmentRoutes) List(options *RESTGetListSegmentsQueryParams) string {
	return appendQuery("/segments", createCursorQuery((*RESTCursorOptions)(options)))
}

func (SegmentRoutes) Create() string {
	return "/segments"
}

func (SegmentRoutes) Sweep() string {
	return "/segments"
}

func (SegmentRoutes) Get(id string) string {
	return fmt.Sprintf("/segments/%s", id)
}

func (SegmentRoutes) Update(id string) string {
	return fmt.Sprintf("/segments/%s", id)
}

func (SegmentRoutes) Delete(id string) string {
	return fmt.Sprintf("/segments/%s", id)
}

func (SegmentContactRoutes) List(id string, options *RESTGetListSegmentContactsQueryParams) string {
	return appendQuery(fmt.Sprintf("/segments/%s/contacts", id), createCursorQuery((*RESTCursorOptions)(options)))
}

func (SegmentContactRoutes) Attach(id string) string {
	return fmt.Sprintf("/segments/%s/contacts", id)
}

func (SegmentContactRoutes) AttachMany(id string) string {
	return fmt.Sprintf("/segments/%s/contacts/attach", id)
}

func (SegmentContactRoutes) DetachMany(id string) string {
	return fmt.Sprintf("/segments/%s/contacts/detach", id)
}

func (SegmentContactRoutes) Detach(id, contactID string) string {
	return fmt.Sprintf("/segments/%s/contacts/%s", id, contactID)
}

func (TagRoutes) List() string {
	return "/tags"
}

func (TagRoutes) Create() string {
	return "/tags"
}

func (TagRoutes) Get(id string) string {
	return fmt.Sprintf("/tags/%s", id)
}

func (TagRoutes) Update(id string) string {
	return fmt.Sprintf("/tags/%s", id)
}

func (TagRoutes) Delete(id string) string {
	return fmt.Sprintf("/tags/%s", id)
}

func (TemplateRoutes) List(options *RESTGetListTemplatesQueryParams) string {
	return appendQuery("/templates", createTemplatesListQuery(options))
}

func (TemplateRoutes) Create() string {
	return "/templates"
}

func (TemplateRoutes) Sweep() string {
	return "/templates"
}

func (TemplateRoutes) Update(id string) string {
	return fmt.Sprintf("/templates/%s", id)
}

func (TemplateRoutes) Delete(id string) string {
	return fmt.Sprintf("/templates/%s", id)
}

func (TemplateRoutes) Duplicate(id string) string {
	return fmt.Sprintf("/templates/%s/duplicate", id)
}

func (TemplateRoutes) Get(identifier string, options ...*RESTGetTemplateQueryParams) string {
	return appendQuery(fmt.Sprintf("/templates/%s", identifier), createTemplateGetQuery(firstTemplateQueryOptions(options)))
}

func (WebhookRoutes) List(options *RESTGetListWebhooksQueryParams) string {
	return appendQuery("/webhooks", createCursorQuery((*RESTCursorOptions)(options)))
}

func (WebhookRoutes) Create() string {
	return "/webhooks"
}

func (WebhookRoutes) Sweep() string {
	return "/webhooks"
}

func (WebhookRoutes) Update(id string) string {
	return fmt.Sprintf("/webhooks/%s", id)
}

func (WebhookRoutes) Delete(id string) string {
	return fmt.Sprintf("/webhooks/%s", id)
}

func (WebhookRoutes) Get(id string) string {
	return fmt.Sprintf("/webhooks/%s", id)
}

func (WebhookRoutes) Deliveries(id string, options *RESTGetListWebhookDeliveriesQueryParams) string {
	return appendQuery(fmt.Sprintf("/webhooks/%s/deliveries", id), createWebhookDeliveriesListQuery(options))
}

func (r WebhookRoutes) Logs(id string, options *RESTGetListWebhookLogsQueryParams) string {
	return r.Deliveries(id, (*RESTGetListWebhookDeliveriesQueryParams)(options))
}

func (WebhookLogRoutes) List(id string, options *RESTGetListWebhookLogsQueryParams) string {
	return Routes.Webhooks.Logs(id, options)
}

func appendQuery(route, query string) string {
	if query == "" {
		return route
	}
	return route + "?" + query
}

func createCursorQuery(options *RESTCursorOptions) string {
	limit := 15
	if options != nil && options.Limit > 0 {
		limit = options.Limit
	}

	parts := []string{"limit=" + url.QueryEscape(strconv.Itoa(limit))}
	if options != nil {
		if options.After != "" {
			parts = append(parts, "after="+url.QueryEscape(string(options.After)))
		}
		if options.Before != "" {
			parts = append(parts, "before="+url.QueryEscape(string(options.Before)))
		}
	}

	return strings.Join(parts, "&")
}

func createTemplatesListQuery(options *RESTGetListTemplatesQueryParams) string {
	cursor := RESTCursorOptions{}
	if options != nil {
		cursor = options.RESTCursorOptions
	}

	parts := []string{createCursorQuery(&cursor)}
	if options != nil {
		if withI18n := options.withI18n(); withI18n != nil {
			parts = append(parts, "withi18n="+url.QueryEscape(strconv.FormatBool(*withI18n)))
		}
	}

	return strings.Join(parts, "&")
}

func createTemplateGetQuery(options *RESTGetTemplateQueryParams) string {
	if options == nil {
		return ""
	}

	parts := make([]string, 0, 1)
	if withI18n := options.withI18n(); withI18n != nil {
		parts = append(parts, "withi18n="+url.QueryEscape(strconv.FormatBool(*withI18n)))
	}

	return strings.Join(parts, "&")
}

func createMessagesListQuery(options *RESTGetListMessagesQueryParams) string {
	cursor := RESTCursorOptions{}
	if options != nil {
		cursor = options.RESTCursorOptions
	}

	parts := []string{createCursorQuery(&cursor)}
	if options != nil {
		if options.Status != "" {
			parts = append(parts, "status="+url.QueryEscape(string(options.Status)))
		}
		if options.Country != "" {
			parts = append(parts, "country="+url.QueryEscape(string(options.Country)))
		}
		if options.Encoding != "" {
			parts = append(parts, "encoding="+url.QueryEscape(string(options.Encoding)))
		}
		if options.Scheduled != nil {
			parts = append(parts, "scheduled="+url.QueryEscape(strconv.FormatBool(*options.Scheduled)))
		}
		if options.WithCounts != nil {
			parts = append(parts, "withCounts="+url.QueryEscape(strconv.FormatBool(*options.WithCounts)))
		}
	}

	return strings.Join(parts, "&")
}

func createLogsListQuery(options *RESTGetListLogsQueryParams) string {
	cursor := RESTCursorOptions{}
	if options != nil {
		cursor = options.RESTCursorOptions
	}

	parts := []string{createCursorQuery(&cursor)}
	if options != nil {
		if options.Code > 0 {
			parts = append(parts, "code="+url.QueryEscape(strconv.Itoa(options.Code)))
		}
		if options.Method != "" {
			parts = append(parts, "method="+url.QueryEscape(options.Method))
		}
		if options.Endpoint != "" {
			parts = append(parts, "endpoint="+url.QueryEscape(options.Endpoint))
		}
	}

	return strings.Join(parts, "&")
}

func createDeliveriesListQuery(options *RESTGetListDeliveriesQueryParams) string {
	cursor := RESTCursorOptions{}
	if options != nil {
		cursor = options.RESTCursorOptions
	}

	parts := []string{createCursorQuery(&cursor)}
	if options != nil {
		if options.WebhookID != "" {
			parts = append(parts, "webhookId="+url.QueryEscape(string(options.WebhookID)))
		}
		if options.MessageID != "" {
			parts = append(parts, "messageId="+url.QueryEscape(string(options.MessageID)))
		}
		if options.Type != "" {
			parts = append(parts, "type="+url.QueryEscape(string(options.Type)))
		}
	}

	return strings.Join(parts, "&")
}

func createWebhookDeliveriesListQuery(options *RESTGetListWebhookDeliveriesQueryParams) string {
	cursor := RESTCursorOptions{}
	if options != nil {
		cursor = options.RESTCursorOptions
	}

	parts := []string{createCursorQuery(&cursor)}
	if options != nil {
		if options.Type != "" {
			parts = append(parts, "type="+url.QueryEscape(string(options.Type)))
		}
		if options.Status != "" {
			parts = append(parts, "status="+url.QueryEscape(string(options.Status)))
		}
		if options.Code > 0 {
			parts = append(parts, "code="+url.QueryEscape(strconv.Itoa(options.Code)))
		}
		if options.Attempt > 0 {
			parts = append(parts, "attempt="+url.QueryEscape(strconv.Itoa(options.Attempt)))
		}
		if options.MessageID != "" {
			parts = append(parts, "messageId="+url.QueryEscape(string(options.MessageID)))
		}
	}

	return strings.Join(parts, "&")
}

func firstTemplateQueryOptions(options []*RESTGetTemplateQueryParams) *RESTGetTemplateQueryParams {
	if len(options) == 0 {
		return nil
	}
	return options[0]
}

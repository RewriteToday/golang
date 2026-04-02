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
	Logs:        LogRoutes{},
	Messages:    MessageRoutes{},
	OTP:         OTPRoutes{},
	Segments:    SegmentRoutes{Contacts: SegmentContactRoutes{}},
	Templates:   TemplateRoutes{},
	WebhookLogs: WebhookLogRoutes{},
	Webhooks:    WebhookRoutes{},
}

// RouteRegistry groups route builders by public resource.
type RouteRegistry struct {
	APIKeys     APIKeyRoutes
	Contacts    ContactRoutes
	Logs        LogRoutes
	Messages    MessageRoutes
	OTP         OTPRoutes
	Segments    SegmentRoutes
	Templates   TemplateRoutes
	WebhookLogs WebhookLogRoutes
	Webhooks    WebhookRoutes
}

// APIKeyRoutes builds public API key endpoints.
type APIKeyRoutes struct{}

// ContactRoutes builds contact endpoints.
type ContactRoutes struct{}

// LogRoutes builds log endpoints.
type LogRoutes struct{}

// MessageRoutes builds message endpoints.
type MessageRoutes struct{}

// OTPRoutes builds OTP endpoints.
type OTPRoutes struct{}

// SegmentRoutes builds segment endpoints.
type SegmentRoutes struct {
	Contacts SegmentContactRoutes
}

// SegmentContactRoutes builds nested segment contact endpoints.
type SegmentContactRoutes struct{}

// TemplateRoutes builds template endpoints.
type TemplateRoutes struct{}

// WebhookLogRoutes builds webhook log endpoints.
type WebhookLogRoutes struct{}

// WebhookRoutes builds webhook endpoints.
type WebhookRoutes struct{}

// Delete returns DELETE /api-keys/:apiKeyId.
func (APIKeyRoutes) Delete(apiKeyID string) string {
	return fmt.Sprintf("/api-keys/%s", apiKeyID)
}

// Get returns GET /logs/:logId.
func (LogRoutes) Get(logID string) string {
	return fmt.Sprintf("/logs/%s", logID)
}

// Send returns POST /messages.
func (MessageRoutes) Send() string {
	return "/messages"
}

// Create is a compatibility alias for Send.
func (r MessageRoutes) Create() string {
	return r.Send()
}

// Batch returns POST /messages/batch.
func (MessageRoutes) Batch() string {
	return "/messages/batch"
}

// Cancel returns POST /messages/:messageId/cancel.
func (MessageRoutes) Cancel(messageID string) string {
	return fmt.Sprintf("/messages/%s/cancel", messageID)
}

// Get returns GET /messages/:messageId.
func (MessageRoutes) Get(messageID string) string {
	return fmt.Sprintf("/messages/%s", messageID)
}

// List returns GET /messages with cursor and filter query params.
func (MessageRoutes) List(options *RESTGetListMessagesQueryParams) string {
	return appendQuery("/messages", createMessagesListQuery(options))
}

// Send returns POST /otp.
func (OTPRoutes) Send() string {
	return "/otp"
}

// Create is a compatibility alias for Send.
func (r OTPRoutes) Create() string {
	return r.Send()
}

// Verify returns POST /otp/:otpId/verify.
func (OTPRoutes) Verify(otpID string) string {
	return fmt.Sprintf("/otp/%s/verify", otpID)
}

// List returns GET /contacts with cursor query params.
func (ContactRoutes) List(options *RESTGetListContactsQueryParams) string {
	return appendQuery("/contacts", createCursorQuery((*RESTCursorOptions)(options)))
}

// Create returns POST /contacts.
func (ContactRoutes) Create() string {
	return "/contacts"
}

// Get returns GET /contacts/:identifier.
func (ContactRoutes) Get(identifier string) string {
	return fmt.Sprintf("/contacts/%s", identifier)
}

// Update returns PATCH /contacts/:id.
func (ContactRoutes) Update(id string) string {
	return fmt.Sprintf("/contacts/%s", id)
}

// Delete returns DELETE /contacts/:id.
func (ContactRoutes) Delete(id string) string {
	return fmt.Sprintf("/contacts/%s", id)
}

// List returns GET /segments with cursor query params.
func (SegmentRoutes) List(options *RESTGetListSegmentsQueryParams) string {
	return appendQuery("/segments", createCursorQuery((*RESTCursorOptions)(options)))
}

// Create returns POST /segments.
func (SegmentRoutes) Create() string {
	return "/segments"
}

// Get returns GET /segments/:id.
func (SegmentRoutes) Get(id string) string {
	return fmt.Sprintf("/segments/%s", id)
}

// Update returns PATCH /segments/:id.
func (SegmentRoutes) Update(id string) string {
	return fmt.Sprintf("/segments/%s", id)
}

// Delete returns DELETE /segments/:id.
func (SegmentRoutes) Delete(id string) string {
	return fmt.Sprintf("/segments/%s", id)
}

// List returns GET /segments/:id/contacts with cursor query params.
func (SegmentContactRoutes) List(id string, options *RESTGetListSegmentContactsQueryParams) string {
	return appendQuery(
		fmt.Sprintf("/segments/%s/contacts", id),
		createCursorQuery((*RESTCursorOptions)(options)),
	)
}

// Attach returns POST /segments/:id/contacts.
func (SegmentContactRoutes) Attach(id string) string {
	return fmt.Sprintf("/segments/%s/contacts", id)
}

// Detach returns DELETE /segments/:id/contacts/:contactId.
func (SegmentContactRoutes) Detach(id, contactID string) string {
	return fmt.Sprintf("/segments/%s/contacts/%s", id, contactID)
}

// List returns GET /templates with cursor query and optional i18n expansion.
func (TemplateRoutes) List(options *RESTGetListTemplatesQueryParams) string {
	return appendQuery("/templates", createTemplatesListQuery(options))
}

// Create returns POST /templates.
func (TemplateRoutes) Create() string {
	return "/templates"
}

// Update returns PATCH /templates/:templateId.
func (TemplateRoutes) Update(templateID string) string {
	return fmt.Sprintf("/templates/%s", templateID)
}

// Delete returns DELETE /templates/:templateId.
func (TemplateRoutes) Delete(templateID string) string {
	return fmt.Sprintf("/templates/%s", templateID)
}

// Get returns GET /templates/:identifier with optional query params.
func (TemplateRoutes) Get(identifier string, options ...*RESTGetTemplateQueryParams) string {
	query := createTemplateGetQuery(firstTemplateQueryOptions(options))
	return appendQuery(fmt.Sprintf("/templates/%s", identifier), query)
}

// List returns GET /webhooks with cursor query.
func (WebhookRoutes) List(options *RESTGetListWebhooksQueryParams) string {
	return appendQuery("/webhooks", createCursorQuery(options))
}

// Create returns POST /webhooks.
func (WebhookRoutes) Create() string {
	return "/webhooks"
}

// Update returns PATCH /webhooks/:webhookId.
func (WebhookRoutes) Update(webhookID string) string {
	return fmt.Sprintf("/webhooks/%s", webhookID)
}

// Delete returns DELETE /webhooks/:webhookId.
func (WebhookRoutes) Delete(webhookID string) string {
	return fmt.Sprintf("/webhooks/%s", webhookID)
}

// Get returns GET /webhooks/:webhookId.
func (WebhookRoutes) Get(webhookID string) string {
	return fmt.Sprintf("/webhooks/%s", webhookID)
}

// Logs returns GET /webhooks/:id/logs with cursor query and filters.
func (WebhookRoutes) Logs(webhookID string, options *RESTGetListWebhookLogsQueryParams) string {
	return appendQuery(fmt.Sprintf("/webhooks/%s/logs", webhookID), createWebhookLogsListQuery(options))
}

// List is a compatibility alias for Routes.Webhooks.Logs.
func (WebhookLogRoutes) List(webhookID string, options *RESTGetListWebhookLogsQueryParams) string {
	return Routes.Webhooks.Logs(webhookID, options)
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
	if withI18n := options.withI18n(); withI18n != nil {
		parts = append(parts, "withi18n="+url.QueryEscape(strconv.FormatBool(*withI18n)))
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
	}

	return strings.Join(parts, "&")
}

func createWebhookLogsListQuery(options *RESTGetListWebhookLogsQueryParams) string {
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
	}

	return strings.Join(parts, "&")
}

func firstTemplateQueryOptions(options []*RESTGetTemplateQueryParams) *RESTGetTemplateQueryParams {
	if len(options) == 0 {
		return nil
	}

	return options[0]
}

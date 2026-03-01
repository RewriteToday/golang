package api

// Snowflake is the unique identifier format used by Rewrite resources.
type Snowflake string

// Cursor describes pagination state for list responses.
type Cursor struct {
	Persist bool       `json:"persist"`
	Next    *Snowflake `json:"next,omitempty"`
}

// APIAPIKey represents an API key entity returned by the Rewrite API.
type APIAPIKey struct {
	ID        Snowflake     `json:"id"`
	Name      string        `json:"name"`
	Prefix    string        `json:"prefix"`
	Scopes    []APIKeyScope `json:"scopes"`
	CreatedAt string        `json:"createdAt"`
}

// APICreatedAPIKey represents the one-time response payload from API key creation.
type APICreatedAPIKey struct {
	ID        Snowflake `json:"id"`
	Key       string    `json:"key"`
	CreatedAt string    `json:"createdAt"`
}

// APIKeyScope enumerates the supported API key permissions.
type APIKeyScope string

const (
	// APIKeyScopeWildcard grants every available permission.
	APIKeyScopeWildcard APIKeyScope = "*"
	// APIKeyScopeReadProject allows reading project details.
	APIKeyScopeReadProject APIKeyScope = "project:read"
	// APIKeyScopeWriteProject allows updating project details.
	APIKeyScopeWriteProject APIKeyScope = "project:write"
	// APIKeyScopeReadAPIKeys allows listing API keys.
	APIKeyScopeReadAPIKeys APIKeyScope = "project:api_keys:read"
	// APIKeyScopeWriteTemplate allows creating and updating templates.
	APIKeyScopeWriteTemplate APIKeyScope = "project:templates:write"
	// APIKeyScopeReadTemplates allows listing and reading templates.
	APIKeyScopeReadTemplates APIKeyScope = "project:templates:read"
	// APIKeyScopeReadPayments allows reading payment information.
	APIKeyScopeReadPayments APIKeyScope = "projects:payments:read"
	// APIKeyScopeReadWebhooks allows listing and reading webhooks.
	APIKeyScopeReadWebhooks APIKeyScope = "projects:webhooks:read"
	// APIKeyScopeWriteWebhooks allows creating and updating webhooks.
	APIKeyScopeWriteWebhooks APIKeyScope = "projects:webhooks:write"
)

// APITemplate represents a message template.
type APITemplate struct {
	ID        Snowflake             `json:"id"`
	Name      string                `json:"name"`
	Content   *string               `json:"content"`
	Variables []APITemplateVariable `json:"variables"`
	CreatedAt string                `json:"createdAt"`
}

// APICreatedTemplate represents the create-template response payload.
type APICreatedTemplate struct {
	ID        Snowflake `json:"id"`
	CreatedAt string    `json:"createdAt"`
}

// APITemplateVariable represents a named variable in a template.
type APITemplateVariable struct {
	Name     string `json:"name"`
	Fallback string `json:"fallback"`
}

// APIWebhook represents a webhook endpoint configuration.
type APIWebhook struct {
	ID        Snowflake          `json:"id"`
	Name      *string            `json:"name"`
	Endpoint  string             `json:"endpoint"`
	Events    []WebhookEventType `json:"events"`
	Status    WebhookStatus      `json:"status"`
	CreatedAt string             `json:"createdAt"`
}

// APICreatedWebhook represents the create-webhook response payload.
type APICreatedWebhook struct {
	ID Snowflake `json:"id"`
}

// WebhookEventType enumerates webhook events exposed by Rewrite.
type WebhookEventType string

const (
	// WebhookEventTypeSMSQueued fires when an SMS enters the queue.
	WebhookEventTypeSMSQueued WebhookEventType = "sms.queued"
	// WebhookEventTypeSMSDelivered fires when an SMS is delivered.
	WebhookEventTypeSMSDelivered WebhookEventType = "sms.delivered"
	// WebhookEventTypeSMSScheduled fires when an SMS is scheduled.
	WebhookEventTypeSMSScheduled WebhookEventType = "sms.scheduled"
	// WebhookEventTypeSMSFailed fires when delivery fails.
	WebhookEventTypeSMSFailed WebhookEventType = "sms.failed"
	// WebhookEventTypeSMSCanceled fires when sending is canceled.
	WebhookEventTypeSMSCanceled WebhookEventType = "sms.canceled"
)

// WebhookStatus represents the activation status for a webhook.
type WebhookStatus string

const (
	// WebhookStatusActive means the webhook is receiving events.
	WebhookStatusActive WebhookStatus = "ACTIVE"
	// WebhookStatusInactive means the webhook is paused.
	WebhookStatusInactive WebhookStatus = "INACTIVE"
)

// RESTCursorOptions configures cursor-based pagination.
type RESTCursorOptions struct {
	Limit  int       `json:"limit,omitempty"`
	After  Snowflake `json:"after,omitempty"`
	Before Snowflake `json:"before,omitempty"`
}

// APIValidationError describes validation details returned by the API.
type APIValidationError struct {
	Message  string         `json:"message"`
	Detailed map[string]any `json:"detailed,omitempty"`
}

// APIResponse is the standard Rewrite API response envelope.
type APIResponse[T any] struct {
	OK      bool                `json:"ok"`
	Data    T                   `json:"data"`
	Cursor  *Cursor             `json:"cursor,omitempty"`
	Code    string              `json:"code,omitempty"`
	Message string              `json:"message,omitempty"`
	Errors  *APIValidationError `json:"errors,omitempty"`
}

// RESTGetWebhookData corresponds to GET /projects/:id/webhooks/:webhookId.
type RESTGetWebhookData = APIResponse[APIWebhook]

// RESTPostCreateWebhookData corresponds to POST /projects/:id/webhooks.
type RESTPostCreateWebhookData = APIResponse[APICreatedWebhook]

// RESTPostCreateWebhookBody is the request body for webhook creation.
type RESTPostCreateWebhookBody struct {
	Name     string             `json:"name,omitempty"`
	Endpoint string             `json:"endpoint"`
	Events   []WebhookEventType `json:"events"`
}

// RESTDeleteWebhookData corresponds to DELETE /projects/:id/webhooks/:webhookId.
type RESTDeleteWebhookData = APIResponse[any]

// RESTPatchUpdateWebhookData corresponds to PATCH /projects/:id/webhooks/:webhookId.
type RESTPatchUpdateWebhookData = APIResponse[any]

// RESTPatchUpdateWebhookBody is the request body for webhook updates.
type RESTPatchUpdateWebhookBody struct {
	Name     *string            `json:"name,omitempty"`
	Endpoint string             `json:"endpoint,omitempty"`
	Events   []WebhookEventType `json:"events,omitempty"`
	Status   WebhookStatus      `json:"status,omitempty"`
}

// RESTGetListWebhooksData corresponds to GET /projects/:id/webhooks.
type RESTGetListWebhooksData = APIResponse[[]APIWebhook]

// RESTGetListWebhooksQueryParams corresponds to webhook list query params.
type RESTGetListWebhooksQueryParams = RESTCursorOptions

// RESTGetListTemplatesData corresponds to GET /projects/:id/templates.
type RESTGetListTemplatesData = APIResponse[[]APITemplate]

// RESTGetListTemplatesQueryParams corresponds to template list query params.
type RESTGetListTemplatesQueryParams = RESTCursorOptions

// RESTPostCreateTemplateData corresponds to POST /projects/:id/templates.
type RESTPostCreateTemplateData = APIResponse[APICreatedTemplate]

// RESTPostCreateTemplateBody is the request body for template creation.
type RESTPostCreateTemplateBody struct {
	Name      string                `json:"name"`
	Content   string                `json:"content"`
	Variables []APITemplateVariable `json:"variables"`
}

// RESTPatchUpdateTemplateData corresponds to PATCH /projects/:id/templates/:templateId.
type RESTPatchUpdateTemplateData = APIResponse[any]

// RESTPatchUpdateTemplateBody is the request body for template updates.
type RESTPatchUpdateTemplateBody struct {
	Content   string                `json:"content,omitempty"`
	Variables []APITemplateVariable `json:"variables,omitempty"`
}

// RESTDeleteTemplateData corresponds to DELETE /projects/:id/templates/:templateId.
type RESTDeleteTemplateData = APIResponse[any]

// RESTGetTemplateData corresponds to GET /projects/:id/templates/:identifier.
type RESTGetTemplateData = APIResponse[APITemplate]

// RESTGetListAPIKeysData corresponds to GET /projects/:id/api-keys.
type RESTGetListAPIKeysData = APIResponse[[]APIAPIKey]

// RESTGetListAPIKeysQueryParams corresponds to API key list query params.
type RESTGetListAPIKeysQueryParams = RESTCursorOptions

// RESTPostCreateAPIKeyData corresponds to POST /projects/:id/api-keys.
type RESTPostCreateAPIKeyData = APIResponse[APICreatedAPIKey]

// RESTPostCreateAPIKeyBody is the request body for API key creation.
type RESTPostCreateAPIKeyBody struct {
	Name   string        `json:"name"`
	Scopes []APIKeyScope `json:"scopes,omitempty"`
}

// RESTDeleteAPIKeyData corresponds to DELETE /projects/:id/api-keys/:apiKeyId.
type RESTDeleteAPIKeyData = APIResponse[any]

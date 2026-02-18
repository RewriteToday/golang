package api

// Snowflake is the unique identifier format used by Rewrite resources.
type Snowflake string

// APIAPIKey represents an API key entity returned by the Rewrite API.
type APIAPIKey struct {
	ID        Snowflake     `json:"id"`
	Name      string        `json:"name"`
	ProjectID Snowflake     `json:"projectId"`
	Scopes    []APIKeyScope `json:"scopes"`
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
	APIKeyScopeReadAPIKeys APIKeyScope = "project:api_key:read"
	// APIKeyScopeWriteTemplate allows creating and updating templates.
	APIKeyScopeWriteTemplate APIKeyScope = "project:template:write"
	// APIKeyScopeReadTemplates allows listing and reading templates.
	APIKeyScopeReadTemplates APIKeyScope = "project:template:read"
	// APIKeyScopeReadPayments allows reading payment information.
	APIKeyScopeReadPayments APIKeyScope = "project:payment:read"
	// APIKeyScopeReadWebhooks allows listing and reading webhooks.
	APIKeyScopeReadWebhooks APIKeyScope = "project:webhook:read"
	// APIKeyScopeWriteWebhooks allows creating and updating webhooks.
	APIKeyScopeWriteWebhooks APIKeyScope = "project:webhook:write"
)

// APIProject represents a Rewrite project.
type APIProject struct {
	ID      Snowflake `json:"id"`
	Name    string    `json:"name"`
	OwnerID Snowflake `json:"ownerId"`
	Icon    *string   `json:"icon,omitempty"`
}

// APITemplate represents a message template.
type APITemplate struct {
	ID        Snowflake             `json:"id"`
	Name      string                `json:"name"`
	ProjectID Snowflake             `json:"projectId"`
	Variables []APITemplateVariable `json:"variables"`
}

// APITemplateVariable represents a named variable in a template.
type APITemplateVariable struct {
	Name     string  `json:"name"`
	Fallback *string `json:"fallback,omitempty"`
}

// APIWebhook represents a webhook endpoint configuration.
type APIWebhook struct {
	ID        Snowflake          `json:"id"`
	Name      string             `json:"name"`
	Endpoint  string             `json:"endpoint"`
	Events    []WebhookEventType `json:"events"`
	Status    WebhookStatus      `json:"status"`
	ProjectID Snowflake          `json:"projectId"`
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
	Data    T                   `json:"data,omitempty"`
	Code    string              `json:"code,omitempty"`
	Message string              `json:"message,omitempty"`
	Errors  *APIValidationError `json:"errors,omitempty"`
}

// RESTGetWebhookData corresponds to GET /projects/:id/webhooks/:webhookId.
type RESTGetWebhookData = APIResponse[APIWebhook]

// RESTPostCreateWebhookData corresponds to POST /projects/:id/webhooks.
type RESTPostCreateWebhookData = APIResponse[APIWebhook]

// RESTPostCreateWebhookBody is the request body for webhook creation.
type RESTPostCreateWebhookBody struct {
	Name     string             `json:"name"`
	Endpoint string             `json:"endpoint"`
	Events   []WebhookEventType `json:"events"`
}

// RESTDeleteWebhookData corresponds to DELETE /projects/:id/webhooks/:webhookId.
type RESTDeleteWebhookData = APIResponse[any]

// RESTPatchUpdateWebhookData corresponds to PATCH /projects/:id/webhooks/:webhookId.
type RESTPatchUpdateWebhookData = APIResponse[APIWebhook]

// RESTPatchUpdateWebhookBody is the request body for webhook updates.
type RESTPatchUpdateWebhookBody struct {
	Name     string             `json:"name,omitempty"`
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
type RESTPostCreateTemplateData = APIResponse[APITemplate]

// RESTPostCreateTemplateBody is the request body for template creation.
type RESTPostCreateTemplateBody struct {
	Name      string                `json:"name"`
	Variables []APITemplateVariable `json:"variables"`
}

// RESTPatchUpdateTemplateData corresponds to PATCH /projects/:id/templates/:templateId.
type RESTPatchUpdateTemplateData = APIResponse[APITemplate]

// RESTPatchUpdateTemplateBody is the request body for template updates.
type RESTPatchUpdateTemplateBody struct {
	Name      string                `json:"name"`
	Variables []APITemplateVariable `json:"variables"`
}

// RESTDeleteTemplateData corresponds to DELETE /projects/:id/templates/:templateId.
type RESTDeleteTemplateData = APIResponse[any]

// RESTGetTemplateData corresponds to GET /projects/:id/templates/:templateId.
type RESTGetTemplateData = APIResponse[APITemplate]

// RESTGetListAPIKeysData corresponds to GET /projects/:id/api-keys.
type RESTGetListAPIKeysData = APIResponse[[]APIAPIKey]

// RESTGetListAPIKeysQueryParams corresponds to API key list query params.
type RESTGetListAPIKeysQueryParams = RESTCursorOptions

// RESTPostCreateAPIKeyData corresponds to POST /projects/:id/api-keys.
type RESTPostCreateAPIKeyData = APIResponse[APIAPIKey]

// RESTPostCreateAPIKeyBody is the request body for API key creation.
type RESTPostCreateAPIKeyBody struct {
	Name   string        `json:"name"`
	Scopes []APIKeyScope `json:"scopes"`
}

// RESTDeleteAPIKeyData corresponds to DELETE /projects/:id/api-keys/:apiKeyId.
type RESTDeleteAPIKeyData = APIResponse[any]

// RESTPostCreateProjectData corresponds to POST /projects.
type RESTPostCreateProjectData = APIResponse[APIProject]

// RESTPostCreateProjectBody is the request body for project creation.
type RESTPostCreateProjectBody struct {
	Name string `json:"name"`
}

// RESTPatchUpdateProjectData corresponds to PATCH /projects/:id.
type RESTPatchUpdateProjectData = APIResponse[APIProject]

// Null serializes as JSON null.
type Null struct{}

// MarshalJSON implements json.Marshaler.
func (Null) MarshalJSON() ([]byte, error) {
	return []byte("null"), nil
}

// RESTPatchUpdateProjectBody is the request body for project updates.
type RESTPatchUpdateProjectBody struct {
	Name string `json:"name,omitempty"`
	Icon *Null  `json:"icon,omitempty"`
}

// RESTDeleteProjectData corresponds to DELETE /projects/:id.
type RESTDeleteProjectData = APIResponse[any]

// RESTGetProjectData corresponds to GET /projects/:id.
type RESTGetProjectData = APIResponse[APIProject]

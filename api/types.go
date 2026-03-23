package api

import "encoding/json"

// Snowflake is the unique identifier format used by Rewrite resources.
type Snowflake string

// CountryCode is an ISO-like country identifier used by the Rewrite API.
type CountryCode string

// Cursor describes pagination state for list responses.
type Cursor struct {
	Persist bool       `json:"persist"`
	Next    *Snowflake `json:"next,omitempty"`
	Prev    *Snowflake `json:"prev,omitempty"`
}

// APIError describes the structured error payload returned by the API.
type APIError struct {
	Code     string         `json:"code"`
	Message  string         `json:"message"`
	Detailed map[string]any `json:"detailed,omitempty"`
}

// APIValidationError is kept as a compatibility alias for APIError.
type APIValidationError = APIError

// APIErrorResponse is the standard Rewrite API error envelope.
type APIErrorResponse struct {
	OK    bool      `json:"ok"`
	Error *APIError `json:"error,omitempty"`
}

// APIResponse is the standard Rewrite API success envelope.
type APIResponse[T any] struct {
	OK     bool      `json:"ok"`
	Data   T         `json:"data"`
	Cursor *Cursor   `json:"cursor,omitempty"`
	Error  *APIError `json:"error,omitempty"`
}

// NullableString represents an optional string field that can also be set to null.
type NullableString struct {
	set   bool
	value *string
}

// NewNullableString creates a set nullable string with a concrete value.
func NewNullableString(value string) NullableString {
	return NullableString{
		set:   true,
		value: &value,
	}
}

// NullString creates a set nullable string whose JSON value is null.
func NullString() NullableString {
	return NullableString{set: true}
}

// IsZero reports whether the field should be omitted during JSON encoding.
func (s NullableString) IsZero() bool {
	return !s.set
}

// MarshalJSON implements json.Marshaler.
func (s NullableString) MarshalJSON() ([]byte, error) {
	if s.value == nil {
		return []byte("null"), nil
	}

	return json.Marshal(*s.value)
}

// JSONNull represents an optional explicit JSON null value.
type JSONNull struct {
	set bool
}

// NewNull creates an explicit null marker.
func NewNull() JSONNull {
	return JSONNull{set: true}
}

// IsZero reports whether the null marker should be omitted during JSON encoding.
func (n JSONNull) IsZero() bool {
	return !n.set
}

// MarshalJSON implements json.Marshaler.
func (JSONNull) MarshalJSON() ([]byte, error) {
	return []byte("null"), nil
}

// APIKeyScope enumerates the supported API key permissions.
type APIKeyScope string

const (
	APIKeyScopeWildcard      APIKeyScope = "*"
	APIKeyScopeReadProject   APIKeyScope = "project:read"
	APIKeyScopeReadAPIKeys   APIKeyScope = "project:api_keys:read"
	APIKeyScopeWriteProject  APIKeyScope = "project:write"
	APIKeyScopeReadWebhooks  APIKeyScope = "project:webhooks:read"
	APIKeyScopeWriteTemplate APIKeyScope = "project:templates:write"
	APIKeyScopeReadTemplates APIKeyScope = "project:templates:read"
	APIKeyScopeWriteWebhooks APIKeyScope = "project:webhooks:write"
	APIKeyScopeWriteMessages APIKeyScope = "message:write"
	APIKeyScopeReadMessages  APIKeyScope = "message:read"
	APIKeyScopeReadLogs      APIKeyScope = "project:logs:read"
)

// MessageProvider identifies the provider that handled a message.
type MessageProvider string

const (
	MessageProviderComtele MessageProvider = "COMTELE"
)

// MessageType identifies the logical message kind.
type MessageType string

const (
	MessageTypeSMS MessageType = "SMS"
	MessageTypeOTP MessageType = "OTP"
)

// MessageEncoding identifies the persisted SMS encoding.
type MessageEncoding string

const (
	MessageEncodingGMS7 MessageEncoding = "GMS7"
	MessageEncodingGSM7 MessageEncoding = MessageEncodingGMS7
	MessageEncodingUCS2 MessageEncoding = "UCS2"
)

// MessageAnalysisEncoding identifies the analysis-time SMS encoding.
type MessageAnalysisEncoding string

const (
	MessageAnalysisEncodingGSM7 MessageAnalysisEncoding = "gsm7"
	MessageAnalysisEncodingUCS2 MessageAnalysisEncoding = "ucs2"
)

// MessageStatus identifies the current lifecycle state of a message.
type MessageStatus string

const (
	MessageStatusSent      MessageStatus = "SENT"
	MessageStatusQueued    MessageStatus = "QUEUED"
	MessageStatusFailed    MessageStatus = "FAILED"
	MessageStatusCanceled  MessageStatus = "CANCELED"
	MessageStatusScheduled MessageStatus = "SCHEDULED"
	MessageStatusDelivered MessageStatus = "DELIVERED"
)

// MessageSegmentationMode defines what the API should do when the content exceeds the segment limit.
type MessageSegmentationMode string

const (
	MessageSegmentationModeFail MessageSegmentationMode = "fail"
	MessageSegmentationModeTrim MessageSegmentationMode = "trim"
	MessageSegmentationModeSend MessageSegmentationMode = "send"
)

// MessageAnalysisReason explains the segmentation result.
type MessageAnalysisReason string

const (
	MessageAnalysisReasonFitsSingleSegment         MessageAnalysisReason = "fits"
	MessageAnalysisReasonSmartEncodingApplied      MessageAnalysisReason = "smart"
	MessageAnalysisReasonExceedsSingleSegmentLimit MessageAnalysisReason = "singleLimit"
	MessageAnalysisReasonContainsNonGsm7Characters MessageAnalysisReason = "nonGsm7"
	MessageAnalysisReasonContainsNonGSM7Characters MessageAnalysisReason = MessageAnalysisReasonContainsNonGsm7Characters
)

// OtpAttemptStatus identifies the current OTP attempt state.
type OtpAttemptStatus string

const (
	OtpAttemptStatusFailed     OtpAttemptStatus = "FAILED"
	OtpAttemptStatusPending    OtpAttemptStatus = "PENDING"
	OtpAttemptStatusExpired    OtpAttemptStatus = "EXPIRED"
	OtpAttemptStatusVerified   OtpAttemptStatus = "VERIFIED"
	OtpAttemptStatusSuperseded OtpAttemptStatus = "SUPERSEDED"
)

// WebhookEventType enumerates the supported webhook events.
type WebhookEventType string

const (
	WebhookEventTypeSMSOTP           WebhookEventType = "sms.otp"
	WebhookEventTypeMessageSent      WebhookEventType = "message.sent"
	WebhookEventTypeMessageBatch     WebhookEventType = "message.batch"
	WebhookEventTypeMessageQueued    WebhookEventType = "message.queued"
	WebhookEventTypeMessageFailed    WebhookEventType = "message.failed"
	WebhookEventTypeMessageCanceled  WebhookEventType = "message.canceled"
	WebhookEventTypeMessageDelivered WebhookEventType = "message.delivered"
	WebhookEventTypeMessageScheduled WebhookEventType = "message.scheduled"
)

// WebhookStatus identifies whether a webhook is active.
type WebhookStatus string

const (
	WebhookStatusActive   WebhookStatus = "ACTIVE"
	WebhookStatusInactive WebhookStatus = "INACTIVE"
)

// WebhookDeliveryStatus identifies a webhook delivery attempt outcome.
type WebhookDeliveryStatus string

const (
	WebhookDeliveryStatusFailed  WebhookDeliveryStatus = "FAILED"
	WebhookDeliveryStatusSuccess WebhookDeliveryStatus = "SUCCESS"
)

// APITemplateVariable represents a named variable in a template.
type APITemplateVariable struct {
	Name     string `json:"name"`
	Fallback string `json:"fallback,omitempty"`
}

// APITemplate represents a message template.
type APITemplate struct {
	ID          Snowflake              `json:"id"`
	Name        string                 `json:"name"`
	ProjectID   Snowflake              `json:"projectId,omitempty"`
	I18N        map[CountryCode]string `json:"i18n,omitempty"`
	Content     string                 `json:"content"`
	Variables   []APITemplateVariable  `json:"variables"`
	Description *string                `json:"description,omitempty"`
	CreatedAt   string                 `json:"createdAt"`
}

// APICreatedTemplate represents the create-template response payload.
type APICreatedTemplate struct {
	ID        Snowflake `json:"id"`
	CreatedAt string    `json:"createdAt"`
}

// APIWebhook represents a webhook endpoint configuration.
type APIWebhook struct {
	ID        Snowflake          `json:"id"`
	Name      string             `json:"name"`
	Secret    string             `json:"secret,omitempty"`
	Endpoint  string             `json:"endpoint"`
	Events    []WebhookEventType `json:"events"`
	Status    WebhookStatus      `json:"status"`
	ProjectID Snowflake          `json:"projectId,omitempty"`
	CreatedAt string             `json:"createdAt"`
}

// APICreatedWebhook represents the create-webhook response payload.
type APICreatedWebhook struct {
	ID        Snowflake `json:"id"`
	Secret    string    `json:"secret"`
	CreatedAt string    `json:"createdAt"`
}

// APIMessageTag represents a message tag.
type APIMessageTag struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

// APIMessageSegmentationOptions controls segmentation behavior on send.
type APIMessageSegmentationOptions struct {
	Max   int                     `json:"max"`
	Mode  MessageSegmentationMode `json:"mode,omitempty"`
	Smart bool                    `json:"smart,omitempty"`
}

// APIMessageAnalysisSegments describes the calculated segment usage.
type APIMessageAnalysisSegments struct {
	Concat int                   `json:"concat"`
	Count  int                   `json:"count"`
	Reason MessageAnalysisReason `json:"reason"`
	Single int                   `json:"single"`
}

// APIMessageAnalysis describes the calculated message metrics.
type APIMessageAnalysis struct {
	Characters int                        `json:"characters"`
	Encoding   MessageAnalysisEncoding    `json:"encoding"`
	Segments   APIMessageAnalysisSegments `json:"segments"`
}

// APIMessageError describes a provider-side error.
type APIMessageError struct {
	Code    any    `json:"code,omitempty"`
	Message string `json:"message"`
}

// APIMessage represents a persisted message read model.
type APIMessage struct {
	ID           Snowflake          `json:"id"`
	CreatedAt    Snowflake          `json:"createdAt"`
	Analysis     APIMessageAnalysis `json:"analysis"`
	To           string             `json:"to"`
	Type         MessageType        `json:"type"`
	Tags         []APIMessageTag    `json:"tags"`
	Status       MessageStatus      `json:"status"`
	Country      CountryCode        `json:"country"`
	Content      string             `json:"content"`
	Encoding     MessageEncoding    `json:"encoding"`
	TemplateID   *Snowflake         `json:"templateId,omitempty"`
	DeliveredAt  *string            `json:"deliveredAt,omitempty"`
	ScheduledAt  *string            `json:"scheduledAt,omitempty"`
	IsPayAsYouGo bool               `json:"isPayAsYouGo"`
}

// APIMessageWithoutAnalysis is the read model returned by the Node SDK for GET/LIST message endpoints.
type APIMessageWithoutAnalysis struct {
	ID           Snowflake       `json:"id"`
	CreatedAt    Snowflake       `json:"createdAt"`
	To           string          `json:"to"`
	Type         MessageType     `json:"type"`
	Tags         []APIMessageTag `json:"tags"`
	Status       MessageStatus   `json:"status"`
	Country      CountryCode     `json:"country"`
	Content      string          `json:"content"`
	Encoding     MessageEncoding `json:"encoding"`
	TemplateID   *Snowflake      `json:"templateId,omitempty"`
	DeliveredAt  *string         `json:"deliveredAt,omitempty"`
	ScheduledAt  *string         `json:"scheduledAt,omitempty"`
	IsPayAsYouGo bool            `json:"isPayAsYouGo"`
}

// APICreatedMessage represents the send-message response payload.
type APICreatedMessage struct {
	ID        Snowflake          `json:"id"`
	CreatedAt Snowflake          `json:"createdAt"`
	Analysis  APIMessageAnalysis `json:"analysis"`
}

// APIOTPMessage represents the send-OTP response payload.
type APIOTPMessage struct {
	ID        Snowflake `json:"id"`
	To        string    `json:"to"`
	Prefix    string    `json:"prefix"`
	CreatedAt string    `json:"createdAt"`
	ExpiresAt string    `json:"expiresAt"`
}

// APIOTPVerification represents the verify-OTP response payload.
type APIOTPVerification struct {
	ID         Snowflake `json:"id"`
	Valid      bool      `json:"valid"`
	VerifiedAt *string   `json:"verifiedAt"`
}

// APIWebhookOTPMetadata carries OTP metadata included in webhook payloads.
type APIWebhookOTPMetadata struct {
	ExpiresIn int    `json:"expiresIn"`
	ExpiresAt string `json:"expiresAt"`
	Prefix    string `json:"prefix"`
}

// APIWebhookEventData represents the message or batch payload embedded in webhook deliveries.
type APIWebhookEventData struct {
	ID          Snowflake              `json:"id,omitempty"`
	IDs         []Snowflake            `json:"ids,omitempty"`
	ProjectID   Snowflake              `json:"projectId"`
	Analysis    *APIMessageAnalysis    `json:"analysis,omitempty"`
	Content     string                 `json:"content,omitempty"`
	Country     CountryCode            `json:"country,omitempty"`
	DeliveredAt *string                `json:"deliveredAt,omitempty"`
	Error       *APIMessageError       `json:"error,omitempty"`
	OTP         *APIWebhookOTPMetadata `json:"otp,omitempty"`
	ScheduledAt *string                `json:"scheduledAt,omitempty"`
	Status      MessageStatus          `json:"status,omitempty"`
	Tags        []APIMessageTag        `json:"tags,omitempty"`
	TemplateID  *Snowflake             `json:"templateId,omitempty"`
	To          string                 `json:"to,omitempty"`
	Type        MessageType            `json:"type,omitempty"`
}

// APIWebhookEvent represents the stored webhook event payload.
type APIWebhookEvent struct {
	ID        Snowflake           `json:"id"`
	CreatedAt string              `json:"createdAt"`
	Type      WebhookEventType    `json:"type"`
	Data      APIWebhookEventData `json:"data"`
}

// APIWebhookLog represents a stored webhook delivery log.
type APIWebhookLog struct {
	ID        Snowflake             `json:"id"`
	CreatedAt string                `json:"createdAt"`
	WebhookID *Snowflake            `json:"webhookId,omitempty"`
	MessageID *Snowflake            `json:"messageId,omitempty"`
	Type      WebhookEventType      `json:"type"`
	Error     *string               `json:"error,omitempty"`
	Status    WebhookDeliveryStatus `json:"status"`
	URL       string                `json:"url"`
	Code      *int                  `json:"code,omitempty"`
	Payload   any                   `json:"payload,omitempty"`
	Attempt   int                   `json:"attempt"`
	Latency   *int                  `json:"latency,omitempty"`
	RetryAt   *string               `json:"retryAt,omitempty"`
}

// APIWebhookLogListItem matches the list item shape exposed through @rewritetoday/types for webhook logs.
type APIWebhookLogListItem struct {
	ID        Snowflake             `json:"id"`
	CreatedAt string                `json:"createdAt"`
	WebhookID *Snowflake            `json:"webhookId,omitempty"`
	MessageID *Snowflake            `json:"messageId,omitempty"`
	Type      WebhookEventType      `json:"type"`
	Error     *string               `json:"error,omitempty"`
	Status    WebhookDeliveryStatus `json:"status"`
	URL       string                `json:"url"`
	Code      *int                  `json:"code,omitempty"`
	Attempt   int                   `json:"attempt"`
	Latency   *int                  `json:"latency,omitempty"`
	RetryAt   *string               `json:"retryAt,omitempty"`
}

// APIWebhookLogList is kept for compatibility with older exports.
type APIWebhookLogList struct {
	Data   []APIWebhookLog `json:"data"`
	Cursor *Cursor         `json:"cursor,omitempty"`
}

// APIHealth represents the public health payload.
type APIHealth struct {
	Uptime int `json:"uptime"`
}

// RESTCursorOptions configures cursor-based pagination.
type RESTCursorOptions struct {
	Limit  int       `json:"limit,omitempty"`
	After  Snowflake `json:"after,omitempty"`
	Before Snowflake `json:"before,omitempty"`
}

// RESTPatchUpdateProjectBody is the request body for project updates.
type RESTPatchUpdateProjectBody struct {
	Name string   `json:"name,omitempty"`
	Icon JSONNull `json:"icon,omitempty"`
}

// RESTPatchUpdateProjectData corresponds to PATCH /projects.
type RESTPatchUpdateProjectData = APIResponse[any]

// RESTGetHealthData corresponds to GET /health.
type RESTGetHealthData = APIResponse[APIHealth]

// RESTGetListWebhooksData corresponds to GET /webhooks.
type RESTGetListWebhooksData = APIResponse[[]APIWebhook]

// RESTGetListWebhooksQueryParams corresponds to webhook list query params.
type RESTGetListWebhooksQueryParams = RESTCursorOptions

// RESTGetWebhookData corresponds to GET /webhooks/:id.
type RESTGetWebhookData = APIResponse[APIWebhook]

// RESTPostCreateWebhookData corresponds to POST /webhooks.
type RESTPostCreateWebhookData = APIResponse[APICreatedWebhook]

// RESTPostCreateWebhookBody is the request body for webhook creation.
type RESTPostCreateWebhookBody struct {
	Name     string             `json:"name,omitempty"`
	Endpoint string             `json:"endpoint"`
	Events   []WebhookEventType `json:"events"`
	Secret   string             `json:"secret,omitempty"`
}

// RESTPatchUpdateWebhookData corresponds to PATCH /webhooks/:id.
type RESTPatchUpdateWebhookData = APIResponse[APIWebhook]

// RESTPatchUpdateWebhookBody is the request body for webhook updates.
type RESTPatchUpdateWebhookBody struct {
	Name     NullableString     `json:"name,omitempty"`
	Endpoint string             `json:"endpoint,omitempty"`
	Events   []WebhookEventType `json:"events,omitempty"`
	Secret   string             `json:"secret,omitempty"`
	Status   WebhookStatus      `json:"status,omitempty"`
}

// RESTDeleteWebhookData corresponds to DELETE /webhooks/:id.
type RESTDeleteWebhookData = APIResponse[any]

// RESTDeleteWebhooksBody is the request body for bulk webhook deletion.
type RESTDeleteWebhooksBody struct {
	IDs []Snowflake `json:"ids"`
}

// RESTDeleteWebhooksData corresponds to DELETE /webhooks.
type RESTDeleteWebhooksData = APIResponse[[]Snowflake]

// RESTGetListWebhookLogsQueryParams corresponds to webhook log list query params.
type RESTGetListWebhookLogsQueryParams struct {
	RESTCursorOptions
	Type   WebhookEventType `json:"type,omitempty"`
	Status WebhookStatus    `json:"status,omitempty"`
}

// RESTGetListWebhookLogsData corresponds to GET /webhooks/:id/logs.
type RESTGetListWebhookLogsData = APIResponse[[]APIWebhookLogListItem]

// RESTGetListTemplatesQueryParams corresponds to template list query params.
type RESTGetListTemplatesQueryParams struct {
	RESTCursorOptions
	With18n  *bool `json:"with18n,omitempty"`
	WithI18N *bool `json:"-"`
}

// RESTGetListTemplatesData corresponds to GET /templates.
type RESTGetListTemplatesData = APIResponse[[]APITemplate]

// RESTGetTemplateQueryParams is kept for compatibility with older exports.
type RESTGetTemplateQueryParams struct {
	With18n  *bool `json:"with18n,omitempty"`
	WithI18N *bool `json:"-"`
}

// RESTGetTemplateData corresponds to GET /templates/:id.
type RESTGetTemplateData = APIResponse[APITemplate]

// RESTPostCreateTemplateData corresponds to POST /templates.
type RESTPostCreateTemplateData = APIResponse[APICreatedTemplate]

// RESTPostCreateTemplateBody is the request body for template creation.
type RESTPostCreateTemplateBody struct {
	Name        string                 `json:"name"`
	I18N        map[CountryCode]string `json:"i18n,omitempty"`
	Content     string                 `json:"content"`
	Variables   []APITemplateVariable  `json:"variables"`
	Description string                 `json:"description,omitempty"`
}

// RESTPatchUpdateTemplateData corresponds to PATCH /templates/:id.
type RESTPatchUpdateTemplateData = APIResponse[any]

// RESTPatchUpdateTemplateBody is the request body for template updates.
type RESTPatchUpdateTemplateBody struct {
	I18N        map[CountryCode]string `json:"i18n,omitempty"`
	Content     string                 `json:"content,omitempty"`
	Variables   []APITemplateVariable  `json:"variables,omitempty"`
	Description string                 `json:"description,omitempty"`
}

// RESTDeleteTemplateData corresponds to DELETE /templates/:id.
type RESTDeleteTemplateData = APIResponse[any]

// RESTDeleteTemplatesBody is the request body for bulk template deletion.
type RESTDeleteTemplatesBody struct {
	IDs []Snowflake `json:"ids"`
}

// RESTDeleteTemplatesData corresponds to DELETE /templates.
type RESTDeleteTemplatesData = APIResponse[[]Snowflake]

// RESTPostSendMessageBody is the request body for message sending.
type RESTPostSendMessageBody struct {
	To           string                         `json:"to"`
	Tags         []APIMessageTag                `json:"tags,omitempty"`
	ScheduledAt  string                         `json:"scheduledAt,omitempty"`
	Segmentation *APIMessageSegmentationOptions `json:"segmentation,omitempty"`
	Content      string                         `json:"content,omitempty"`
	TemplateID   Snowflake                      `json:"templateId,omitempty"`
	Variables    map[string]string              `json:"variables,omitempty"`
}

// RESTPostSendMessageData corresponds to POST /messages.
type RESTPostSendMessageData = APIResponse[APICreatedMessage]

// RESTPostSendBatchMessagesBody is the request body for POST /messages/batch.
type RESTPostSendBatchMessagesBody []RESTPostSendMessageBody

// APIBatchMessagesResult is the response payload for POST /messages/batch.
type APIBatchMessagesResult struct {
	IDs []Snowflake `json:"ids"`
}

// RESTPostSendBatchMessagesData corresponds to POST /messages/batch.
type RESTPostSendBatchMessagesData = APIResponse[APIBatchMessagesResult]

// RESTPostCancelMessageBody is kept for compatibility with older SDK versions.
type RESTPostCancelMessageBody struct {
	ID Snowflake `json:"id"`
}

// RESTPostCancelMessageData corresponds to POST /messages/:id/cancel.
type RESTPostCancelMessageData = APIResponse[any]

// RESTGetMessageData corresponds to GET /messages/:id.
type RESTGetMessageData = APIResponse[APIMessageWithoutAnalysis]

// RESTGetListMessagesQueryParams corresponds to message list query params.
type RESTGetListMessagesQueryParams struct {
	RESTCursorOptions
	Status  MessageStatus `json:"status,omitempty"`
	Country CountryCode   `json:"country,omitempty"`
}

// RESTGetListMessagesData corresponds to GET /messages.
type RESTGetListMessagesData = APIResponse[[]APIMessageWithoutAnalysis]

// RESTPostSendOTPMessageBody is the request body for OTP creation.
type RESTPostSendOTPMessageBody struct {
	To        string `json:"to"`
	Prefix    string `json:"prefix,omitempty"`
	ExpiresIn int    `json:"expiresIn,omitempty"`
}

// RESTPostSendOTPMessageData corresponds to POST /otp.
type RESTPostSendOTPMessageData = APIResponse[APIOTPMessage]

// RESTPostVerifyOTPCodeBody is the request body for OTP verification.
type RESTPostVerifyOTPCodeBody struct {
	To   string `json:"to"`
	Code string `json:"code"`
}

// RESTPostVerifyOTPCodeData corresponds to POST /otp/:id/verify.
type RESTPostVerifyOTPCodeData = APIResponse[APIOTPVerification]

// RESTGetWebhookLogData corresponds to GET /logs/:id.
type RESTGetWebhookLogData = APIResponse[APIWebhookLog]

// APIAPIKey represents a project API key entity.
type APIAPIKey struct {
	ID          Snowflake     `json:"id"`
	Name        string        `json:"name"`
	ProjectID   Snowflake     `json:"projectId,omitempty"`
	Prefix      string        `json:"prefix"`
	Scopes      []APIKeyScope `json:"scopes"`
	LastUsedAt  *string       `json:"lastUsedAt,omitempty"`
	Description *string       `json:"description,omitempty"`
	CreatedAt   string        `json:"createdAt"`
}

// APICreatedAPIKey represents the one-time response payload from API key creation.
type APICreatedAPIKey struct {
	ID          Snowflake `json:"id"`
	Key         string    `json:"key"`
	CreatedAt   string    `json:"createdAt"`
	Description *string   `json:"description,omitempty"`
}

// RESTDeleteAPIKeyData corresponds to DELETE /api-keys/:id.
type RESTDeleteAPIKeyData = APIResponse[any]

// Compatibility aliases for previous Go SDK releases.
type APIOTPCreateResponseData = APIOTPMessage
type APIOTPVerifyResponseData = APIOTPVerification
type RESTPostCreateMessageBody = RESTPostSendMessageBody
type RESTPostCreateMessageData = RESTPostSendMessageData
type RESTPostCreateMessagesBatchBody = RESTPostSendBatchMessagesBody
type RESTPostCreateMessagesBatchData = RESTPostSendBatchMessagesData
type RESTPostCreateOTPBody = RESTPostSendOTPMessageBody
type RESTPostCreateOTPData = RESTPostSendOTPMessageData
type RESTPostVerifyOTPBody = RESTPostVerifyOTPCodeBody
type RESTPostVerifyOTPData = RESTPostVerifyOTPCodeData
type RESTGetLogData = RESTGetWebhookLogData

func (q *RESTGetListTemplatesQueryParams) with18n() *bool {
	if q == nil {
		return nil
	}
	if q.With18n != nil {
		return q.With18n
	}
	return q.WithI18N
}

func (q *RESTGetTemplateQueryParams) with18n() *bool {
	if q == nil {
		return nil
	}
	if q.With18n != nil {
		return q.With18n
	}
	return q.WithI18N
}

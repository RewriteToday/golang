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

// APIContact represents a stored contact.
type APIContact struct {
	ID        Snowflake      `json:"id"`
	CreatedAt string         `json:"createdAt"`
	UpdatedAt string         `json:"updatedAt"`
	Name      *string        `json:"name"`
	Phone     string         `json:"phone"`
	Country   CountryCode    `json:"country"`
	Channel   *MessageType   `json:"channel"`
	Tags      map[string]any `json:"tags"`
}

// APICreatedContact represents the contact payload returned by POST /contacts.
type APICreatedContact struct {
	ID        Snowflake   `json:"id"`
	Phone     string      `json:"phone"`
	Country   CountryCode `json:"country"`
	CreatedAt string      `json:"createdAt"`
}

// APISegment represents a stored segment.
type APISegment struct {
	ID            Snowflake `json:"id"`
	CreatedAt     string    `json:"createdAt"`
	UpdatedAt     string    `json:"updatedAt"`
	Name          string    `json:"name"`
	Color         *string   `json:"color"`
	Description   *string   `json:"description"`
	ContactsCount int       `json:"contactsCount"`
}

// APITemplate represents a message template.
type APITemplate struct {
	ID          Snowflake              `json:"id"`
	Name        string                 `json:"name"`
	Content     string                 `json:"content"`
	Description *string                `json:"description"`
	I18N        map[CountryCode]string `json:"i18n,omitempty"`
	Variables   []APITemplateVariable  `json:"variables"`
	Tags        []APITemplateTag       `json:"tags"`
	CreatedAt   string                 `json:"createdAt"`
}

// APICreatedTemplate represents the create-template response payload.
type APICreatedTemplate struct {
	ID        Snowflake `json:"id"`
	CreatedAt string    `json:"createdAt"`
}

// APITemplateTag represents a static template tag.
type APITemplateTag struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

// WebhookEventSelection is the selector accepted by webhook create/update endpoints.
type WebhookEventSelection = string

const (
	// WebhookAllEvents subscribes a webhook to every supported event.
	WebhookAllEvents WebhookEventSelection = "*"
)

// APIWebhookDelivery represents delivery settings persisted on a webhook.
type APIWebhookDelivery struct {
	Timeout int `json:"timeout"`
	Retries int `json:"retries"`
}

// APIWebhook represents a webhook endpoint configuration.
type APIWebhook struct {
	ID        Snowflake               `json:"id"`
	Name      *string                 `json:"name"`
	Secret    string                  `json:"secret"`
	Endpoint  string                  `json:"endpoint"`
	Events    []WebhookEventSelection `json:"events"`
	Status    WebhookStatus           `json:"status"`
	Timeout   int                     `json:"timeout"`
	Retries   int                     `json:"retries"`
	CreatedAt string                  `json:"createdAt"`
}

// APIWebhookSummary represents a webhook list item.
type APIWebhookSummary struct {
	ID        Snowflake               `json:"id"`
	Name      *string                 `json:"name"`
	Endpoint  string                  `json:"endpoint"`
	Events    []WebhookEventSelection `json:"events"`
	Status    WebhookStatus           `json:"status"`
	Timeout   int                     `json:"timeout"`
	Retries   int                     `json:"retries"`
	CreatedAt string                  `json:"createdAt"`
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
	Count  int                   `json:"count"`
	Single int                   `json:"single"`
	Concat int                   `json:"concat"`
	Reason MessageAnalysisReason `json:"reason"`
}

// APIMessageAnalysis describes the calculated message metrics.
type APIMessageAnalysis struct {
	Characters int                        `json:"characters"`
	Encoding   MessageAnalysisEncoding    `json:"encoding"`
	Segments   APIMessageAnalysisSegments `json:"segments"`
}

// MessageError describes a provider-side error.
type MessageError struct {
	Code    any    `json:"code,omitempty"`
	Message string `json:"message"`
}

// APIMessageError is kept as a compatibility alias for MessageError.
type APIMessageError = MessageError

// APIMessage represents a persisted message read model.
type APIMessage struct {
	ID           Snowflake          `json:"id"`
	CreatedAt    string             `json:"createdAt"`
	Analysis     APIMessageAnalysis `json:"analysis"`
	To           string             `json:"to"`
	From         *Snowflake         `json:"from"`
	ContactID    *Snowflake         `json:"contactId"`
	Type         MessageType        `json:"type"`
	Tags         []APIMessageTag    `json:"tags"`
	Status       MessageStatus      `json:"status"`
	Country      CountryCode        `json:"country"`
	Content      string             `json:"content"`
	Encoding     MessageEncoding    `json:"encoding"`
	TemplateID   *Snowflake         `json:"templateId"`
	DeliveredAt  *string            `json:"deliveredAt"`
	ScheduledAt  *string            `json:"scheduledAt"`
	IsPayAsYouGo bool               `json:"isPayAsYouGo"`
}

// APIMessageWithoutAnalysis is the read model returned by message GET/LIST endpoints.
type APIMessageWithoutAnalysis struct {
	ID           Snowflake       `json:"id"`
	CreatedAt    string          `json:"createdAt"`
	To           string          `json:"to"`
	From         *Snowflake      `json:"from"`
	ContactID    *Snowflake      `json:"contactId"`
	Type         MessageType     `json:"type"`
	Tags         []APIMessageTag `json:"tags"`
	Status       MessageStatus   `json:"status"`
	Country      CountryCode     `json:"country"`
	Content      string          `json:"content"`
	Encoding     MessageEncoding `json:"encoding"`
	TemplateID   *Snowflake      `json:"templateId"`
	DeliveredAt  *string         `json:"deliveredAt"`
	ScheduledAt  *string         `json:"scheduledAt"`
	IsPayAsYouGo bool            `json:"isPayAsYouGo"`
}

// APICreatedMessage represents the send-message response payload.
type APICreatedMessage struct {
	ID        Snowflake          `json:"id"`
	CreatedAt string             `json:"createdAt"`
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
	VerifiedAt string    `json:"verifiedAt"`
}

// WebhookOTPMetadata carries OTP metadata included in webhook payloads.
type WebhookOTPMetadata struct {
	Prefix    string `json:"prefix"`
	ExpiresAt string `json:"expiresAt"`
	ExpiresIn int    `json:"expiresIn"`
}

// APIWebhookOTPMetadata is kept as a compatibility alias.
type APIWebhookOTPMetadata = WebhookOTPMetadata

// WebhookMessagePayload represents the message payload embedded in webhook deliveries.
type WebhookMessagePayload struct {
	ID          Snowflake          `json:"id"`
	ProjectID   Snowflake          `json:"projectId"`
	To          string             `json:"to"`
	Contact     *string            `json:"contact"`
	ContactID   *Snowflake         `json:"contactId"`
	Tags        []APIMessageTag    `json:"tags"`
	Type        MessageType        `json:"type"`
	Status      MessageStatus      `json:"status"`
	Country     CountryCode        `json:"country"`
	Content     string             `json:"content"`
	Analysis    APIMessageAnalysis `json:"analysis"`
	TemplateID  *Snowflake         `json:"templateId"`
	ScheduledAt *string            `json:"scheduledAt"`
	DeliveredAt *string            `json:"deliveredAt"`
	Error       *MessageError      `json:"error"`
}

// APIWebhookEventData is kept as a compatibility shape for broad webhook payload handling.
type APIWebhookEventData struct {
	ID          Snowflake           `json:"id,omitempty"`
	IDs         []Snowflake         `json:"ids,omitempty"`
	ProjectID   Snowflake           `json:"projectId"`
	To          string              `json:"to,omitempty"`
	Contact     *string             `json:"contact,omitempty"`
	ContactID   *Snowflake          `json:"contactId,omitempty"`
	Tags        []APIMessageTag     `json:"tags,omitempty"`
	Type        MessageType         `json:"type,omitempty"`
	Status      MessageStatus       `json:"status,omitempty"`
	Country     CountryCode         `json:"country,omitempty"`
	Content     string              `json:"content,omitempty"`
	Analysis    *APIMessageAnalysis `json:"analysis,omitempty"`
	TemplateID  *Snowflake          `json:"templateId,omitempty"`
	ScheduledAt *string             `json:"scheduledAt,omitempty"`
	DeliveredAt *string             `json:"deliveredAt,omitempty"`
	Error       *MessageError       `json:"error,omitempty"`
	OTP         *WebhookOTPMetadata `json:"otp,omitempty"`
}

// APIWebhookEvent represents a webhook event payload.
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
	WebhookID *Snowflake            `json:"webhookId"`
	MessageID *Snowflake            `json:"messageId"`
	Type      WebhookEventType      `json:"type"`
	Error     *string               `json:"error"`
	Status    WebhookDeliveryStatus `json:"status"`
	URL       string                `json:"url"`
	Code      *int                  `json:"code"`
	Payload   map[string]any        `json:"payload"`
	Attempt   int                   `json:"attempt"`
	Latency   *int                  `json:"latency"`
	RetryAt   *string               `json:"retryAt"`
}

// APIWebhookLogSummary matches the list item shape returned by GET /webhooks/:id/logs.
type APIWebhookLogSummary struct {
	ID        Snowflake             `json:"id"`
	CreatedAt string                `json:"createdAt"`
	MessageID *Snowflake            `json:"messageId"`
	Type      WebhookEventType      `json:"type"`
	Error     *string               `json:"error"`
	Status    WebhookDeliveryStatus `json:"status"`
	URL       string                `json:"url"`
	Code      *int                  `json:"code"`
	Attempt   int                   `json:"attempt"`
	Latency   *int                  `json:"latency"`
	RetryAt   *string               `json:"retryAt"`
}

// APIWebhookLogListItem is kept as a compatibility alias.
type APIWebhookLogListItem = APIWebhookLogSummary

// APIWebhookLogList is kept for compatibility with older exports.
type APIWebhookLogList struct {
	Data   []APIWebhookLogSummary `json:"data"`
	Cursor *Cursor                `json:"cursor,omitempty"`
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

// RESTGetContactData corresponds to GET /contacts/:identifier.
type RESTGetContactData = APIResponse[APIContact]

// RESTGetListContactsData corresponds to GET /contacts.
type RESTGetListContactsData = APIResponse[[]APIContact]

// RESTGetListContactsQueryParams corresponds to contact list query params.
type RESTGetListContactsQueryParams = RESTCursorOptions

// RESTPostCreateContactData corresponds to POST /contacts.
type RESTPostCreateContactData = APIResponse[APICreatedContact]

// RESTPostCreateContactBody is the request body for contact creation.
type RESTPostCreateContactBody struct {
	Phone   string         `json:"phone"`
	Name    string         `json:"name,omitempty"`
	Channel MessageType    `json:"channel,omitempty"`
	Tags    map[string]any `json:"tags,omitempty"`
}

// RESTPatchUpdateContactData corresponds to PATCH /contacts/:id.
type RESTPatchUpdateContactData = APIResponse[any]

// RESTPatchUpdateContactBody is the request body for contact updates.
type RESTPatchUpdateContactBody struct {
	Phone   string         `json:"phone,omitempty"`
	Name    string         `json:"name,omitempty"`
	Channel MessageType    `json:"channel,omitempty"`
	Tags    map[string]any `json:"tags,omitempty"`
}

// RESTDeleteContactData corresponds to DELETE /contacts/:id.
type RESTDeleteContactData = APIResponse[any]

// RESTGetSegmentData corresponds to GET /segments/:id.
type RESTGetSegmentData = APIResponse[APISegment]

// RESTGetListSegmentsData corresponds to GET /segments.
type RESTGetListSegmentsData = APIResponse[[]APISegment]

// RESTGetListSegmentsQueryParams corresponds to segment list query params.
type RESTGetListSegmentsQueryParams = RESTCursorOptions

// RESTPostCreateSegmentData corresponds to POST /segments.
type RESTPostCreateSegmentData = APIResponse[APISegment]

// RESTPostCreateSegmentBody is the request body for segment creation.
type RESTPostCreateSegmentBody struct {
	Name        string         `json:"name"`
	Color       NullableString `json:"color,omitempty"`
	Description NullableString `json:"description,omitempty"`
}

// RESTPatchUpdateSegmentData corresponds to PATCH /segments/:id.
type RESTPatchUpdateSegmentData = APIResponse[any]

// RESTPatchUpdateSegmentBody is the request body for segment updates.
type RESTPatchUpdateSegmentBody struct {
	Name        string         `json:"name,omitempty"`
	Color       NullableString `json:"color,omitempty"`
	Description NullableString `json:"description,omitempty"`
}

// RESTDeleteSegmentData corresponds to DELETE /segments/:id.
type RESTDeleteSegmentData = APIResponse[any]

// RESTGetListSegmentContactsData corresponds to GET /segments/:id/contacts.
type RESTGetListSegmentContactsData = APIResponse[[]APIContact]

// RESTGetListSegmentContactsQueryParams corresponds to segment contact list query params.
type RESTGetListSegmentContactsQueryParams = RESTCursorOptions

// RESTPostAttachSegmentContactBody is the request body for POST /segments/:id/contacts.
type RESTPostAttachSegmentContactBody struct {
	ContactID Snowflake `json:"contactId"`
}

// RESTPostAttachSegmentContactData corresponds to POST /segments/:id/contacts.
type RESTPostAttachSegmentContactData = APIResponse[any]

// RESTDeleteDetachSegmentContactData corresponds to DELETE /segments/:id/contacts/:contactId.
type RESTDeleteDetachSegmentContactData = APIResponse[any]

// RESTGetListWebhooksData corresponds to GET /webhooks.
type RESTGetListWebhooksData = APIResponse[[]APIWebhookSummary]

// RESTGetListWebhooksQueryParams corresponds to webhook list query params.
type RESTGetListWebhooksQueryParams = RESTCursorOptions

// RESTGetWebhookData corresponds to GET /webhooks/:id.
type RESTGetWebhookData = APIResponse[APIWebhook]

// RESTPostCreateWebhookData corresponds to POST /webhooks.
type RESTPostCreateWebhookData = APIResponse[APICreatedWebhook]

// RESTWebhookDeliveryBody is the partial delivery config accepted by webhook create/update endpoints.
type RESTWebhookDeliveryBody struct {
	Timeout *int `json:"timeout,omitempty"`
	Retries *int `json:"retries,omitempty"`
}

// RESTPostCreateWebhookBody is the request body for webhook creation.
type RESTPostCreateWebhookBody struct {
	Name     string                   `json:"name,omitempty"`
	Endpoint string                   `json:"endpoint"`
	Events   []WebhookEventSelection  `json:"events"`
	Secret   string                   `json:"secret,omitempty"`
	Delivery *RESTWebhookDeliveryBody `json:"delivery,omitempty"`
}

// RESTPatchUpdateWebhookData corresponds to PATCH /webhooks/:id.
type RESTPatchUpdateWebhookData = APIResponse[any]

// RESTPatchUpdateWebhookBody is the request body for webhook updates.
type RESTPatchUpdateWebhookBody struct {
	Name     NullableString           `json:"name,omitempty"`
	Endpoint string                   `json:"endpoint,omitempty"`
	Events   []WebhookEventSelection  `json:"events,omitempty"`
	Secret   string                   `json:"secret,omitempty"`
	Status   WebhookStatus            `json:"status,omitempty"`
	Delivery *RESTWebhookDeliveryBody `json:"delivery,omitempty"`
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
	Type   WebhookEventType      `json:"type,omitempty"`
	Status WebhookDeliveryStatus `json:"status,omitempty"`
}

// RESTGetListWebhookLogsData corresponds to GET /webhooks/:id/logs.
type RESTGetListWebhookLogsData = APIResponse[[]APIWebhookLogSummary]

// RESTGetListTemplatesQueryParams corresponds to template list query params.
type RESTGetListTemplatesQueryParams struct {
	RESTCursorOptions
	WithI18n *bool `json:"withi18n,omitempty"`
	With18n  *bool `json:"-"`
	WithI18N *bool `json:"-"`
}

// RESTGetListTemplatesData corresponds to GET /templates.
type RESTGetListTemplatesData = APIResponse[[]APITemplate]

// RESTGetTemplateQueryParams corresponds to GET /templates/:identifier query params.
type RESTGetTemplateQueryParams struct {
	WithI18n *bool `json:"withi18n,omitempty"`
	With18n  *bool `json:"-"`
	WithI18N *bool `json:"-"`
}

// RESTGetTemplateData corresponds to GET /templates/:identifier.
type RESTGetTemplateData = APIResponse[APITemplate]

// RESTPostCreateTemplateData corresponds to POST /templates.
type RESTPostCreateTemplateData = APIResponse[APICreatedTemplate]

// RESTPostCreateTemplateBody is the request body for template creation.
type RESTPostCreateTemplateBody struct {
	Name        string                `json:"name"`
	Content     string                `json:"content"`
	Variables   []APITemplateVariable `json:"variables"`
	Description NullableString        `json:"description,omitempty"`
	Tags        []APITemplateTag      `json:"tags,omitempty"`
}

// RESTPatchUpdateTemplateData corresponds to PATCH /templates/:id.
type RESTPatchUpdateTemplateData = APIResponse[any]

// RESTPatchUpdateTemplateBody is the request body for template updates.
type RESTPatchUpdateTemplateBody struct {
	Content     string                `json:"content,omitempty"`
	Variables   []APITemplateVariable `json:"variables,omitempty"`
	Description NullableString        `json:"description,omitempty"`
	Tags        []APITemplateTag      `json:"tags,omitempty"`
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
	To           string                         `json:"to,omitempty"`
	Contact      string                         `json:"contact,omitempty"`
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

// APIBatchMessagesResult is kept for compatibility with older SDK versions.
type APIBatchMessagesResult struct {
	IDs []Snowflake `json:"ids"`
}

// RESTPostSendBatchMessagesData corresponds to POST /messages/batch.
type RESTPostSendBatchMessagesData = APIResponse[[]APICreatedMessage]

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

func (q *RESTGetListTemplatesQueryParams) withI18n() *bool {
	if q == nil {
		return nil
	}
	if q.WithI18n != nil {
		return q.WithI18n
	}
	if q.With18n != nil {
		return q.With18n
	}
	return q.WithI18N
}

func (q *RESTGetTemplateQueryParams) withI18n() *bool {
	if q == nil {
		return nil
	}
	if q.WithI18n != nil {
		return q.WithI18n
	}
	if q.With18n != nil {
		return q.With18n
	}
	return q.WithI18N
}

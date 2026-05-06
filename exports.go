package rewrite

import (
	"github.com/rewritetoday/golang/api"
	"github.com/rewritetoday/golang/resources"
	"github.com/rewritetoday/golang/rest"
)

const (
	// APIBaseURL is the canonical Rewrite API base URL.
	APIBaseURL = api.APIBaseURL
)

var (
	// Routes exposes typed route builders equivalent to the public API surface.
	Routes = api.Routes
)

// Low-level REST aliases.
type (
	RESTOptions          = rest.Options
	RetryOptions         = rest.RetryOptions
	FetchOptions         = rest.FetchOptions
	RetryCallbackOptions = rest.HandleErrorOptions
	RetryResponseMeta    = rest.ResponseMeta
	HTTPError            = rest.HTTPError
)

// Resource option aliases.
type (
	APIKeyManager             = resources.APIKeyManager
	ContactManager            = resources.ContactManager
	DeliveryManager           = resources.DeliveryManager
	LogManager                = resources.LogManager
	MessageManager            = resources.MessageManager
	OTPManager                = resources.OTPManager
	SegmentManager            = resources.SegmentManager
	TagManager                = resources.TagManager
	TemplateManager           = resources.TemplateManager
	WebhookManager            = resources.WebhookManager
	SendMessageOptions        = resources.SendMessageOptions
	SendBatchMessageOptions   = resources.SendBatchMessageOptions
	SendOTPMessageOptions     = resources.SendOTPMessageOptions
	VerifyOTPOptions          = resources.VerifyOTPOptions
	VerifyWebhookOptions      = resources.VerifyWebhookOptions
	CreateMessageOptions      = resources.CreateMessageOptions
	CreateMessageBatchOptions = resources.CreateMessageBatchOptions
	CreateOTPOptions          = resources.CreateOTPOptions
)

// API model aliases.
type (
	Snowflake                     = api.Snowflake
	CountryCode                   = api.CountryCode
	Cursor                        = api.Cursor
	APIError                      = api.APIError
	APIErrorResponse              = api.APIErrorResponse
	APIValidationError            = api.APIValidationError
	NullableString                = api.NullableString
	JSONNull                      = api.JSONNull
	APIAPIKey                     = api.APIAPIKey
	APICreatedAPIKey              = api.APICreatedAPIKey
	APIContact                    = api.APIContact
	APICreatedContact             = api.APICreatedContact
	APISegment                    = api.APISegment
	APITemplate                   = api.APITemplate
	APICreatedTemplate            = api.APICreatedTemplate
	APITemplateVariable           = api.APITemplateVariable
	APITemplateTag                = api.APITemplateTag
	APIWebhook                    = api.APIWebhook
	APICompactDelivery            = api.APICompactDelivery
	APIDeliverySummary            = api.APIDeliverySummary
	APIDelivery                   = api.APIDelivery
	APIWebhookSummary             = api.APIWebhookSummary
	APIWebhookDelivery            = api.APIWebhookDelivery
	APICreatedWebhook             = api.APICreatedWebhook
	APITag                        = api.APITag
	APICreatedTag                 = api.APICreatedTag
	APIRequestLogSummary          = api.APIRequestLogSummary
	APIRequestLog                 = api.APIRequestLog
	APIMessageTag                 = api.APIMessageTag
	APIMessageSegmentationOptions = api.APIMessageSegmentationOptions
	APIMessageAnalysisSegments    = api.APIMessageAnalysisSegments
	APIMessageAnalysis            = api.APIMessageAnalysis
	MessageError                  = api.MessageError
	APIMessageError               = api.APIMessageError
	APIMessage                    = api.APIMessage
	APIMessageWithoutAnalysis     = api.APIMessageWithoutAnalysis
	APICreatedMessage             = api.APICreatedMessage
	APIBatchMessagesResult        = api.APIBatchMessagesResult
	APIOTPMessage                 = api.APIOTPMessage
	APIOTPVerification            = api.APIOTPVerification
	APIOTPCreateResponseData      = api.APIOTPCreateResponseData
	APIOTPVerifyResponseData      = api.APIOTPVerifyResponseData
	APIWebhookOTPMetadata         = api.APIWebhookOTPMetadata
	APIWebhookEventData           = api.APIWebhookEventData
	APIWebhookEvent               = api.APIWebhookEvent
	APIWebhookLog                 = api.APIWebhookLog
	APIWebhookLogSummary          = api.APIWebhookLogSummary
	APIWebhookLogListItem         = api.APIWebhookLogListItem
	APIWebhookLogList             = api.APIWebhookLogList
	APIKeyScope                   = api.APIKeyScope
	MessageProvider               = api.MessageProvider
	MessageType                   = api.MessageType
	MessageEncoding               = api.MessageEncoding
	MessageAnalysisEncoding       = api.MessageAnalysisEncoding
	MessageStatus                 = api.MessageStatus
	MessageSegmentationMode       = api.MessageSegmentationMode
	MessageAnalysisReason         = api.MessageAnalysisReason
	OtpAttemptStatus              = api.OtpAttemptStatus
	WebhookEventSelection         = api.WebhookEventSelection
	WebhookEventType              = api.WebhookEventType
	WebhookStatus                 = api.WebhookStatus
	WebhookDeliveryStatus         = api.WebhookDeliveryStatus
	RequestLogSource              = api.RequestLogSource
	RESTCursorOptions             = api.RESTCursorOptions
)

// API response/body aliases.
type (
	RESTGetContactData                      = api.RESTGetContactData
	RESTGetListContactsData                 = api.RESTGetListContactsData
	RESTGetListContactsQueryParams          = api.RESTGetListContactsQueryParams
	RESTGetListAPIKeysQueryParams           = api.RESTGetListAPIKeysQueryParams
	RESTGetListAPIKeysData                  = api.RESTGetListAPIKeysData
	RESTGetAPIKeyData                       = api.RESTGetAPIKeyData
	RESTPostCreateAPIKeyBody                = api.RESTPostCreateAPIKeyBody
	RESTPostCreateAPIKeyData                = api.RESTPostCreateAPIKeyData
	RESTPatchUpdateAPIKeyBody               = api.RESTPatchUpdateAPIKeyBody
	RESTPatchUpdateAPIKeyData               = api.RESTPatchUpdateAPIKeyData
	RESTDeleteAPIKeysData                   = api.RESTDeleteAPIKeysData
	RESTGetListAPIKeyLogsQueryParams        = api.RESTGetListAPIKeyLogsQueryParams
	RESTGetListAPIKeyLogsData               = api.RESTGetListAPIKeyLogsData
	RESTPostCreateContactData               = api.RESTPostCreateContactData
	RESTPostCreateContactBody               = api.RESTPostCreateContactBody
	RESTPatchUpdateContactData              = api.RESTPatchUpdateContactData
	RESTPatchUpdateContactBody              = api.RESTPatchUpdateContactBody
	RESTDeleteContactData                   = api.RESTDeleteContactData
	RESTDeleteContactsData                  = api.RESTDeleteContactsData
	RESTPostBatchContactsBody               = api.RESTPostBatchContactsBody
	RESTPostBatchContactsData               = api.RESTPostBatchContactsData
	RESTPostAttachContactTagsBody           = api.RESTPostAttachContactTagsBody
	RESTPostAttachContactTagsData           = api.RESTPostAttachContactTagsData
	RESTDeleteDetachContactTagsBody         = api.RESTDeleteDetachContactTagsBody
	RESTDeleteDetachContactTagsData         = api.RESTDeleteDetachContactTagsData
	RESTGetSegmentData                      = api.RESTGetSegmentData
	RESTGetListSegmentsData                 = api.RESTGetListSegmentsData
	RESTGetListSegmentsQueryParams          = api.RESTGetListSegmentsQueryParams
	RESTPostCreateSegmentData               = api.RESTPostCreateSegmentData
	RESTPostCreateSegmentBody               = api.RESTPostCreateSegmentBody
	RESTPatchUpdateSegmentData              = api.RESTPatchUpdateSegmentData
	RESTPatchUpdateSegmentBody              = api.RESTPatchUpdateSegmentBody
	RESTDeleteSegmentData                   = api.RESTDeleteSegmentData
	RESTDeleteSegmentsData                  = api.RESTDeleteSegmentsData
	RESTGetListSegmentContactsData          = api.RESTGetListSegmentContactsData
	RESTGetListSegmentContactsQueryParams   = api.RESTGetListSegmentContactsQueryParams
	RESTPostAttachSegmentContactBody        = api.RESTPostAttachSegmentContactBody
	RESTPostAttachSegmentContactData        = api.RESTPostAttachSegmentContactData
	RESTPostAttachSegmentContactsBody       = api.RESTPostAttachSegmentContactsBody
	RESTPostAttachSegmentContactsData       = api.RESTPostAttachSegmentContactsData
	RESTPostDetachSegmentContactsBody       = api.RESTPostDetachSegmentContactsBody
	RESTPostDetachSegmentContactsData       = api.RESTPostDetachSegmentContactsData
	RESTDeleteDetachSegmentContactData      = api.RESTDeleteDetachSegmentContactData
	RESTGetListWebhooksData                 = api.RESTGetListWebhooksData
	RESTGetListWebhooksQueryParams          = api.RESTGetListWebhooksQueryParams
	RESTGetWebhookData                      = api.RESTGetWebhookData
	RESTPostCreateWebhookData               = api.RESTPostCreateWebhookData
	RESTWebhookDeliveryBody                 = api.RESTWebhookDeliveryBody
	RESTPostCreateWebhookBody               = api.RESTPostCreateWebhookBody
	RESTPatchUpdateWebhookData              = api.RESTPatchUpdateWebhookData
	RESTPatchUpdateWebhookBody              = api.RESTPatchUpdateWebhookBody
	RESTDeleteWebhookData                   = api.RESTDeleteWebhookData
	RESTDeleteWebhooksData                  = api.RESTDeleteWebhooksData
	RESTGetListWebhookLogsQueryParams       = api.RESTGetListWebhookLogsQueryParams
	RESTGetListWebhookLogsData              = api.RESTGetListWebhookLogsData
	RESTGetListWebhookDeliveriesQueryParams = api.RESTGetListWebhookDeliveriesQueryParams
	RESTGetListWebhookDeliveriesData        = api.RESTGetListWebhookDeliveriesData
	RESTGetListTemplatesQueryParams         = api.RESTGetListTemplatesQueryParams
	RESTGetListTemplatesData                = api.RESTGetListTemplatesData
	RESTGetTemplateQueryParams              = api.RESTGetTemplateQueryParams
	RESTGetTemplateData                     = api.RESTGetTemplateData
	RESTPostCreateTemplateData              = api.RESTPostCreateTemplateData
	RESTPostCreateTemplateBody              = api.RESTPostCreateTemplateBody
	RESTPatchUpdateTemplateData             = api.RESTPatchUpdateTemplateData
	RESTPatchUpdateTemplateBody             = api.RESTPatchUpdateTemplateBody
	RESTDeleteTemplateData                  = api.RESTDeleteTemplateData
	RESTDeleteTemplatesData                 = api.RESTDeleteTemplatesData
	RESTPostDuplicateTemplateBody           = api.RESTPostDuplicateTemplateBody
	RESTPostDuplicateTemplateData           = api.RESTPostDuplicateTemplateData
	RESTPostSendMessageBody                 = api.RESTPostSendMessageBody
	RESTPostSendMessageData                 = api.RESTPostSendMessageData
	RESTPostSendBatchMessagesBody           = api.RESTPostSendBatchMessagesBody
	RESTPostSendBatchMessagesData           = api.RESTPostSendBatchMessagesData
	RESTPostCreateMessageBody               = api.RESTPostCreateMessageBody
	RESTPostCreateMessageData               = api.RESTPostCreateMessageData
	RESTPostCreateMessagesBatchBody         = api.RESTPostCreateMessagesBatchBody
	RESTPostCreateMessagesBatchData         = api.RESTPostCreateMessagesBatchData
	RESTPostCancelMessageBody               = api.RESTPostCancelMessageBody
	RESTPostCancelMessageData               = api.RESTPostCancelMessageData
	RESTGetMessageData                      = api.RESTGetMessageData
	RESTGetListMessagesQueryParams          = api.RESTGetListMessagesQueryParams
	RESTGetListMessagesData                 = api.RESTGetListMessagesData
	RESTPostSendOTPMessageBody              = api.RESTPostSendOTPMessageBody
	RESTPostSendOTPMessageData              = api.RESTPostSendOTPMessageData
	RESTPostVerifyOTPCodeBody               = api.RESTPostVerifyOTPCodeBody
	RESTPostVerifyOTPCodeData               = api.RESTPostVerifyOTPCodeData
	RESTGetListTagsData                     = api.RESTGetListTagsData
	RESTGetTagData                          = api.RESTGetTagData
	RESTPostCreateTagBody                   = api.RESTPostCreateTagBody
	RESTPostCreateTagData                   = api.RESTPostCreateTagData
	RESTPatchUpdateTagBody                  = api.RESTPatchUpdateTagBody
	RESTPatchUpdateTagData                  = api.RESTPatchUpdateTagData
	RESTDeleteTagData                       = api.RESTDeleteTagData
	RESTGetListLogsQueryParams              = api.RESTGetListLogsQueryParams
	RESTGetListLogsData                     = api.RESTGetListLogsData
	RESTGetListDeliveriesQueryParams        = api.RESTGetListDeliveriesQueryParams
	RESTGetListDeliveriesData               = api.RESTGetListDeliveriesData
	RESTGetDeliveryData                     = api.RESTGetDeliveryData
	RESTPostCreateOTPBody                   = api.RESTPostCreateOTPBody
	RESTPostCreateOTPData                   = api.RESTPostCreateOTPData
	RESTPostVerifyOTPBody                   = api.RESTPostVerifyOTPBody
	RESTPostVerifyOTPData                   = api.RESTPostVerifyOTPData
	RESTGetWebhookLogData                   = api.RESTGetWebhookLogData
	RESTGetLogData                          = api.RESTGetLogData
	RESTDeleteAPIKeyData                    = api.RESTDeleteAPIKeyData
)

// Nullable helper constructors.
func NewNullableString(value string) NullableString {
	return api.NewNullableString(value)
}

// NullString creates an explicit JSON null string value.
func NullString() NullableString {
	return api.NullString()
}

// NullJSON creates an explicit JSON null marker.
func NullJSON() JSONNull {
	return api.NewNull()
}

// APIKey scope constants.
const (
	APIKeyScopeWildcard      = api.APIKeyScopeWildcard
	APIKeyScopeReadProject   = api.APIKeyScopeReadProject
	APIKeyScopeReadAPIKeys   = api.APIKeyScopeReadAPIKeys
	APIKeyScopeWriteProject  = api.APIKeyScopeWriteProject
	APIKeyScopeReadWebhooks  = api.APIKeyScopeReadWebhooks
	APIKeyScopeWriteTemplate = api.APIKeyScopeWriteTemplate
	APIKeyScopeReadTemplates = api.APIKeyScopeReadTemplates
	APIKeyScopeWriteWebhooks = api.APIKeyScopeWriteWebhooks
	APIKeyScopeWriteMessages = api.APIKeyScopeWriteMessages
	APIKeyScopeReadMessages  = api.APIKeyScopeReadMessages
	APIKeyScopeReadLogs      = api.APIKeyScopeReadLogs
)

// Message constants.
const (
	MessageProviderComtele                         = api.MessageProviderComtele
	MessageTypeSMS                                 = api.MessageTypeSMS
	MessageTypeOTP                                 = api.MessageTypeOTP
	MessageEncodingGMS7                            = api.MessageEncodingGMS7
	MessageEncodingGSM7                            = api.MessageEncodingGSM7
	MessageEncodingUCS2                            = api.MessageEncodingUCS2
	MessageAnalysisEncodingGSM7                    = api.MessageAnalysisEncodingGSM7
	MessageAnalysisEncodingUCS2                    = api.MessageAnalysisEncodingUCS2
	MessageStatusSent                              = api.MessageStatusSent
	MessageStatusQueued                            = api.MessageStatusQueued
	MessageStatusFailed                            = api.MessageStatusFailed
	MessageStatusCanceled                          = api.MessageStatusCanceled
	MessageStatusScheduled                         = api.MessageStatusScheduled
	MessageStatusDelivered                         = api.MessageStatusDelivered
	MessageSegmentationModeFail                    = api.MessageSegmentationModeFail
	MessageSegmentationModeTrim                    = api.MessageSegmentationModeTrim
	MessageSegmentationModeSend                    = api.MessageSegmentationModeSend
	MessageAnalysisReasonFitsSingleSegment         = api.MessageAnalysisReasonFitsSingleSegment
	MessageAnalysisReasonSmartEncodingApplied      = api.MessageAnalysisReasonSmartEncodingApplied
	MessageAnalysisReasonExceedsSingleSegmentLimit = api.MessageAnalysisReasonExceedsSingleSegmentLimit
	MessageAnalysisReasonContainsNonGsm7Characters = api.MessageAnalysisReasonContainsNonGsm7Characters
	MessageAnalysisReasonContainsNonGSM7Characters = api.MessageAnalysisReasonContainsNonGSM7Characters
)

// OTP constants.
const (
	OtpAttemptStatusFailed     = api.OtpAttemptStatusFailed
	OtpAttemptStatusPending    = api.OtpAttemptStatusPending
	OtpAttemptStatusExpired    = api.OtpAttemptStatusExpired
	OtpAttemptStatusVerified   = api.OtpAttemptStatusVerified
	OtpAttemptStatusSuperseded = api.OtpAttemptStatusSuperseded
)

// Webhook constants.
const (
	WebhookAllEvents                 = api.WebhookAllEvents
	WebhookEventTypeSMSOTP           = api.WebhookEventTypeSMSOTP
	WebhookEventTypeMessageSent      = api.WebhookEventTypeMessageSent
	WebhookEventTypeMessageBatch     = api.WebhookEventTypeMessageBatch
	WebhookEventTypeMessageQueued    = api.WebhookEventTypeMessageQueued
	WebhookEventTypeMessageFailed    = api.WebhookEventTypeMessageFailed
	WebhookEventTypeMessageCanceled  = api.WebhookEventTypeMessageCanceled
	WebhookEventTypeMessageDelivered = api.WebhookEventTypeMessageDelivered
	WebhookEventTypeMessageScheduled = api.WebhookEventTypeMessageScheduled
	WebhookStatusActive              = api.WebhookStatusActive
	WebhookStatusInactive            = api.WebhookStatusInactive
	WebhookDeliveryStatusFailed      = api.WebhookDeliveryStatusFailed
	WebhookDeliveryStatusSuccess     = api.WebhookDeliveryStatusSuccess
)

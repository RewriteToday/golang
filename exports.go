package rewrite

import (
	"github.com/rewritetoday/golang/api"
	"github.com/rewritetoday/golang/resources"
	"github.com/rewritetoday/golang/rest"
)

const (
	// APIBaseURL is the canonical Rewrite API origin.
	APIBaseURL = api.APIBaseURL
)

var (
	// Routes exposes typed route builders equivalent to @rewritejs/types.
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
	CreateAPIKeyOptions   = resources.CreateAPIKeyOptions
	CreateTemplateOptions = resources.CreateTemplateOptions
	UpdateTemplateOptions = resources.UpdateTemplateOptions
	CreateWebhookOptions  = resources.CreateWebhookOptions
	UpdateWebhookOptions  = resources.UpdateWebhookOptions
)

// API model aliases.
type (
	Snowflake           = api.Snowflake
	APIAPIKey           = api.APIAPIKey
	APIProject          = api.APIProject
	APITemplate         = api.APITemplate
	APITemplateVariable = api.APITemplateVariable
	APIWebhook          = api.APIWebhook
	APIValidationError  = api.APIValidationError
	APIKeyScope         = api.APIKeyScope
	WebhookEventType    = api.WebhookEventType
	WebhookStatus       = api.WebhookStatus
	RESTCursorOptions   = api.RESTCursorOptions
	Null                = api.Null
)

// API response/body aliases.
type (
	RESTGetWebhookData              = api.RESTGetWebhookData
	RESTPostCreateWebhookData       = api.RESTPostCreateWebhookData
	RESTPostCreateWebhookBody       = api.RESTPostCreateWebhookBody
	RESTDeleteWebhookData           = api.RESTDeleteWebhookData
	RESTPatchUpdateWebhookData      = api.RESTPatchUpdateWebhookData
	RESTPatchUpdateWebhookBody      = api.RESTPatchUpdateWebhookBody
	RESTGetListWebhooksData         = api.RESTGetListWebhooksData
	RESTGetListWebhooksQueryParams  = api.RESTGetListWebhooksQueryParams
	RESTGetListTemplatesData        = api.RESTGetListTemplatesData
	RESTGetListTemplatesQueryParams = api.RESTGetListTemplatesQueryParams
	RESTPostCreateTemplateData      = api.RESTPostCreateTemplateData
	RESTPostCreateTemplateBody      = api.RESTPostCreateTemplateBody
	RESTPatchUpdateTemplateData     = api.RESTPatchUpdateTemplateData
	RESTPatchUpdateTemplateBody     = api.RESTPatchUpdateTemplateBody
	RESTDeleteTemplateData          = api.RESTDeleteTemplateData
	RESTGetTemplateData             = api.RESTGetTemplateData
	RESTGetListAPIKeysData          = api.RESTGetListAPIKeysData
	RESTGetListAPIKeysQueryParams   = api.RESTGetListAPIKeysQueryParams
	RESTPostCreateAPIKeyData        = api.RESTPostCreateAPIKeyData
	RESTPostCreateAPIKeyBody        = api.RESTPostCreateAPIKeyBody
	RESTDeleteAPIKeyData            = api.RESTDeleteAPIKeyData
	RESTPostCreateProjectData       = api.RESTPostCreateProjectData
	RESTPostCreateProjectBody       = api.RESTPostCreateProjectBody
	RESTPatchUpdateProjectData      = api.RESTPatchUpdateProjectData
	RESTPatchUpdateProjectBody      = api.RESTPatchUpdateProjectBody
	RESTDeleteProjectData           = api.RESTDeleteProjectData
	RESTGetProjectData              = api.RESTGetProjectData
)

// APIKey scope constants.
const (
	APIKeyScopeWildcard      = api.APIKeyScopeWildcard
	APIKeyScopeReadProject   = api.APIKeyScopeReadProject
	APIKeyScopeWriteProject  = api.APIKeyScopeWriteProject
	APIKeyScopeReadAPIKeys   = api.APIKeyScopeReadAPIKeys
	APIKeyScopeWriteTemplate = api.APIKeyScopeWriteTemplate
	APIKeyScopeReadTemplates = api.APIKeyScopeReadTemplates
	APIKeyScopeReadPayments  = api.APIKeyScopeReadPayments
	APIKeyScopeReadWebhooks  = api.APIKeyScopeReadWebhooks
	APIKeyScopeWriteWebhooks = api.APIKeyScopeWriteWebhooks
)

// Webhook event/status constants.
const (
	WebhookEventTypeSMSQueued    = api.WebhookEventTypeSMSQueued
	WebhookEventTypeSMSDelivered = api.WebhookEventTypeSMSDelivered
	WebhookEventTypeSMSScheduled = api.WebhookEventTypeSMSScheduled
	WebhookEventTypeSMSFailed    = api.WebhookEventTypeSMSFailed
	WebhookEventTypeSMSCanceled  = api.WebhookEventTypeSMSCanceled
	WebhookStatusActive          = api.WebhookStatusActive
	WebhookStatusInactive        = api.WebhookStatusInactive
)

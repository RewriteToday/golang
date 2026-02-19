package rest

import "time"

// Options configures the low-level REST client.
type Options struct {
	// BaseURL is the API origin. When empty, https://api.rewritetoday.com is used.
	BaseURL string
	// Auth is the Rewrite API secret used in the Bearer token.
	Auth string
	// Timeout is the default per-request timeout.
	Timeout time.Duration
	// Headers are merged into every outgoing request.
	Headers map[string]string
	// Retry configures retry behavior for retryable HTTP statuses.
	Retry *RetryOptions
}

// RetryOptions controls retry behavior for failed requests.
type RetryOptions struct {
	// Max is the maximum number of retries.
	Max int
	// Delay customizes wait duration before each retry attempt.
	Delay func(attempt int) time.Duration
	// OnRetry runs before each retry attempt.
	OnRetry func(options HandleErrorOptions)
}

// FetchOptions customizes an individual REST request.
type FetchOptions struct {
	// Headers overrides or adds headers for this request.
	Headers map[string]string
	// Timeout overrides the client default timeout for this request.
	Timeout time.Duration
	// Query appends query params. Supported values: string, map[string]string, url.Values, [][2]string.
	Query any

	method  string
	data    any
	hasData bool
}

// HandleErrorOptions are passed to RetryOptions.OnRetry.
type HandleErrorOptions struct {
	Method   string
	Route    string
	Attempt  int
	Response *ResponseMeta
	Options  FetchOptions
}

// ResponseMeta provides response metadata to retry callbacks.
type ResponseMeta struct {
	Status int
	URL    string
}

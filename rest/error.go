package rest

import "fmt"

// HTTPError is returned when the API responds with a non-success status.
type HTTPError struct {
	// Message is the parsed API error message.
	Message string
	// Status is the HTTP status code.
	Status int
	// URL is the requested URL.
	URL string
	// Method is the HTTP method used in the request.
	Method string
}

// Error implements the error interface.
func (e *HTTPError) Error() string {
	if e == nil {
		return ""
	}
	if e.Message != "" {
		return e.Message
	}
	return fmt.Sprintf("HTTPError(%d)", e.Status)
}

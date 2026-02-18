package rest

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"math"
	"math/rand"
	"net/url"
	"strings"
	"time"

	"github.com/go-resty/resty/v2"
)

const (
	defaultBaseURL = "https://api.rewritetoday.com"
	fiveSeconds    = 5 * time.Second
	baseDelay      = 300 * time.Millisecond
	maxDelay       = 10 * time.Second
)

var retryableStatus = map[int]struct{}{
	408: {},
	425: {},
	429: {},
	500: {},
	502: {},
	503: {},
	504: {},
}

// Client is the Rewrite low-level REST client.
type Client struct {
	options Options
	headers map[string]string
	client  *resty.Client
}

// New creates a REST client from an auth string or Options struct.
func New(options any) (*Client, error) {
	var resolved Options
	switch v := options.(type) {
	case string:
		resolved = Options{Auth: v}
	case Options:
		resolved = v
	case *Options:
		if v == nil {
			return nil, errors.New("Expected a string for the secret")
		}
		resolved = *v
	default:
		return nil, errors.New("Expected a string for the secret")
	}

	headers := make(map[string]string, len(resolved.Headers)+1)
	for k, v := range resolved.Headers {
		headers[k] = v
	}
	headers["Authorization"] = "Bearer " + resolved.Auth

	return &Client{
		options: resolved,
		headers: headers,
		client:  resty.New(),
	}, nil
}

// SetAuth updates the authorization token.
func (c *Client) SetAuth(authorization string) *Client {
	c.options.Auth = authorization
	c.headers["Authorization"] = "Bearer " + authorization
	return c
}

// Get executes a GET request.
func (c *Client) Get(ctx context.Context, route string, out any, options *FetchOptions) error {
	opts := cloneFetchOptions(options)
	opts.method = "GET"
	return c.fetch(ctx, route, out, opts, 0)
}

// Post executes a POST request.
func (c *Client) Post(ctx context.Context, route string, data any, out any, options *FetchOptions) error {
	opts := cloneFetchOptions(options)
	opts.method = "POST"
	opts.data = data
	opts.hasData = true
	return c.fetch(ctx, route, out, opts, 0)
}

// Delete executes a DELETE request.
func (c *Client) Delete(ctx context.Context, route string, out any, options *FetchOptions) error {
	opts := cloneFetchOptions(options)
	opts.method = "DELETE"
	return c.fetch(ctx, route, out, opts, 0)
}

// Put executes a PUT request.
func (c *Client) Put(ctx context.Context, route string, out any, options *FetchOptions) error {
	opts := cloneFetchOptions(options)
	opts.method = "PUT"
	return c.fetch(ctx, route, out, opts, 0)
}

// Patch executes a PATCH request.
func (c *Client) Patch(ctx context.Context, route string, data any, out any, options *FetchOptions) error {
	opts := cloneFetchOptions(options)
	opts.method = "PATCH"
	opts.data = data
	opts.hasData = true
	return c.fetch(ctx, route, out, opts, 0)
}

func (c *Client) fetch(ctx context.Context, route string, out any, options FetchOptions, attempt int) error {
	response, err := c.execute(ctx, route, options)
	if err != nil {
		return err
	}

	if response.IsError() {
		return c.handleError(ctx, route, out, options, attempt, response)
	}

	if out == nil || len(response.Body()) == 0 {
		return nil
	}

	var envelope struct {
		Data json.RawMessage `json:"data"`
	}
	if err := json.Unmarshal(response.Body(), &envelope); err != nil {
		return err
	}
	if len(envelope.Data) == 0 {
		return nil
	}

	if err := json.Unmarshal(envelope.Data, out); err != nil {
		return err
	}

	return nil
}

func (c *Client) execute(ctx context.Context, route string, options FetchOptions) (*resty.Response, error) {
	requestURL, err := CreateURL(route, options.Query, c.options.BaseURL)
	if err != nil {
		return nil, err
	}

	timeout := options.Timeout
	if timeout <= 0 {
		timeout = c.options.Timeout
	}
	if timeout <= 0 {
		timeout = fiveSeconds
	}

	requestCtx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	req := c.client.R().SetContext(requestCtx)
	for k, v := range c.headers {
		req.SetHeader(k, v)
	}
	for k, v := range options.Headers {
		req.SetHeader(k, v)
	}
	if options.hasData {
		req.SetBody(options.data)
	}

	return req.Execute(options.method, requestURL)
}

func (c *Client) handleError(ctx context.Context, route string, out any, options FetchOptions, attempt int, response *resty.Response) error {
	status := response.StatusCode()
	if !isRetryableStatus(status) {
		return &HTTPError{
			Message: readErrorMessage(response.Body()),
			Status:  status,
			URL:     response.Request.URL,
			Method:  options.method,
		}
	}

	maxRetries := 3
	if c.options.Retry != nil && c.options.Retry.Max > 0 {
		maxRetries = c.options.Retry.Max
	}
	if attempt >= maxRetries {
		return &HTTPError{
			Message: "Max retries reached",
			Status:  status,
			URL:     response.Request.URL,
			Method:  options.method,
		}
	}

	if c.options.Retry != nil && c.options.Retry.OnRetry != nil {
		c.options.Retry.OnRetry(HandleErrorOptions{
			Method:  options.method,
			Route:   route,
			Attempt: attempt,
			Response: &ResponseMeta{
				Status: status,
				URL:    response.Request.URL,
			},
			Options: options,
		})
	}

	delay := Backoff(attempt)
	if c.options.Retry != nil && c.options.Retry.Delay != nil {
		delay = c.options.Retry.Delay(attempt)
	}
	if err := sleepWithContext(ctx, delay); err != nil {
		return err
	}

	return c.fetch(ctx, route, out, options, attempt+1)
}

func cloneFetchOptions(in *FetchOptions) FetchOptions {
	if in == nil {
		return FetchOptions{}
	}
	out := *in
	if len(in.Headers) > 0 {
		out.Headers = make(map[string]string, len(in.Headers))
		for k, v := range in.Headers {
			out.Headers[k] = v
		}
	}
	return out
}

// CreateURL builds a URL in the same style used by @rewritejs/rest.
func CreateURL(route string, query any, baseURL string) (string, error) {
	if baseURL == "" {
		baseURL = defaultBaseURL
	}

	route = strings.TrimSpace(route)
	if route == "" {
		route = "/"
	}
	if !strings.HasPrefix(route, "/") {
		route = "/" + route
	}

	baseURL = strings.TrimSuffix(strings.TrimSpace(baseURL), "/")
	if strings.HasSuffix(baseURL, "/v1") {
		full := baseURL + route
		return appendQuery(full, query)
	}

	full := fmt.Sprintf("%s/v1%s", baseURL, route)
	return appendQuery(full, query)
}

func appendQuery(full string, query any) (string, error) {
	if query == nil {
		return full, nil
	}

	encoded, err := encodeQuery(query)
	if err != nil {
		return "", err
	}
	if encoded == "" {
		return full, nil
	}

	return full + "?" + encoded, nil
}

func encodeQuery(query any) (string, error) {
	switch q := query.(type) {
	case string:
		return strings.TrimPrefix(q, "?"), nil
	case map[string]string:
		values := make(url.Values, len(q))
		for k, v := range q {
			values.Set(k, v)
		}
		return values.Encode(), nil
	case url.Values:
		return q.Encode(), nil
	case [][2]string:
		values := make(url.Values)
		for _, pair := range q {
			values.Add(pair[0], pair[1])
		}
		return values.Encode(), nil
	default:
		return "", fmt.Errorf("unsupported query type %T", query)
	}
}

// Backoff computes retry delay using exponential backoff with jitter.
func Backoff(attempt int) time.Duration {
	exp := float64(baseDelay) * math.Pow(2, float64(attempt))
	if exp > float64(maxDelay) {
		exp = float64(maxDelay)
	}
	jitter := rand.Float64() * exp * 0.3
	return time.Duration(math.Floor(exp + jitter))
}

func isRetryableStatus(status int) bool {
	_, ok := retryableStatus[status]
	return ok
}

func sleepWithContext(ctx context.Context, delay time.Duration) error {
	if delay <= 0 {
		return nil
	}
	timer := time.NewTimer(delay)
	defer timer.Stop()

	select {
	case <-ctx.Done():
		return ctx.Err()
	case <-timer.C:
		return nil
	}
}

func readErrorMessage(body []byte) string {
	if len(body) == 0 {
		return "Request failed"
	}

	var parsed struct {
		Error any `json:"error"`
	}
	if err := json.Unmarshal(body, &parsed); err == nil && parsed.Error != nil {
		switch v := parsed.Error.(type) {
		case string:
			return v
		case map[string]any:
			if message, ok := v["message"].(string); ok && message != "" {
				return message
			}
		}
		return fmt.Sprintf("%v", parsed.Error)
	}

	return string(body)
}

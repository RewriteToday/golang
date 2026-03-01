package rewrite

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestNewFromOptionsAndAPIKeysList(t *testing.T) {
	var authHeader string
	var requestPath string
	var requestQuery string

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader = r.Header.Get("Authorization")
		requestPath = r.URL.Path
		requestQuery = r.URL.RawQuery
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"ok":true,"data":[],"cursor":{"persist":false}}`))
	}))
	defer server.Close()

	client, err := New(RewriteOptions{
		Secret: "rw_test",
		Rest: &RESTOptions{
			BaseURL: server.URL,
			Retry: &RetryOptions{
				Max:   1,
				Delay: func(int) time.Duration { return 0 },
			},
		},
	})
	if err != nil {
		t.Fatalf("unexpected constructor error: %v", err)
	}

	apiKeys, err := client.APIKeys.List(context.Background(), "123", nil)
	if err != nil {
		t.Fatalf("unexpected list error: %v", err)
	}

	if authHeader != "Bearer rw_test" {
		t.Fatalf("unexpected auth header: %q", authHeader)
	}
	if requestPath != "/v1/projects/123/api-keys" {
		t.Fatalf("unexpected path: %q", requestPath)
	}
	if requestQuery != "limit=15" {
		t.Fatalf("unexpected query: %q", requestQuery)
	}
	if !apiKeys.OK || len(apiKeys.Data) != 0 {
		t.Fatalf("unexpected payload: %+v", apiKeys)
	}
	if apiKeys.Cursor == nil || apiKeys.Cursor.Persist {
		t.Fatalf("unexpected cursor: %+v", apiKeys.Cursor)
	}
}

func TestConstructorSecretTypeError(t *testing.T) {
	_, err := New(123)
	if err == nil || err.Error() != "Expected a string for the secret" {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestAPIKeysListRouteDefaultLimit(t *testing.T) {
	route := Routes.APIKeys.List("abc", nil)
	if route != "/projects/abc/api-keys?limit=15" {
		t.Fatalf("unexpected route: %s", route)
	}
}

func TestTemplatesCreateDoesNotSendProjectInBody(t *testing.T) {
	var payload map[string]any

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/v1/projects/p1/templates" {
			t.Fatalf("unexpected path: %s", r.URL.Path)
		}
		if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
			t.Fatalf("decode body: %v", err)
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"ok":true,"data":{"id":"1","createdAt":"2026-02-19T20:01:09.000Z"}}`))
	}))
	defer server.Close()

	client, err := New(RewriteOptions{Secret: "rw", Rest: &RESTOptions{BaseURL: server.URL}})
	if err != nil {
		t.Fatalf("unexpected constructor error: %v", err)
	}

	_, err = client.Templates.Create(context.Background(), CreateTemplateOptions{
		Project: "p1",
		RESTPostCreateTemplateBody: RESTPostCreateTemplateBody{
			Name:    "welcome",
			Content: "Hello {{name}}",
			Variables: []APITemplateVariable{
				{Name: "name", Fallback: "customer"},
			},
		},
	})
	if err != nil {
		t.Fatalf("unexpected create error: %v", err)
	}

	if _, ok := payload["project"]; ok {
		t.Fatalf("did not expect project in body: %#v", payload["project"])
	}
	if payload["name"] != "welcome" {
		t.Fatalf("unexpected name in body: %#v", payload["name"])
	}
	if payload["content"] != "Hello {{name}}" {
		t.Fatalf("unexpected content in body: %#v", payload["content"])
	}
}

func TestAPIKeysListReturnsHTTPError(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(`{"error":"invalid project id"}`))
	}))
	defer server.Close()

	client, err := New(RewriteOptions{Secret: "rw", Rest: &RESTOptions{BaseURL: server.URL}})
	if err != nil {
		t.Fatalf("unexpected constructor error: %v", err)
	}

	_, err = client.APIKeys.List(context.Background(), "abc", nil)
	if err == nil {
		t.Fatal("expected error")
	}

	var httpErr *HTTPError
	if !errors.As(err, &httpErr) {
		t.Fatalf("expected HTTPError, got %T", err)
	}
	if httpErr.Status != http.StatusBadRequest {
		t.Fatalf("unexpected status: %d", httpErr.Status)
	}
	if httpErr.Message != "invalid project id" {
		t.Fatalf("unexpected message: %s", httpErr.Message)
	}
}

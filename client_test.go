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

func TestNewFromOptionsAndProjectsGet(t *testing.T) {
	var authHeader string
	var requestPath string

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader = r.Header.Get("Authorization")
		requestPath = r.URL.Path
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"data":{"ok":true,"data":{"id":"123","name":"Test","ownerId":"999"}}}`))
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

	project, err := client.Projects.Get(context.Background(), "123")
	if err != nil {
		t.Fatalf("unexpected get error: %v", err)
	}

	if authHeader != "Bearer rw_test" {
		t.Fatalf("unexpected auth header: %q", authHeader)
	}
	if requestPath != "/v1/projects/123" {
		t.Fatalf("unexpected path: %q", requestPath)
	}
	if !project.OK || project.Data.ID != "123" {
		t.Fatalf("unexpected payload: %+v", project)
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

func TestTemplatesCreateSendsProjectInBody(t *testing.T) {
	var payload map[string]any

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/v1/projects/p1/templates" {
			t.Fatalf("unexpected path: %s", r.URL.Path)
		}
		if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
			t.Fatalf("decode body: %v", err)
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"data":{"ok":true,"data":{"id":"1","name":"welcome","projectId":"p1","variables":[]}}}`))
	}))
	defer server.Close()

	client, err := New(RewriteOptions{Secret: "rw", Rest: &RESTOptions{BaseURL: server.URL}})
	if err != nil {
		t.Fatalf("unexpected constructor error: %v", err)
	}

	_, err = client.Templates.Create(context.Background(), CreateTemplateOptions{
		Project: "p1",
		RESTPostCreateTemplateBody: RESTPostCreateTemplateBody{
			Name:      "welcome",
			Variables: []APITemplateVariable{},
		},
	})
	if err != nil {
		t.Fatalf("unexpected create error: %v", err)
	}

	if payload["project"] != "p1" {
		t.Fatalf("expected project in body, got %#v", payload["project"])
	}
}

func TestProjectsGetReturnsHTTPError(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(`{"error":"invalid project id"}`))
	}))
	defer server.Close()

	client, err := New(RewriteOptions{Secret: "rw", Rest: &RESTOptions{BaseURL: server.URL}})
	if err != nil {
		t.Fatalf("unexpected constructor error: %v", err)
	}

	_, err = client.Projects.Get(context.Background(), "abc")
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

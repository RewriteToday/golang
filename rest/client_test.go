package rest

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/rewritetoday/golang/api"
)

func TestRetryOnRetryableStatus(t *testing.T) {
	attempts := 0
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		attempts++
		if attempts == 1 {
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte(`{"error":"temporary"}`))
			return
		}

		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"data":{"ok":true,"data":{"id":"1","name":"ok","ownerId":"2"}}}`))
	}))
	defer server.Close()

	client, err := New(Options{
		Auth:    "rw",
		BaseURL: server.URL,
		Retry: &RetryOptions{
			Max:   2,
			Delay: func(int) time.Duration { return 0 },
		},
	})
	if err != nil {
		t.Fatalf("unexpected constructor error: %v", err)
	}

	var out api.RESTGetProjectData
	if err := client.Get(context.Background(), "/projects/1", &out, nil); err != nil {
		t.Fatalf("unexpected get error: %v", err)
	}
	if !out.OK || out.Data.ID != "1" {
		t.Fatalf("unexpected output: %+v", out)
	}
	if attempts != 2 {
		t.Fatalf("expected 2 attempts, got %d", attempts)
	}
}

func TestCreateURL(t *testing.T) {
	url, err := CreateURL("/projects/1", map[string]string{"limit": "20"}, "https://example.com")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if url != "https://example.com/v1/projects/1?limit=20" {
		t.Fatalf("unexpected URL: %s", url)
	}
}

func TestCreateURLDefaultBaseURL(t *testing.T) {
	url, err := CreateURL("/projects/1", nil, "")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if url != "https://api.rewritetoday.com/v1/projects/1" {
		t.Fatalf("unexpected URL: %s", url)
	}
}

func TestCreateURLAvoidsDoubleV1(t *testing.T) {
	url, err := CreateURL("/projects/1", nil, "https://api.rewritetoday.com/v1")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if url != "https://api.rewritetoday.com/v1/projects/1" {
		t.Fatalf("unexpected URL: %s", url)
	}
}

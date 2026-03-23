package rest

import (
	"context"
	"io"
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
			_, _ = w.Write([]byte(`{"ok":false,"error":{"message":"temporary"}}`))
			return
		}

		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"ok":true,"data":{"id":"1","name":"Delivery","endpoint":"https://example.com/hook","events":["message.queued"],"status":"ACTIVE","createdAt":"2026-02-19T16:05:00.000Z"}}`))
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

	var out api.RESTGetWebhookData
	if err := client.Get(context.Background(), "/webhooks/1", &out, nil); err != nil {
		t.Fatalf("unexpected get error: %v", err)
	}
	if !out.OK || out.Data.ID != "1" || out.Data.Endpoint != "https://example.com/hook" {
		t.Fatalf("unexpected output: %+v", out)
	}
	if attempts != 2 {
		t.Fatalf("expected 2 attempts, got %d", attempts)
	}
}

func TestDeleteWithBody(t *testing.T) {
	var method string
	var body string

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		method = r.Method
		buf, _ := io.ReadAll(r.Body)
		body = string(buf)
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"ok":true,"data":["1"]}`))
	}))
	defer server.Close()

	client, err := New(Options{Auth: "rw", BaseURL: server.URL})
	if err != nil {
		t.Fatalf("unexpected constructor error: %v", err)
	}

	var out api.RESTDeleteTemplatesData
	if err := client.DeleteWithBody(context.Background(), "/templates", map[string]any{"ids": []string{"1"}}, &out, nil); err != nil {
		t.Fatalf("unexpected delete error: %v", err)
	}
	if method != http.MethodDelete {
		t.Fatalf("unexpected method: %s", method)
	}
	if body != "{\"ids\":[\"1\"]}" {
		t.Fatalf("unexpected body: %s", body)
	}
	if len(out.Data) != 1 || out.Data[0] != "1" {
		t.Fatalf("unexpected output: %+v", out)
	}
}

func TestCreateURL(t *testing.T) {
	url, err := CreateURL("/messages/1", map[string]string{"limit": "20"}, "https://example.com")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if url != "https://example.com/v1/messages/1?limit=20" {
		t.Fatalf("unexpected URL: %s", url)
	}
}

func TestCreateURLDefaultBaseURL(t *testing.T) {
	url, err := CreateURL("/health", nil, "")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if url != "https://api.rewritetoday.com/v1/health" {
		t.Fatalf("unexpected URL: %s", url)
	}
}

func TestCreateURLAvoidsDoubleV1(t *testing.T) {
	url, err := CreateURL("/health", nil, api.APIBaseURL)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if url != "https://api.rewritetoday.com/v1/health" {
		t.Fatalf("unexpected URL: %s", url)
	}
}

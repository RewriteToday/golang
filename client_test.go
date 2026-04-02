package rewrite

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestNewFromOptionsAndTemplatesList(t *testing.T) {
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

	templates, err := client.Templates.List(context.Background(), nil)
	if err != nil {
		t.Fatalf("unexpected list error: %v", err)
	}

	if authHeader != "Bearer rw_test" {
		t.Fatalf("unexpected auth header: %q", authHeader)
	}
	if requestPath != "/v1/templates" {
		t.Fatalf("unexpected path: %q", requestPath)
	}
	if requestQuery != "limit=15" {
		t.Fatalf("unexpected query: %q", requestQuery)
	}
	if !templates.OK || len(templates.Data) != 0 {
		t.Fatalf("unexpected payload: %+v", templates)
	}
	if templates.Cursor == nil || templates.Cursor.Persist {
		t.Fatalf("unexpected cursor: %+v", templates.Cursor)
	}
	if client.Contacts == nil || client.Segments == nil {
		t.Fatalf("expected contacts and segments managers to be initialized")
	}
}

func TestConstructorSecretTypeError(t *testing.T) {
	_, err := New(123)
	if err == nil || err.Error() != "Expected a string for the secret" {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestRoutesMatchPublicAPI(t *testing.T) {
	if route := Routes.Templates.List(nil); route != "/templates?limit=15" {
		t.Fatalf("unexpected templates list route: %s", route)
	}
	withI18n := true
	if route := Routes.Templates.List(&RESTGetListTemplatesQueryParams{WithI18n: &withI18n}); route != "/templates?limit=15&withi18n=true" {
		t.Fatalf("unexpected templates list route with withi18n: %s", route)
	}
	if route := Routes.Messages.List(nil); route != "/messages?limit=15" {
		t.Fatalf("unexpected messages list route: %s", route)
	}
	if route := Routes.Messages.Send(); route != "/messages" {
		t.Fatalf("unexpected message send route: %s", route)
	}
	if route := Routes.Messages.Get("abc"); route != "/messages/abc" {
		t.Fatalf("unexpected message get route: %s", route)
	}
	if route := Routes.Messages.Cancel("abc"); route != "/messages/abc/cancel" {
		t.Fatalf("unexpected message cancel route: %s", route)
	}
	if route := Routes.Contacts.List(nil); route != "/contacts?limit=15" {
		t.Fatalf("unexpected contacts list route: %s", route)
	}
	if route := Routes.Contacts.Get("abc"); route != "/contacts/abc" {
		t.Fatalf("unexpected contacts get route: %s", route)
	}
	if route := Routes.Segments.List(nil); route != "/segments?limit=15" {
		t.Fatalf("unexpected segments list route: %s", route)
	}
	if route := Routes.Segments.Contacts.List("abc", nil); route != "/segments/abc/contacts?limit=15" {
		t.Fatalf("unexpected segment contacts route: %s", route)
	}
	if route := Routes.Templates.Get("abc", &RESTGetTemplateQueryParams{WithI18n: &withI18n}); route != "/templates/abc?withi18n=true" {
		t.Fatalf("unexpected template get route with withi18n: %s", route)
	}
	if route := Routes.Webhooks.Logs("abc", nil); route != "/webhooks/abc/logs?limit=15" {
		t.Fatalf("unexpected webhook logs route: %s", route)
	}
	if route := Routes.Webhooks.Get("abc"); route != "/webhooks/abc" {
		t.Fatalf("unexpected webhook route: %s", route)
	}
}

func TestTemplatesCreateDoesNotSendHiddenFields(t *testing.T) {
	var payload map[string]any

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/v1/templates" {
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

	_, err = client.Templates.Create(context.Background(), RESTPostCreateTemplateBody{
		Name:        "welcome",
		Description: NewNullableString("Welcome message"),
		Content:     "Hello {{name}}",
		Variables: []APITemplateVariable{
			{Name: "name", Fallback: "customer"},
		},
		Tags: []APITemplateTag{
			{Name: "category", Value: "welcome"},
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
	if payload["description"] != "Welcome message" {
		t.Fatalf("unexpected description in body: %#v", payload["description"])
	}
	if _, ok := payload["i18n"]; ok {
		t.Fatalf("did not expect i18n in body: %#v", payload["i18n"])
	}
}

func TestTemplatesListSupportsCurrentAndLegacyI18nFields(t *testing.T) {
	var requestQuery string

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestQuery = r.URL.RawQuery
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"ok":true,"data":[],"cursor":{"persist":false}}`))
	}))
	defer server.Close()

	client, err := New(RewriteOptions{Secret: "rw", Rest: &RESTOptions{BaseURL: server.URL}})
	if err != nil {
		t.Fatalf("unexpected constructor error: %v", err)
	}

	withI18n := true
	if _, err := client.Templates.List(context.Background(), &RESTGetListTemplatesQueryParams{
		WithI18n: &withI18n,
	}); err != nil {
		t.Fatalf("unexpected list error: %v", err)
	}
	if requestQuery != "limit=15&withi18n=true" {
		t.Fatalf("unexpected current query: %q", requestQuery)
	}

	if _, err := client.Templates.List(context.Background(), &RESTGetListTemplatesQueryParams{
		With18n: &withI18n,
	}); err != nil {
		t.Fatalf("unexpected legacy list error: %v", err)
	}
	if requestQuery != "limit=15&withi18n=true" {
		t.Fatalf("unexpected legacy with18n query: %q", requestQuery)
	}

	if _, err := client.Templates.List(context.Background(), &RESTGetListTemplatesQueryParams{
		WithI18N: &withI18n,
	}); err != nil {
		t.Fatalf("unexpected legacy i18n list error: %v", err)
	}
	if requestQuery != "limit=15&withi18n=true" {
		t.Fatalf("unexpected legacy i18n query: %q", requestQuery)
	}
}

func TestTemplatesUpdateUsesCurrentContract(t *testing.T) {
	var requestPath string
	var payload map[string]any

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestPath = r.URL.Path
		if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
			t.Fatalf("decode body: %v", err)
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"ok":true,"data":null}`))
	}))
	defer server.Close()

	client, err := New(RewriteOptions{Secret: "rw", Rest: &RESTOptions{BaseURL: server.URL}})
	if err != nil {
		t.Fatalf("unexpected constructor error: %v", err)
	}

	response, err := client.Templates.Update(context.Background(), "tmpl_123", RESTPatchUpdateTemplateBody{
		Content: "Hello {{name}}",
		Variables: []APITemplateVariable{
			{Name: "name", Fallback: "friend"},
		},
		Tags: []APITemplateTag{
			{Name: "category", Value: "transactional"},
		},
	})
	if err != nil {
		t.Fatalf("unexpected update error: %v", err)
	}

	if requestPath != "/v1/templates/tmpl_123" {
		t.Fatalf("unexpected path: %q", requestPath)
	}
	if _, ok := payload["name"]; ok {
		t.Fatalf("did not expect name in patch payload: %#v", payload["name"])
	}
	if response.OK != true {
		t.Fatalf("unexpected response: %+v", response)
	}
}

func TestTemplatesGetSupportsQueryOptions(t *testing.T) {
	var requestPath string
	var requestQuery string

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestPath = r.URL.Path
		requestQuery = r.URL.RawQuery
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"ok":true,"data":{"id":"1","name":"welcome","content":"Hello {{name}}","description":"Welcome copy","variables":[{"name":"name","fallback":"friend"}],"tags":[{"name":"category","value":"welcome"}],"createdAt":"2026-02-19T20:01:09.000Z"}}`))
	}))
	defer server.Close()

	client, err := New(RewriteOptions{Secret: "rw", Rest: &RESTOptions{BaseURL: server.URL}})
	if err != nil {
		t.Fatalf("unexpected constructor error: %v", err)
	}

	withI18n := true
	if _, err := client.Templates.Get(context.Background(), "welcome", &RESTGetTemplateQueryParams{
		WithI18n: &withI18n,
	}); err != nil {
		t.Fatalf("unexpected get error: %v", err)
	}

	if requestPath != "/v1/templates/welcome" {
		t.Fatalf("unexpected path: %q", requestPath)
	}
	if requestQuery != "withi18n=true" {
		t.Fatalf("unexpected query: %q", requestQuery)
	}
}

func TestMessagesSendSendsIdempotencyKey(t *testing.T) {
	var idempotencyKey string
	var requestPath string

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		idempotencyKey = r.Header.Get("Idempotency-Key")
		requestPath = r.URL.Path
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"ok":true,"data":{"id":"1","createdAt":"2026-02-19T20:01:09.000Z","analysis":{"characters":5,"encoding":"gsm7","segments":{"concat":153,"count":1,"reason":"fits","single":160}}}}`))
	}))
	defer server.Close()

	client, err := New(RewriteOptions{Secret: "rw", Rest: &RESTOptions{BaseURL: server.URL}})
	if err != nil {
		t.Fatalf("unexpected constructor error: %v", err)
	}

	response, err := client.Messages.Send(context.Background(), SendMessageOptions{
		IdempotencyKey: "msg-1",
		RESTPostSendMessageBody: RESTPostSendMessageBody{
			To:      "+5511999999999",
			Content: "hello",
		},
	})
	if err != nil {
		t.Fatalf("unexpected send error: %v", err)
	}

	if requestPath != "/v1/messages" {
		t.Fatalf("unexpected path: %q", requestPath)
	}
	if idempotencyKey != "msg-1" {
		t.Fatalf("unexpected idempotency key: %q", idempotencyKey)
	}
	if response.Data.ID != "1" {
		t.Fatalf("unexpected response: %+v", response)
	}
}

func TestMessagesCancelUsesMessageIDInRoute(t *testing.T) {
	var requestPath string
	var requestBody string

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestPath = r.URL.Path
		body, _ := io.ReadAll(r.Body)
		requestBody = string(body)
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"ok":true,"data":null}`))
	}))
	defer server.Close()

	client, err := New(RewriteOptions{Secret: "rw", Rest: &RESTOptions{BaseURL: server.URL}})
	if err != nil {
		t.Fatalf("unexpected constructor error: %v", err)
	}

	if _, err := client.Messages.Cancel(context.Background(), "msg_123"); err != nil {
		t.Fatalf("unexpected cancel error: %v", err)
	}

	if requestPath != "/v1/messages/msg_123/cancel" {
		t.Fatalf("unexpected path: %q", requestPath)
	}
	if requestBody != "" {
		t.Fatalf("unexpected body: %q", requestBody)
	}
}

func TestMessagesGetUsesNodeRouteBuilder(t *testing.T) {
	var requestPath string

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestPath = r.URL.Path
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"ok":true,"data":{"id":"1","createdAt":"2026-02-19T20:01:09.000Z","to":"+5511999999999","type":"SMS","tags":[],"status":"SENT","country":"br","content":"hello","encoding":"GMS7","isPayAsYouGo":false}}`))
	}))
	defer server.Close()

	client, err := New(RewriteOptions{Secret: "rw", Rest: &RESTOptions{BaseURL: server.URL}})
	if err != nil {
		t.Fatalf("unexpected constructor error: %v", err)
	}

	if _, err := client.Messages.Get(context.Background(), "msg_123"); err != nil {
		t.Fatalf("unexpected get error: %v", err)
	}

	if requestPath != "/v1/messages/msg_123" {
		t.Fatalf("unexpected path: %q", requestPath)
	}
}

func TestMessagesSendOmitsEmptyIdempotencyKey(t *testing.T) {
	var idempotencyKey string

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		idempotencyKey = r.Header.Get("Idempotency-Key")
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"ok":true,"data":{"id":"1","createdAt":"2026-02-19T20:01:09.000Z","analysis":{"characters":5,"encoding":"gsm7","segments":{"concat":153,"count":1,"reason":"fits","single":160}}}}`))
	}))
	defer server.Close()

	client, err := New(RewriteOptions{Secret: "rw", Rest: &RESTOptions{BaseURL: server.URL}})
	if err != nil {
		t.Fatalf("unexpected constructor error: %v", err)
	}

	if _, err := client.Messages.Send(context.Background(), SendMessageOptions{
		RESTPostSendMessageBody: RESTPostSendMessageBody{
			To:      "+5511999999999",
			Content: "hello",
		},
	}); err != nil {
		t.Fatalf("unexpected send error: %v", err)
	}

	if idempotencyKey != "" {
		t.Fatalf("unexpected idempotency key header: %q", idempotencyKey)
	}
}

func TestOTPVerifyUsesMessageIDInPath(t *testing.T) {
	var requestPath string

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestPath = r.URL.Path
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"ok":true,"data":{"id":"1","valid":true,"verifiedAt":"2026-02-19T20:01:09.000Z"}}`))
	}))
	defer server.Close()

	client, err := New(RewriteOptions{Secret: "rw", Rest: &RESTOptions{BaseURL: server.URL}})
	if err != nil {
		t.Fatalf("unexpected constructor error: %v", err)
	}

	_, err = client.OTP.Verify(context.Background(), VerifyOTPOptions{
		ID: "1",
		RESTPostVerifyOTPCodeBody: RESTPostVerifyOTPCodeBody{
			To:   "+5511999999999",
			Code: "123456",
		},
	})
	if err != nil {
		t.Fatalf("unexpected verify error: %v", err)
	}

	if requestPath != "/v1/otp/1/verify" {
		t.Fatalf("unexpected path: %q", requestPath)
	}
}

func TestAPIKeysDeleteReturnsHTTPError(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(`{"ok":false,"error":{"code":"INVALID","message":"invalid api key id"}}`))
	}))
	defer server.Close()

	client, err := New(RewriteOptions{Secret: "rw", Rest: &RESTOptions{BaseURL: server.URL}})
	if err != nil {
		t.Fatalf("unexpected constructor error: %v", err)
	}

	_, err = client.APIKeys.Delete(context.Background(), "abc")
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
	if httpErr.Message != "invalid api key id" {
		t.Fatalf("unexpected message: %s", httpErr.Message)
	}
}

func TestWebhookVerifyReturnsTrueForValidSignature(t *testing.T) {
	client, err := New("rw_test")
	if err != nil {
		t.Fatalf("unexpected constructor error: %v", err)
	}

	payload := `{"type":"message.sent"}`
	mac := hmac.New(sha256.New, []byte("secret"))
	mac.Write([]byte("msg_test.1710000000." + payload))
	signature := base64.StdEncoding.EncodeToString(mac.Sum(nil))

	valid, err := client.Webhooks.Verify(VerifyWebhookOptions{
		Secret:  "whsec_c2VjcmV0",
		Payload: payload,
		Headers: map[string]string{
			"svix-id":        "msg_test",
			"svix-timestamp": "1710000000",
			"svix-signature": "v1," + signature,
		},
	})
	if err != nil {
		t.Fatalf("unexpected verify error: %v", err)
	}
	if !valid {
		t.Fatal("expected valid signature")
	}
}

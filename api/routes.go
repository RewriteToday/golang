package api

import (
	"fmt"
	"net/url"
	"strconv"
	"strings"
)

const (
	// APIBaseURL is the canonical Rewrite API origin from @rewritejs/types.
	APIBaseURL = "https://api.rewritetoday.com"
)

// Routes exposes helper builders for Rewrite API routes.
var Routes = RouteRegistry{
	Webhooks:  WebhookRoutes{},
	Templates: TemplateRoutes{},
	APIKeys:   APIKeyRoutes{},
	Projects:  ProjectRoutes{},
}

// RouteRegistry groups route builders by resource.
type RouteRegistry struct {
	Webhooks  WebhookRoutes
	Templates TemplateRoutes
	APIKeys   APIKeyRoutes
	Projects  ProjectRoutes
}

// WebhookRoutes builds webhook endpoints.
type WebhookRoutes struct{}

// TemplateRoutes builds template endpoints.
type TemplateRoutes struct{}

// APIKeyRoutes builds API key endpoints.
type APIKeyRoutes struct{}

// ProjectRoutes builds project endpoints.
type ProjectRoutes struct{}

// List returns GET /projects/:id/webhooks with cursor query.
func (WebhookRoutes) List(id string, options *RESTCursorOptions) string {
	return fmt.Sprintf("/projects/%s/webhooks?%s", id, createCursorQuery(options))
}

// Create returns POST /projects/:id/webhooks.
func (WebhookRoutes) Create(id string) string {
	return fmt.Sprintf("/projects/%s/webhooks", id)
}

// Update returns PATCH /projects/:id/webhooks/:webhookId.
func (WebhookRoutes) Update(id, webhookID string) string {
	return fmt.Sprintf("/projects/%s/webhooks/%s", id, webhookID)
}

// Delete returns DELETE /projects/:id/webhooks/:webhookId.
func (WebhookRoutes) Delete(id, webhookID string) string {
	return fmt.Sprintf("/projects/%s/webhooks/%s", id, webhookID)
}

// Get returns GET /projects/:id/webhooks/:webhookId.
func (WebhookRoutes) Get(id, webhookID string) string {
	return fmt.Sprintf("/projects/%s/webhooks/%s", id, webhookID)
}

// List returns GET /projects/:id/templates with cursor query.
func (TemplateRoutes) List(id string, options *RESTCursorOptions) string {
	return fmt.Sprintf("/projects/%s/templates?%s", id, createCursorQuery(options))
}

// Create returns POST /projects/:id/templates.
func (TemplateRoutes) Create(id string) string {
	return fmt.Sprintf("/projects/%s/templates", id)
}

// Update returns PATCH /projects/:id/templates/:templateId.
func (TemplateRoutes) Update(id, templateID string) string {
	return fmt.Sprintf("/projects/%s/templates/%s", id, templateID)
}

// Delete returns DELETE /projects/:id/templates/:templateId.
func (TemplateRoutes) Delete(id, templateID string) string {
	return fmt.Sprintf("/projects/%s/templates/%s", id, templateID)
}

// Get returns GET /projects/:id/templates/:templateId.
func (TemplateRoutes) Get(id, templateID string) string {
	return fmt.Sprintf("/projects/%s/templates/%s", id, templateID)
}

// List returns GET /projects/:id/api-keys with cursor query.
func (APIKeyRoutes) List(id string, options *RESTCursorOptions) string {
	return fmt.Sprintf("/projects/%s/api-keys?%s", id, createCursorQuery(options))
}

// Create returns POST /projects/:id/api-keys.
func (APIKeyRoutes) Create(id string) string {
	return fmt.Sprintf("/projects/%s/api-keys", id)
}

// Delete returns DELETE /projects/:id/api-keys/:apiKeyId.
func (APIKeyRoutes) Delete(id, apiKeyID string) string {
	return fmt.Sprintf("/projects/%s/api-keys/%s", id, apiKeyID)
}

// Create returns POST /projects.
func (ProjectRoutes) Create() string {
	return "/projects"
}

// Update returns PATCH /projects/:id.
func (ProjectRoutes) Update(id string) string {
	return fmt.Sprintf("/projects/%s", id)
}

// Delete returns DELETE /projects/:id.
func (ProjectRoutes) Delete(id string) string {
	return fmt.Sprintf("/projects/%s", id)
}

// Get returns GET /projects/:id.
func (ProjectRoutes) Get(id string) string {
	return fmt.Sprintf("/projects/%s", id)
}

func createCursorQuery(options *RESTCursorOptions) string {
	limit := 15
	if options != nil && options.Limit > 0 {
		limit = options.Limit
	}

	parts := []string{"limit=" + url.QueryEscape(strconv.Itoa(limit))}
	if options != nil {
		if options.After != "" {
			parts = append(parts, "after="+url.QueryEscape(string(options.After)))
		}
		if options.Before != "" {
			parts = append(parts, "before="+url.QueryEscape(string(options.Before)))
		}
	}

	return strings.Join(parts, "&")
}

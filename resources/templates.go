package resources

import (
	"context"

	"github.com/rewritetoday/golang/api"
)

// TemplateManager provides template resource operations.
type TemplateManager struct {
	Base
}

// Templates is kept as a compatibility alias for TemplateManager.
type Templates = TemplateManager

// Create creates a template.
func (r *TemplateManager) Create(ctx context.Context, body api.RESTPostCreateTemplateBody) (api.RESTPostCreateTemplateData, error) {
	var out api.RESTPostCreateTemplateData
	err := r.Rest.Post(ctx, api.Routes.Templates.Create(), body, &out, nil)
	return out, err
}

// Update updates a template by ID using the same payload and response shape exposed by the Node SDK.
func (r *TemplateManager) Update(ctx context.Context, id string, body api.RESTPostCreateTemplateBody) (api.RESTPostCreateTemplateData, error) {
	var out api.RESTPostCreateTemplateData
	err := r.Rest.Patch(ctx, api.Routes.Templates.Update(id), body, &out, nil)
	return out, err
}

// Delete deletes a template by ID.
func (r *TemplateManager) Delete(ctx context.Context, id string) (api.RESTDeleteTemplateData, error) {
	var out api.RESTDeleteTemplateData
	err := r.Rest.Delete(ctx, api.Routes.Templates.Delete(id), &out, nil)
	return out, err
}

// List lists templates.
func (r *TemplateManager) List(ctx context.Context, query *api.RESTGetListTemplatesQueryParams) (api.RESTGetListTemplatesData, error) {
	var out api.RESTGetListTemplatesData
	err := r.Rest.Get(ctx, api.Routes.Templates.List(query), &out, nil)
	return out, err
}

// Get fetches a template by ID.
func (r *TemplateManager) Get(ctx context.Context, id string) (api.RESTGetTemplateData, error) {
	var out api.RESTGetTemplateData
	err := r.Rest.Get(ctx, api.Routes.Templates.Get(id), &out, nil)
	return out, err
}

// GetWithQuery is a compatibility helper for the previous Go SDK signature.
func (r *TemplateManager) GetWithQuery(ctx context.Context, id string, _ *api.RESTGetTemplateQueryParams) (api.RESTGetTemplateData, error) {
	return r.Get(ctx, id)
}

package resources

import (
	"context"

	"github.com/rewritetoday/golang/api"
)

// Templates provides template resource operations.
type Templates struct {
	Base
}

// CreateTemplateOptions mirrors RESTPostCreateTemplateBody & { project: string } from the Node SDK.
type CreateTemplateOptions struct {
	Project string `json:"project"`
	api.RESTPostCreateTemplateBody
}

// UpdateTemplateOptions mirrors RESTPostCreateTemplateBody & { project: string } from the Node SDK.
type UpdateTemplateOptions struct {
	Project string `json:"project"`
	api.RESTPostCreateTemplateBody
}

// Create creates a template for a project.
func (r *Templates) Create(ctx context.Context, options CreateTemplateOptions) (api.RESTPostCreateTemplateData, error) {
	var out api.RESTPostCreateTemplateData
	err := r.Rest.Post(ctx, api.Routes.Templates.Create(options.Project), options, &out, nil)
	return out, err
}

// Update updates a template by ID.
func (r *Templates) Update(ctx context.Context, id string, options UpdateTemplateOptions) (api.RESTPostCreateTemplateData, error) {
	var out api.RESTPostCreateTemplateData
	err := r.Rest.Patch(ctx, api.Routes.Templates.Update(options.Project, id), options, &out, nil)
	return out, err
}

// Delete deletes a template by ID.
func (r *Templates) Delete(ctx context.Context, id, project string) error {
	return r.Rest.Delete(ctx, api.Routes.Templates.Delete(project, id), nil, nil)
}

// List lists templates for a project.
func (r *Templates) List(ctx context.Context, project string, query *api.RESTGetListTemplatesQueryParams) (api.RESTGetListTemplatesData, error) {
	var out api.RESTGetListTemplatesData
	err := r.Rest.Get(ctx, api.Routes.Templates.List(project, query), &out, nil)
	return out, err
}

// Get fetches a template by ID.
func (r *Templates) Get(ctx context.Context, id, project string) (api.RESTGetTemplateData, error) {
	var out api.RESTGetTemplateData
	err := r.Rest.Get(ctx, api.Routes.Templates.Get(project, id), &out, nil)
	return out, err
}

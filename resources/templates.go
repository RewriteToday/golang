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

func (r *TemplateManager) Create(ctx context.Context, body api.RESTPostCreateTemplateBody) (api.RESTPostCreateTemplateData, error) {
	var out api.RESTPostCreateTemplateData
	err := r.Rest.Post(ctx, api.Routes.Templates.Create(), body, &out, nil)
	return out, err
}

func (r *TemplateManager) Update(ctx context.Context, id string, body api.RESTPatchUpdateTemplateBody) (api.RESTPatchUpdateTemplateData, error) {
	var out api.RESTPatchUpdateTemplateData
	err := r.Rest.Patch(ctx, api.Routes.Templates.Update(id), body, &out, nil)
	return out, err
}

func (r *TemplateManager) Delete(ctx context.Context, id string) (api.RESTDeleteTemplateData, error) {
	var out api.RESTDeleteTemplateData
	err := r.Rest.Delete(ctx, api.Routes.Templates.Delete(id), &out, nil)
	return out, err
}

func (r *TemplateManager) Sweep(ctx context.Context, body api.RESTDeleteTemplatesBody) (api.RESTDeleteTemplatesData, error) {
	var out api.RESTDeleteTemplatesData
	err := r.Rest.DeleteWithBody(ctx, api.Routes.Templates.Sweep(), body, &out, nil)
	return out, err
}

func (r *TemplateManager) List(ctx context.Context, query *api.RESTGetListTemplatesQueryParams) (api.RESTGetListTemplatesData, error) {
	var out api.RESTGetListTemplatesData
	err := r.Rest.Get(ctx, api.Routes.Templates.List(query), &out, nil)
	return out, err
}

func (r *TemplateManager) Get(ctx context.Context, id string, query ...*api.RESTGetTemplateQueryParams) (api.RESTGetTemplateData, error) {
	var out api.RESTGetTemplateData
	err := r.Rest.Get(ctx, api.Routes.Templates.Get(id, firstTemplateQuery(query)), &out, nil)
	return out, err
}

func (r *TemplateManager) Duplicate(ctx context.Context, id string, body *api.RESTPostDuplicateTemplateBody) (api.RESTPostDuplicateTemplateData, error) {
	var out api.RESTPostDuplicateTemplateData
	err := r.Rest.Post(ctx, api.Routes.Templates.Duplicate(id), body, &out, nil)
	return out, err
}

func (r *TemplateManager) GetWithQuery(ctx context.Context, id string, query *api.RESTGetTemplateQueryParams) (api.RESTGetTemplateData, error) {
	return r.Get(ctx, id, query)
}

func firstTemplateQuery(query []*api.RESTGetTemplateQueryParams) *api.RESTGetTemplateQueryParams {
	if len(query) == 0 {
		return nil
	}
	return query[0]
}

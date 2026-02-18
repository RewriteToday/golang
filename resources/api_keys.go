package resources

import (
	"context"

	"github.com/rewritetoday/golang/api"
)

// APIKeys provides API key resource operations.
type APIKeys struct {
	Base
}

// CreateAPIKeyOptions mirrors RESTPostCreateAPIKeyBody & { project: string } from the Node SDK.
type CreateAPIKeyOptions struct {
	Project string `json:"project"`
	api.RESTPostCreateAPIKeyBody
}

// Create creates an API key for a project.
func (r *APIKeys) Create(ctx context.Context, options CreateAPIKeyOptions) (api.RESTPostCreateAPIKeyData, error) {
	var out api.RESTPostCreateAPIKeyData
	err := r.Rest.Post(ctx, api.Routes.APIKeys.Create(options.Project), options, &out, nil)
	return out, err
}

// Delete deletes an API key by ID.
func (r *APIKeys) Delete(ctx context.Context, id, project string) error {
	return r.Rest.Delete(ctx, api.Routes.APIKeys.Delete(project, id), nil, nil)
}

// List lists API keys for a project.
func (r *APIKeys) List(ctx context.Context, project string, query *api.RESTGetListAPIKeysQueryParams) (api.RESTGetListAPIKeysData, error) {
	var out api.RESTGetListAPIKeysData
	err := r.Rest.Get(ctx, api.Routes.APIKeys.List(project, query), &out, nil)
	return out, err
}

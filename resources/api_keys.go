package resources

import (
	"context"

	"github.com/rewritetoday/golang/api"
)

// APIKeyManager provides public API key resource operations.
type APIKeyManager struct {
	Base
}

// APIKeys is kept as a compatibility alias for APIKeyManager.
type APIKeys = APIKeyManager

func (r *APIKeyManager) List(ctx context.Context, query *api.RESTGetListAPIKeysQueryParams) (api.RESTGetListAPIKeysData, error) {
	var out api.RESTGetListAPIKeysData
	err := r.Rest.Get(ctx, api.Routes.APIKeys.List(query), &out, nil)
	return out, err
}

func (r *APIKeyManager) Create(ctx context.Context, body api.RESTPostCreateAPIKeyBody) (api.RESTPostCreateAPIKeyData, error) {
	var out api.RESTPostCreateAPIKeyData
	err := r.Rest.Post(ctx, api.Routes.APIKeys.Create(), body, &out, nil)
	return out, err
}

func (r *APIKeyManager) Sweep(ctx context.Context, body api.RESTDeleteWebhooksBody) (api.RESTDeleteAPIKeysData, error) {
	var out api.RESTDeleteAPIKeysData
	err := r.Rest.DeleteWithBody(ctx, api.Routes.APIKeys.Sweep(), body, &out, nil)
	return out, err
}

func (r *APIKeyManager) Get(ctx context.Context, id string) (api.RESTGetAPIKeyData, error) {
	var out api.RESTGetAPIKeyData
	err := r.Rest.Get(ctx, api.Routes.APIKeys.Get(id), &out, nil)
	return out, err
}

func (r *APIKeyManager) Update(ctx context.Context, id string, body api.RESTPatchUpdateAPIKeyBody) (api.RESTPatchUpdateAPIKeyData, error) {
	var out api.RESTPatchUpdateAPIKeyData
	err := r.Rest.Patch(ctx, api.Routes.APIKeys.Update(id), body, &out, nil)
	return out, err
}

// Delete deletes an API key by ID.
func (r *APIKeyManager) Delete(ctx context.Context, id string) (api.RESTDeleteAPIKeyData, error) {
	var out api.RESTDeleteAPIKeyData
	err := r.Rest.Delete(ctx, api.Routes.APIKeys.Delete(id), &out, nil)
	return out, err
}

func (r *APIKeyManager) Logs(ctx context.Context, id string, query *api.RESTGetListAPIKeyLogsQueryParams) (api.RESTGetListAPIKeyLogsData, error) {
	var out api.RESTGetListAPIKeyLogsData
	err := r.Rest.Get(ctx, api.Routes.APIKeys.Logs(id, query), &out, nil)
	return out, err
}

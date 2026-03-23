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

// Delete deletes an API key by ID.
func (r *APIKeyManager) Delete(ctx context.Context, id string) (api.RESTDeleteAPIKeyData, error) {
	var out api.RESTDeleteAPIKeyData
	err := r.Rest.Delete(ctx, api.Routes.APIKeys.Delete(id), &out, nil)
	return out, err
}

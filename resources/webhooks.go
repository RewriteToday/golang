package resources

import (
	"context"

	"github.com/rewritetoday/golang/api"
)

// Webhooks provides webhook resource operations.
type Webhooks struct {
	Base
}

// CreateWebhookOptions mirrors RESTPostCreateWebhookBody & { project: string } from the Node SDK.
type CreateWebhookOptions struct {
	Project string `json:"project"`
	api.RESTPostCreateWebhookBody
}

// UpdateWebhookOptions mirrors RESTPatchUpdateWebhookBody & { project: string } from the Node SDK.
type UpdateWebhookOptions struct {
	Project string `json:"project"`
	api.RESTPatchUpdateWebhookBody
}

// Create creates a webhook for a project.
func (r *Webhooks) Create(ctx context.Context, options CreateWebhookOptions) (api.RESTPostCreateWebhookData, error) {
	var out api.RESTPostCreateWebhookData
	err := r.Rest.Post(ctx, api.Routes.Webhooks.Create(options.Project), options, &out, nil)
	return out, err
}

// Update updates a webhook by ID.
func (r *Webhooks) Update(ctx context.Context, id string, options UpdateWebhookOptions) (api.RESTPatchUpdateWebhookData, error) {
	var out api.RESTPatchUpdateWebhookData
	err := r.Rest.Patch(ctx, api.Routes.Webhooks.Update(options.Project, id), options, &out, nil)
	return out, err
}

// Delete deletes a webhook by ID.
func (r *Webhooks) Delete(ctx context.Context, id, project string) error {
	return r.Rest.Delete(ctx, api.Routes.Webhooks.Delete(project, id), nil, nil)
}

// List lists webhooks for a project.
func (r *Webhooks) List(ctx context.Context, project string, query *api.RESTGetListWebhooksQueryParams) (api.RESTGetListWebhooksData, error) {
	var out api.RESTGetListWebhooksData
	err := r.Rest.Get(ctx, api.Routes.Webhooks.List(project, query), &out, nil)
	return out, err
}

// Get fetches a webhook by ID.
func (r *Webhooks) Get(ctx context.Context, id, project string) (api.RESTGetWebhookData, error) {
	var out api.RESTGetWebhookData
	err := r.Rest.Get(ctx, api.Routes.Webhooks.Get(project, id), &out, nil)
	return out, err
}

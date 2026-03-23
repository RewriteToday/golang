package resources

import (
	"context"

	"github.com/rewritetoday/golang/api"
)

// LogManager provides webhook delivery log operations.
type LogManager struct {
	Base
}

// Logs is kept as a compatibility alias for LogManager.
type Logs = LogManager

// List lists delivery logs for a webhook.
func (r *LogManager) List(ctx context.Context, webhookID string, query *api.RESTGetListWebhookLogsQueryParams) (api.RESTGetListWebhookLogsData, error) {
	var out api.RESTGetListWebhookLogsData
	err := r.Rest.Get(ctx, api.Routes.Webhooks.Logs(webhookID, query), &out, nil)
	return out, err
}

// Get fetches a webhook delivery log by ID.
func (r *LogManager) Get(ctx context.Context, id string) (api.RESTGetWebhookLogData, error) {
	var out api.RESTGetWebhookLogData
	err := r.Rest.Get(ctx, api.Routes.Logs.Get(id), &out, nil)
	return out, err
}

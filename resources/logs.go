package resources

import (
	"context"

	"github.com/rewritetoday/golang/api"
)

// LogManager provides request-log operations.
type LogManager struct {
	Base
}

// Logs is kept as a compatibility alias for LogManager.
type Logs = LogManager

func (r *LogManager) List(ctx context.Context, query *api.RESTGetListLogsQueryParams) (api.RESTGetListLogsData, error) {
	var out api.RESTGetListLogsData
	err := r.Rest.Get(ctx, api.Routes.Logs.List(query), &out, nil)
	return out, err
}

func (r *LogManager) Get(ctx context.Context, id string) (api.RESTGetLogData, error) {
	var out api.RESTGetLogData
	err := r.Rest.Get(ctx, api.Routes.Logs.Get(id), &out, nil)
	return out, err
}

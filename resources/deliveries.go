package resources

import (
	"context"

	"github.com/rewritetoday/golang/api"
)

// DeliveryManager provides webhook delivery operations.
type DeliveryManager struct {
	Base
}

// Deliveries is kept as a compatibility alias for DeliveryManager.
type Deliveries = DeliveryManager

func (r *DeliveryManager) List(ctx context.Context, query *api.RESTGetListDeliveriesQueryParams) (api.RESTGetListDeliveriesData, error) {
	var out api.RESTGetListDeliveriesData
	err := r.Rest.Get(ctx, api.Routes.Deliveries.List(query), &out, nil)
	return out, err
}

func (r *DeliveryManager) Get(ctx context.Context, id string) (api.RESTGetDeliveryData, error) {
	var out api.RESTGetDeliveryData
	err := r.Rest.Get(ctx, api.Routes.Deliveries.Get(id), &out, nil)
	return out, err
}

func (r *DeliveryManager) ByWebhook(ctx context.Context, id string, query *api.RESTGetListWebhookDeliveriesQueryParams) (api.RESTGetListWebhookDeliveriesData, error) {
	var out api.RESTGetListWebhookDeliveriesData
	err := r.Rest.Get(ctx, api.Routes.Deliveries.ByWebhook(id, query), &out, nil)
	return out, err
}

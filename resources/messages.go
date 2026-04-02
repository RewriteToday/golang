package resources

import (
	"context"

	"github.com/rewritetoday/golang/api"
	"github.com/rewritetoday/golang/rest"
)

// MessageManager provides message resource operations.
type MessageManager struct {
	Base
}

// Messages is kept as a compatibility alias for MessageManager.
type Messages = MessageManager

// SendMessageOptions carries message input and optional idempotency metadata.
type SendMessageOptions struct {
	IdempotencyKey string `json:"-"`
	api.RESTPostSendMessageBody
}

// CreateMessageOptions is kept as a compatibility alias for SendMessageOptions.
type CreateMessageOptions = SendMessageOptions

// SendBatchMessageOptions carries optional idempotency metadata for message batches.
type SendBatchMessageOptions struct {
	IdempotencyKey string `json:"-"`
}

// CreateMessageBatchOptions is kept for compatibility with the previous Go SDK.
type CreateMessageBatchOptions struct {
	IdempotencyKey string                            `json:"-"`
	Body           api.RESTPostSendBatchMessagesBody `json:"-"`
}

// Send sends a single message.
func (r *MessageManager) Send(ctx context.Context, options SendMessageOptions) (api.RESTPostSendMessageData, error) {
	var out api.RESTPostSendMessageData
	err := r.Rest.Post(
		ctx,
		api.Routes.Messages.Send(),
		options.RESTPostSendMessageBody,
		&out,
		withIdempotencyKey(options.IdempotencyKey),
	)
	return out, err
}

// Create is a compatibility alias for Send.
func (r *MessageManager) Create(ctx context.Context, options CreateMessageOptions) (api.RESTPostSendMessageData, error) {
	return r.Send(ctx, options)
}

// List lists messages.
func (r *MessageManager) List(ctx context.Context, query *api.RESTGetListMessagesQueryParams) (api.RESTGetListMessagesData, error) {
	var out api.RESTGetListMessagesData
	err := r.Rest.Get(ctx, api.Routes.Messages.List(query), &out, nil)
	return out, err
}

// Batch sends a message batch.
func (r *MessageManager) Batch(ctx context.Context, body api.RESTPostSendBatchMessagesBody, options SendBatchMessageOptions) (api.RESTPostSendBatchMessagesData, error) {
	var out api.RESTPostSendBatchMessagesData
	err := r.Rest.Post(
		ctx,
		api.Routes.Messages.Batch(),
		body,
		&out,
		withIdempotencyKey(options.IdempotencyKey),
	)
	return out, err
}

// CreateBatch is a compatibility helper for the previous Go SDK signature.
func (r *MessageManager) CreateBatch(ctx context.Context, options CreateMessageBatchOptions) (api.RESTPostSendBatchMessagesData, error) {
	return r.Batch(ctx, options.Body, SendBatchMessageOptions{
		IdempotencyKey: options.IdempotencyKey,
	})
}

// Cancel cancels a queued or scheduled message.
func (r *MessageManager) Cancel(ctx context.Context, id string) (api.RESTPostCancelMessageData, error) {
	var out api.RESTPostCancelMessageData
	err := r.Rest.Post(ctx, api.Routes.Messages.Cancel(id), nil, &out, nil)
	return out, err
}

// Get fetches a message by ID.
func (r *MessageManager) Get(ctx context.Context, id string) (api.RESTGetMessageData, error) {
	var out api.RESTGetMessageData
	err := r.Rest.Get(ctx, api.Routes.Messages.Get(id), &out, nil)
	return out, err
}

func withIdempotencyKey(value string) *rest.FetchOptions {
	if value == "" {
		return nil
	}

	return &rest.FetchOptions{
		Headers: map[string]string{
			"Idempotency-Key": value,
		},
	}
}

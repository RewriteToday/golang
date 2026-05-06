package resources

import (
	"context"

	"github.com/rewritetoday/golang/api"
)

// TagManager provides reusable tag operations.
type TagManager struct {
	Base
}

// Tags is kept as a compatibility alias for TagManager.
type Tags = TagManager

func (r *TagManager) List(ctx context.Context) (api.RESTGetListTagsData, error) {
	var out api.RESTGetListTagsData
	err := r.Rest.Get(ctx, api.Routes.Tags.List(), &out, nil)
	return out, err
}

func (r *TagManager) Create(ctx context.Context, body api.RESTPostCreateTagBody) (api.RESTPostCreateTagData, error) {
	var out api.RESTPostCreateTagData
	err := r.Rest.Post(ctx, api.Routes.Tags.Create(), body, &out, nil)
	return out, err
}

func (r *TagManager) Get(ctx context.Context, id string) (api.RESTGetTagData, error) {
	var out api.RESTGetTagData
	err := r.Rest.Get(ctx, api.Routes.Tags.Get(id), &out, nil)
	return out, err
}

func (r *TagManager) Update(ctx context.Context, id string, body api.RESTPatchUpdateTagBody) (api.RESTPatchUpdateTagData, error) {
	var out api.RESTPatchUpdateTagData
	err := r.Rest.Patch(ctx, api.Routes.Tags.Update(id), body, &out, nil)
	return out, err
}

func (r *TagManager) Delete(ctx context.Context, id string) (api.RESTDeleteTagData, error) {
	var out api.RESTDeleteTagData
	err := r.Rest.Delete(ctx, api.Routes.Tags.Delete(id), &out, nil)
	return out, err
}

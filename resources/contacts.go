package resources

import (
	"context"

	"github.com/rewritetoday/golang/api"
)

// ContactManager provides contact resource operations.
type ContactManager struct {
	Base
}

// Contacts is kept as a compatibility alias for ContactManager.
type Contacts = ContactManager

func (r *ContactManager) List(ctx context.Context, query *api.RESTGetListContactsQueryParams) (api.RESTGetListContactsData, error) {
	var out api.RESTGetListContactsData
	err := r.Rest.Get(ctx, api.Routes.Contacts.List(query), &out, nil)
	return out, err
}

func (r *ContactManager) Create(ctx context.Context, body api.RESTPostCreateContactBody) (api.RESTPostCreateContactData, error) {
	var out api.RESTPostCreateContactData
	err := r.Rest.Post(ctx, api.Routes.Contacts.Create(), body, &out, nil)
	return out, err
}

func (r *ContactManager) Sweep(ctx context.Context, body api.RESTDeleteWebhooksBody) (api.RESTDeleteContactsData, error) {
	var out api.RESTDeleteContactsData
	err := r.Rest.DeleteWithBody(ctx, api.Routes.Contacts.Sweep(), body, &out, nil)
	return out, err
}

func (r *ContactManager) Batch(ctx context.Context, body api.RESTPostBatchContactsBody) (api.RESTPostBatchContactsData, error) {
	var out api.RESTPostBatchContactsData
	err := r.Rest.Post(ctx, api.Routes.Contacts.Batch(), body, &out, nil)
	return out, err
}

func (r *ContactManager) Get(ctx context.Context, identifier string) (api.RESTGetContactData, error) {
	var out api.RESTGetContactData
	err := r.Rest.Get(ctx, api.Routes.Contacts.Get(identifier), &out, nil)
	return out, err
}

func (r *ContactManager) Update(ctx context.Context, id string, body api.RESTPatchUpdateContactBody) (api.RESTPatchUpdateContactData, error) {
	var out api.RESTPatchUpdateContactData
	err := r.Rest.Patch(ctx, api.Routes.Contacts.Update(id), body, &out, nil)
	return out, err
}

func (r *ContactManager) Delete(ctx context.Context, id string) (api.RESTDeleteContactData, error) {
	var out api.RESTDeleteContactData
	err := r.Rest.Delete(ctx, api.Routes.Contacts.Delete(id), &out, nil)
	return out, err
}

func (r *ContactManager) AddTags(ctx context.Context, id string, body api.RESTPostAttachContactTagsBody) (api.RESTPostAttachContactTagsData, error) {
	var out api.RESTPostAttachContactTagsData
	err := r.Rest.Post(ctx, api.Routes.Contacts.AddTags(id), body, &out, nil)
	return out, err
}

func (r *ContactManager) RemoveTags(ctx context.Context, id string, body api.RESTDeleteDetachContactTagsBody) (api.RESTDeleteDetachContactTagsData, error) {
	var out api.RESTDeleteDetachContactTagsData
	err := r.Rest.DeleteWithBody(ctx, api.Routes.Contacts.RemoveTags(id), body, &out, nil)
	return out, err
}

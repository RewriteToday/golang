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

// List lists contacts.
func (r *ContactManager) List(ctx context.Context, query *api.RESTGetListContactsQueryParams) (api.RESTGetListContactsData, error) {
	var out api.RESTGetListContactsData
	err := r.Rest.Get(ctx, api.Routes.Contacts.List(query), &out, nil)
	return out, err
}

// Create creates a contact.
func (r *ContactManager) Create(ctx context.Context, body api.RESTPostCreateContactBody) (api.RESTPostCreateContactData, error) {
	var out api.RESTPostCreateContactData
	err := r.Rest.Post(ctx, api.Routes.Contacts.Create(), body, &out, nil)
	return out, err
}

// Get fetches a contact by ID or phone number.
func (r *ContactManager) Get(ctx context.Context, identifier string) (api.RESTGetContactData, error) {
	var out api.RESTGetContactData
	err := r.Rest.Get(ctx, api.Routes.Contacts.Get(identifier), &out, nil)
	return out, err
}

// Update updates a contact by ID.
func (r *ContactManager) Update(ctx context.Context, id string, body api.RESTPatchUpdateContactBody) (api.RESTPatchUpdateContactData, error) {
	var out api.RESTPatchUpdateContactData
	err := r.Rest.Patch(ctx, api.Routes.Contacts.Update(id), body, &out, nil)
	return out, err
}

// Delete deletes a contact by ID.
func (r *ContactManager) Delete(ctx context.Context, id string) (api.RESTDeleteContactData, error) {
	var out api.RESTDeleteContactData
	err := r.Rest.Delete(ctx, api.Routes.Contacts.Delete(id), &out, nil)
	return out, err
}

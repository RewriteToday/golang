package resources

import (
	"context"

	"github.com/rewritetoday/golang/api"
)

// SegmentManager provides segment resource operations.
type SegmentManager struct {
	Base
}

// Segments is kept as a compatibility alias for SegmentManager.
type Segments = SegmentManager

// List lists segments.
func (r *SegmentManager) List(ctx context.Context, query *api.RESTGetListSegmentsQueryParams) (api.RESTGetListSegmentsData, error) {
	var out api.RESTGetListSegmentsData
	err := r.Rest.Get(ctx, api.Routes.Segments.List(query), &out, nil)
	return out, err
}

// Create creates a segment.
func (r *SegmentManager) Create(ctx context.Context, body api.RESTPostCreateSegmentBody) (api.RESTPostCreateSegmentData, error) {
	var out api.RESTPostCreateSegmentData
	err := r.Rest.Post(ctx, api.Routes.Segments.Create(), body, &out, nil)
	return out, err
}

// Get fetches a segment by ID.
func (r *SegmentManager) Get(ctx context.Context, id string) (api.RESTGetSegmentData, error) {
	var out api.RESTGetSegmentData
	err := r.Rest.Get(ctx, api.Routes.Segments.Get(id), &out, nil)
	return out, err
}

// Update updates a segment by ID.
func (r *SegmentManager) Update(ctx context.Context, id string, body api.RESTPatchUpdateSegmentBody) (api.RESTPatchUpdateSegmentData, error) {
	var out api.RESTPatchUpdateSegmentData
	err := r.Rest.Patch(ctx, api.Routes.Segments.Update(id), body, &out, nil)
	return out, err
}

// Delete deletes a segment by ID.
func (r *SegmentManager) Delete(ctx context.Context, id string) (api.RESTDeleteSegmentData, error) {
	var out api.RESTDeleteSegmentData
	err := r.Rest.Delete(ctx, api.Routes.Segments.Delete(id), &out, nil)
	return out, err
}

// ListContacts lists contacts attached to a segment.
func (r *SegmentManager) ListContacts(ctx context.Context, id string, query *api.RESTGetListSegmentContactsQueryParams) (api.RESTGetListSegmentContactsData, error) {
	var out api.RESTGetListSegmentContactsData
	err := r.Rest.Get(ctx, api.Routes.Segments.Contacts.List(id, query), &out, nil)
	return out, err
}

// AttachContact attaches a contact to a segment.
func (r *SegmentManager) AttachContact(ctx context.Context, id string, body api.RESTPostAttachSegmentContactBody) (api.RESTPostAttachSegmentContactData, error) {
	var out api.RESTPostAttachSegmentContactData
	err := r.Rest.Post(ctx, api.Routes.Segments.Contacts.Attach(id), body, &out, nil)
	return out, err
}

// DetachContact detaches a contact from a segment.
func (r *SegmentManager) DetachContact(ctx context.Context, id, contactID string) (api.RESTDeleteDetachSegmentContactData, error) {
	var out api.RESTDeleteDetachSegmentContactData
	err := r.Rest.Delete(ctx, api.Routes.Segments.Contacts.Detach(id, contactID), &out, nil)
	return out, err
}

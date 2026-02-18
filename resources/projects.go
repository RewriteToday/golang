package resources

import (
	"context"

	"github.com/rewritetoday/golang/api"
)

// Projects provides project resource operations.
type Projects struct {
	Base
}

// Create creates a new project.
func (r *Projects) Create(ctx context.Context, options api.RESTPostCreateProjectBody) (api.RESTPostCreateProjectData, error) {
	var out api.RESTPostCreateProjectData
	err := r.Rest.Post(ctx, api.Routes.Projects.Create(), options, &out, nil)
	return out, err
}

// Update updates a project by ID.
func (r *Projects) Update(ctx context.Context, id string, options api.RESTPatchUpdateProjectBody) (api.RESTPatchUpdateProjectData, error) {
	var out api.RESTPatchUpdateProjectData
	err := r.Rest.Patch(ctx, api.Routes.Projects.Update(id), options, &out, nil)
	return out, err
}

// Delete deletes a project by ID.
func (r *Projects) Delete(ctx context.Context, id string) error {
	return r.Rest.Delete(ctx, api.Routes.Projects.Delete(id), nil, nil)
}

// Get fetches a project by ID.
func (r *Projects) Get(ctx context.Context, id string) (api.RESTGetProjectData, error) {
	var out api.RESTGetProjectData
	err := r.Rest.Get(ctx, api.Routes.Projects.Get(id), &out, nil)
	return out, err
}

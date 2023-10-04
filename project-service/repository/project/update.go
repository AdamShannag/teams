package project

import (
	"context"
	"project-service/repository/ent"
	t "project-service/repository/ent/project"
	"project-service/resource/project"
	"time"
)

func (r repository) Update(ctx context.Context, request *project.UpdateRequest) (*ent.Project, error) {
	query := r.client.UpdateOneID(*request.ProjectId)

	query.SetNillableDescription(request.Description)
	setNillableStatus(request.Status, query)
	setNillableName(request.Name, query)

	return query.SetUpdatedAt(time.Now()).
		Save(ctx)
}

func setNillableStatus(status *t.Status, query *ent.ProjectUpdateOne) {
	if status != nil {
		query.SetStatus(*status)
	}
}

func setNillableName(name *string, query *ent.ProjectUpdateOne) {
	if name != nil {
		query.SetName(*name)
	}
}

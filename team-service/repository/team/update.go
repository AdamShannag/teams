package team

import (
	"context"
	"team-service/repository/ent"
	t "team-service/repository/ent/team"
	"team-service/resource/team"
	"time"
)

func (r repository) Update(ctx context.Context, request *team.UpdateRequest) (*ent.Team, error) {
	query := r.client.UpdateOneID(*request.TeamId)

	query.SetNillableDescription(request.Description)
	setNillableStatus(request.Status, query)
	setNillableName(request.Name, query)

	return query.SetUpdatedAt(time.Now()).
		Save(ctx)
}

func setNillableStatus(status *t.Status, query *ent.TeamUpdateOne) {
	if status != nil {
		query.SetStatus(*status)
	}
}

func setNillableName(name *string, query *ent.TeamUpdateOne) {
	if name != nil {
		query.SetName(*name)
	}
}

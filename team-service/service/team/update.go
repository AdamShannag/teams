package team

import (
	"context"
	mapper "team-service/mapper/team"
	"team-service/repository/ent"
	"team-service/resource/team"
	v "team-service/validation"
	"time"
)

type update struct {
	client     *ent.Client
	validation *v.Validation
	mapper     mapper.Mapper
}

func (s update) Update(ctx context.Context, request *team.UpdateRequest) (*team.Resource, error) {
	updated, err := s.client.Team.
		UpdateOneID(request.ID).
		SetName(request.Name).
		SetDescription(request.Description).
		SetStatus(request.Status).
		SetUpdatedAt(time.Now()).
		Save(ctx)

	if err != nil {
		return nil, err
	}

	resource := s.mapper.ToResource(updated)

	return &resource, nil
}

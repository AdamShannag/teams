package team

import (
	"context"
	"github.com/google/uuid"
	mapper "team-service/mapper/team"
	"team-service/repository/ent"
	t "team-service/repository/ent/team"
	"team-service/resource/team"
	v "team-service/validation"
)

type create struct {
	client     *ent.Client
	validation *v.Validation
	mapper     mapper.Mapper
}

func (s create) Create(ctx context.Context, request *team.Request) (*team.Resource, error) {
	if err := s.validation.Validate(*request, ctx); err != nil {
		return nil, err
	}

	created, err := s.client.Team.
		Create().
		SetID(uuid.New().String()).
		SetStatus(t.StatusNEW).
		SetName(request.Name).
		SetDescription(request.Description).
		SetCreatedBy(request.CreatedBy).
		Save(ctx)

	if err != nil {
		return nil, err
	}

	resource := s.mapper.ToResource(created)

	return &resource, nil
}

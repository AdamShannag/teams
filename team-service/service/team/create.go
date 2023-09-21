package team

import (
	"context"
	"github.com/google/uuid"
	mapper "team-service/mapper/team"
	"team-service/repository/ent"
	t "team-service/repository/ent/team"
	"team-service/resource/team"
	"team-service/validation/create_team_validation"
	"team-service/validation/violation"
)

type create struct {
	client     *ent.Client
	validation *create_team_validation.Validation
	mapper     mapper.Mapper
}

func (s create) Create(ctx context.Context, request *team.Request) (*team.Resource, []violation.Violation) {
	if violations := s.validation.Validate(*request, ctx); len(violations) != 0 {
		return nil, violations
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
		return nil, []violation.Violation{violation.FieldViolation("noField", err)}
	}

	resource := s.mapper.ToResource(created)

	return &resource, nil
}

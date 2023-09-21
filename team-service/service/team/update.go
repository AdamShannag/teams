package team

import (
	"context"
	mapper "team-service/mapper/team"
	"team-service/repository/ent"
	"team-service/resource/team"
	"team-service/validation/update_team_validation"
	"team-service/validation/violation"
	"time"
)

type update struct {
	client     *ent.Client
	validation *update_team_validation.Validation
	mapper     mapper.Mapper
}

func (s update) Update(ctx context.Context, request *team.UpdateRequest) (*team.Resource, []violation.Violation) {
	if violations := s.validation.Validate(*request, ctx); len(violations) != 0 {
		return nil, violations
	}

	updated, err := s.client.Team.
		UpdateOneID(request.TeamId).
		SetName(request.Name).
		SetDescription(request.Description).
		SetStatus(request.Status).
		SetUpdatedAt(time.Now()).
		Save(ctx)

	if err != nil {
		return nil, []violation.Violation{violation.FieldViolation("noField", err)}
	}

	resource := s.mapper.ToResource(updated)

	return &resource, nil
}

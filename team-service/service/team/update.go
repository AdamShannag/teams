package team

import (
	"context"
	"team-service/resource/team"
	"team-service/service/log"
	"team-service/validation/validator"
	"team-service/validation/violation"
)

type update struct {
	commonDependencies
	validator validator.Validator[team.UpdateRequest]
	log       log.Update
}

func (s update) Update(ctx context.Context, request *team.UpdateRequest) (*team.Resource, []violation.Violation) {
	if violations := s.validator.Validate(*request, ctx); len(violations) != 0 {
		return nil, violations
	}

	updated, err := s.repository.Update(ctx, request)

	if err != nil {

		s.log.Failed("team", err, *request.TeamId)
		return nil, []violation.Violation{violation.FieldViolation("noField", err)}
	}

	resource := s.mapper.ToResource(updated)

	s.log.Success("team", resource.ID)
	return &resource, nil
}

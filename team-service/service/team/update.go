package team

import (
	"context"
	"team-service/constant/message"
	"team-service/resource/team"
	"team-service/validation/validator"
	"team-service/validation/violation"
)

type upd struct {
	commonDependencies
	validator validator.Validator[team.UpdateRequest]
}

func (s upd) Update(ctx context.Context, request *team.UpdateRequest) (*team.Resource, []violation.Violation) {
	if violations := s.validator.Validate(*request, ctx); len(violations) != 0 {
		return nil, violations
	}

	updated, err := s.repository.Update(ctx, request)

	if err != nil {
		s.log.Error().Err(err).Msgf(message.UPDATED_FAILED, "team", request.TeamId)
		return nil, []violation.Violation{violation.FieldViolation("noField", err)}
	}

	resource := s.mapper.ToResource(updated)

	s.log.Info().Msgf(message.UPDATED_SUCCESSFULLY, "team", request.TeamId)
	return &resource, nil
}

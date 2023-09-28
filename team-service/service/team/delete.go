package team

import (
	"context"
	"team-service/resource/team"
	"team-service/service/log"
	"team-service/validation/validator"
	"team-service/validation/violation"
)

type delete struct {
	commonDependencies
	validator validator.Validator[team.DeleteRequest]
	log       log.Delete
}

func (s delete) Delete(ctx context.Context, request team.DeleteRequest) []violation.Violation {
	if violations := s.validator.Validate(request, ctx); len(violations) != 0 {
		return violations
	}

	err := s.repository.DeleteAll(ctx, request.TeamIds)

	if err != nil {
		s.log.Failed("teams", err, request.TeamIds...)
		return []violation.Violation{violation.FieldViolation("noField", err)}
	}
	s.log.Success("teams", request.TeamIds)
	return nil
}

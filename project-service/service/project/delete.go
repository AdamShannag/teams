package project

import (
	"context"
	"project-service/resource/project"
	"project-service/service/log"
	"project-service/validation/validator"
	"project-service/validation/violation"
)

type delete struct {
	commonDependencies
	validator validator.Validator[project.DeleteRequest]
	log       log.Delete
}

func (s delete) Delete(ctx context.Context, request project.DeleteRequest) []violation.Violation {
	if violations := s.validator.Validate(request, ctx); len(violations) != 0 {
		return violations
	}

	err := s.repository.DeleteAll(ctx, request.ProjectIds)

	if err != nil {
		s.log.Failed("projects", err, request.ProjectIds...)
		return []violation.Violation{violation.FieldViolation("noField", err)}
	}
	s.log.Success("projects", request.ProjectIds)
	return nil
}

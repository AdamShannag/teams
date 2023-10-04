package project

import (
	"context"
	"project-service/resource/project"
	"project-service/service/log"
	"project-service/validation/validator"
	"project-service/validation/violation"
)

type update struct {
	commonDependencies
	validator validator.Validator[project.UpdateRequest]
	log       log.Update
}

func (s update) Update(ctx context.Context, request *project.UpdateRequest) (*project.Resource, []violation.Violation) {
	if violations := s.validator.Validate(*request, ctx); len(violations) != 0 {
		return nil, violations
	}

	updated, err := s.repository.Update(ctx, request)

	if err != nil {

		s.log.Failed("project", err, *request.ProjectId)
		return nil, []violation.Violation{violation.FieldViolation("noField", err)}
	}

	resource := s.mapper.ToResource(updated)

	s.log.Success("project", resource.ID)
	return &resource, nil
}

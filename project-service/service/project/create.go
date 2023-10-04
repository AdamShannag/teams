package project

import (
	"context"
	"github.com/google/uuid"
	"project-service/repository/ent"
	p "project-service/repository/ent/project"
	"project-service/resource/project"
	"project-service/service/log"
	"project-service/validation/validator"
	"project-service/validation/violation"
)

type create struct {
	commonDependencies
	validator     validator.Validator[project.Request]
	userValidator validator.Validator[string]
	log           log.Create
}

func (s create) Create(ctx context.Context, request *project.Request, userId string) (*project.Resource, []violation.Violation) {
	if violations := s.userValidator.Validate(userId, ctx); len(violations) != 0 {
		return nil, violations
	}
	if violations := s.validator.Validate(*request, ctx); len(violations) != 0 {
		return nil, violations
	}

	entity := s.toEntity(request, userId)

	saved, err := s.repository.Save(ctx, &entity)

	if err != nil {
		s.log.Failed("project", err)
		return nil, []violation.Violation{violation.FieldViolation("noField", err)}
	}

	resource := s.mapper.ToResource(saved)

	s.log.Success("project", resource.ID)
	return &resource, nil
}

func (s create) toEntity(request *project.Request, userId string) ent.Project {
	entity := s.mapper.ToEntity(*request)
	entity.ID = uuid.New().String()
	entity.Status = p.StatusAVAILABLE
	entity.CreatedBy = userId
	return entity
}

package team

import (
	"context"
	"github.com/google/uuid"
	"team-service/constant/message"
	"team-service/repository/ent"
	t "team-service/repository/ent/team"
	"team-service/resource/team"
	"team-service/validation/validator"
	"team-service/validation/violation"
)

type create struct {
	commonDependencies
	validator     validator.Validator[team.Request]
	userValidator validator.Validator[string]
}

func (s create) Create(ctx context.Context, request *team.Request, userId string) (*team.Resource, []violation.Violation) {
	if violations := s.userValidator.Validate(userId, ctx); len(violations) != 0 {
		return nil, violations
	}
	if violations := s.validator.Validate(*request, ctx); len(violations) != 0 {
		return nil, violations
	}

	entity := s.toEntity(request, userId)

	saved, err := s.repository.Save(ctx, &entity)

	if err != nil {
		s.log.Error().Err(err).Msgf(message.CREATED_FAILED, "team")
		return nil, []violation.Violation{violation.FieldViolation("noField", err)}
	}

	resource := s.mapper.ToResource(saved)

	s.log.Info().Msgf(message.CREATED_SUCCESSFULLY, "team", resource.ID)
	return &resource, nil
}

func (s create) toEntity(request *team.Request, userId string) ent.Team {
	entity := s.mapper.ToEntity(*request)
	entity.ID = uuid.New().String()
	entity.Status = t.StatusAVAILABLE
	entity.CreatedBy = userId
	return entity
}

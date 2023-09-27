package member

import (
	"context"
	"team-service/repository/ent/member"
	memberrepo "team-service/repository/member"
	resource "team-service/resource/member"
	"team-service/validation/validator"
	"team-service/validation/violation"
)

type assign struct {
	repository    memberrepo.Repository
	validator     validator.Validator[resource.AssignRequest]
	userValidator validator.Validator[string]
}

func (h assign) Assign(ctx context.Context, request *resource.AssignRequest, userId string) []violation.Violation {
	if violations := h.userValidator.Validate(userId, ctx); len(violations) != 0 {
		return violations
	}
	if violations := h.validator.Validate(*request, ctx); len(violations) != 0 {
		return violations
	}

	resource := toResource(request, userId)

	var err error

	if request.Assign == nil || *request.Assign {
		err = h.repository.Assign(ctx, resource)
	} else {
		resource.Status = member.StatusFREE
		err = h.repository.UnAssign(ctx, resource)
	}

	if err != nil {
		return []violation.Violation{violation.FieldViolation("noField", err)}
	}

	return nil
}

func toResource(request *resource.AssignRequest, userId string) resource.AssignResource {
	return resource.AssignResource{
		Status:  member.StatusPENDING,
		TeamId:  request.TeamId,
		Members: request.Members,
		UserID:  userId,
	}
}

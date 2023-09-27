package member

import (
	"context"
	"team-service/repository/ent/member"
	memberrepo "team-service/repository/member"
	resource "team-service/resource/member"
	"team-service/validation/validator"
	"team-service/validation/violation"
)

type approval struct {
	repository    memberrepo.Repository
	validator     validator.Validator[resource.ApprovalRequest]
	userValidator validator.Validator[string]
}

func (h approval) AssignApproval(ctx context.Context, request *resource.ApprovalRequest, userId string) []violation.Violation {
	if violations := h.userValidator.Validate(userId, ctx); len(violations) != 0 {
		return violations
	}
	if violations := h.validator.Validate(*request, ctx); len(violations) != 0 {
		return violations
	}

	resource := resource.AssignResource{
		Status:  member.StatusIN_TEAM,
		Members: request.Members,
		UserID:  userId,
	}

	var err error

	if request.IsApprove() {
		err = h.repository.Approve(ctx, resource)
	} else {
		resource.Status = member.StatusFREE
		err = h.repository.Reject(ctx, resource)
	}

	if err != nil {
		return []violation.Violation{violation.FieldViolation("noField", err)}
	}

	return nil
}

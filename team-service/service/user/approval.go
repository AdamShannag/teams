package user

import (
	"context"
	"fmt"
	"team-service/repository/ent/user"
	resource "team-service/resource/user"
	"team-service/validation/validator"
	"team-service/validation/violation"
)

type approval struct {
	commonDependencies
	validator validator.Validator[resource.ApprovalRequest]
}

func (h approval) AssignApproval(ctx context.Context, request *resource.ApprovalRequest, userId string) []violation.Violation {
	if violations := h.userValidator.Validate(userId, ctx); len(violations) != 0 {
		return violations
	}
	if violations := h.validator.Validate(*request, ctx); len(violations) != 0 {
		return violations
	}

	resource := resource.AssignResource{
		Status: user.StatusIN_TEAM,
		Users:  request.Users,
		UserID: userId,
	}

	var err error

	if request.IsApprove() {
		err = h.repository.Approve(ctx, resource)
	} else {
		resource.Status = user.StatusFREE
		err = h.repository.Reject(ctx, resource)
	}

	if err != nil {
		h.log.Failed(failedApproveMessage(request, resource), err)
		return []violation.Violation{violation.FieldViolation("noField", err)}
	}

	h.log.Success(successApproveMessage(request, resource))
	return nil
}

func failedApproveMessage(request *resource.ApprovalRequest, resource resource.AssignResource) string {
	return fmt.Sprintf("Failed %s users %s to team", getPrefixApproveMessage(request.IsApprove()), resource.Users)
}

func successApproveMessage(request *resource.ApprovalRequest, resource resource.AssignResource) string {
	return fmt.Sprintf("Successfully %s users %s to team", getPrefixApproveMessage(request.IsApprove()), resource.Users)
}

func getPrefixApproveMessage(isApprove bool) string {
	if isApprove {
		return "approved"
	}
	return "rejected"
}
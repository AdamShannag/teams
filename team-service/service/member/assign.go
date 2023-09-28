package member

import (
	"context"
	"fmt"
	"team-service/repository/ent/member"
	resource "team-service/resource/member"
	"team-service/validation/validator"
	"team-service/validation/violation"
)

type assign struct {
	commonDependencies
	validator validator.Validator[resource.AssignRequest]
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

	if request.IsAssign() {
		err = h.repository.Assign(ctx, resource)
	} else {
		resource.Status = member.StatusFREE
		err = h.repository.UnAssign(ctx, resource)
	}

	if err != nil {
		h.log.Failed(failedMessage(request, resource), err)
		return []violation.Violation{violation.FieldViolation("noField", err)}
	}

	h.log.Success(successMessage(request, resource))
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

func failedMessage(request *resource.AssignRequest, resource resource.AssignResource) string {
	return fmt.Sprintf("Failed %s members %s to team [%s]", getPrefixMessage(request.IsAssign()), resource.Members, request.TeamId)
}

func successMessage(request *resource.AssignRequest, resource resource.AssignResource) string {
	return fmt.Sprintf("Successfully %s members %s to team [%s]", getPrefixMessage(request.IsAssign()), resource.Members, request.TeamId)
}

func getPrefixMessage(isAssign bool) string {
	if isAssign {
		return "assigned"
	}
	return "unassigned"
}

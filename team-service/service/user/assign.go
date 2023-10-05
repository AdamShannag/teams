package user

import (
	"context"
	"fmt"
	"team-service/repository/ent/user"
	resource "team-service/resource/user"
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
		resource.Status = user.StatusFREE
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
		Status: user.StatusPENDING,
		TeamId: request.TeamId,
		Users:  request.Users,
		UserID: userId,
	}
}

func failedMessage(request *resource.AssignRequest, resource resource.AssignResource) string {
	return fmt.Sprintf("Failed %s users %s to team [%s]", getPrefixMessage(request.IsAssign()), resource.Users, request.TeamId)
}

func successMessage(request *resource.AssignRequest, resource resource.AssignResource) string {
	return fmt.Sprintf("Successfully %s users %s to team [%s]", getPrefixMessage(request.IsAssign()), resource.Users, request.TeamId)
}

func getPrefixMessage(isAssign bool) string {
	if isAssign {
		return "assigned"
	}
	return "unassigned"
}

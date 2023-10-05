package assign

import (
	"context"
	"errors"
	"fmt"
	"team-service/repository/ent/team"
	"team-service/repository/ent/user"
	resource "team-service/resource/user"
	"team-service/validation/common"
	"team-service/validation/violation"
)

func (v *Validator) validateTeamId(teamId string, ctx context.Context) (violations []violation.Violation) {
	if err := common.IsEmptyString(teamId); err != nil {
		return []violation.Violation{violation.FieldViolation("teamId", err)}
	}
	if err := v.existById(teamId, ctx); err != nil {
		return []violation.Violation{violation.FieldViolation("teamId", err)}
	}
	return violations
}

func (v *Validator) validateUsers(request resource.AssignRequest, ctx context.Context) (violations []violation.Violation) {
	if err := common.IsEmpty(request.Users); err != nil {
		return []violation.Violation{violation.FieldViolation("users", err)}
	}
	for _, userId := range request.Users {
		if ok := v.isExistUser(userId, ctx); !ok {
			violations = append(violations, violation.FieldViolation("users", fmt.Errorf("user [%s] not found", userId)))
		}
		if ok, vo := v.checkUserAssignability(request.IsAssign(), userId, ctx); !ok {
			violations = append(violations, vo)
		}
	}
	return violations
}

func (v *Validator) checkUserAssignability(isAssign bool, userID string, ctx context.Context) (ok bool, vio violation.Violation) {
	isAssignable := v.assignableUser(userID, ctx)

	if isAssign {
		if !isAssignable {
			return false, violation.FieldViolation("users", fmt.Errorf("user [%s] is already assigned to a team", userID))
		}
	} else if isAssignable {
		return false, violation.FieldViolation("users", fmt.Errorf("user [%s] is not assigned to a team", userID))
	}
	return true, vio
}

func (v *Validator) existById(id string, ctx context.Context) error {
	ok, err := v.teamRepo.ExistByIdAndStatusNot(ctx, id, team.StatusDELETED)
	if err != nil {
		return err
	}
	if !ok {
		return errors.New(fmt.Sprintf("team id [%s] dose not exist", id))
	}
	return nil
}

func (v *Validator) assignableUser(userId string, ctx context.Context) (ok bool) {
	ok, _ = v.userRepo.ExistByIdAndStatus(ctx, userId, user.StatusFREE)
	return
}

func (v *Validator) isExistUser(userId string, ctx context.Context) (ok bool) {
	ok, _ = v.userRepo.ExistById(ctx, userId)
	return
}

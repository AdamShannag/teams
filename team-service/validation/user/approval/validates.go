package approval

import (
	"context"
	"fmt"
	"team-service/repository/ent/user"
	"team-service/validation/common"
	"team-service/validation/violation"
)

func (v *Validator) validateAssignApproval(userIds []string, ctx context.Context) (violations []violation.Violation) {
	if err := common.IsEmpty(userIds); err != nil {
		return []violation.Violation{violation.FieldViolation("users", err)}
	}
	for _, userId := range userIds {
		if ok := v.isExistUser(userId, ctx); !ok {
			violations = append(violations, violation.FieldViolation("users", fmt.Errorf("user [%s] not found", userId)))
		}
		if ok := v.isPendingUser(userId, ctx); !ok {
			violations = append(violations, violation.FieldViolation("users", fmt.Errorf("user [%s] is not pending assignation", userId)))
		}
	}
	return violations
}

func (v *Validator) isPendingUser(userId string, ctx context.Context) (ok bool) {
	ok, _ = v.repository.ExistByIdAndStatus(ctx, userId, user.StatusPENDING)
	return
}

func (v *Validator) isExistUser(userId string, ctx context.Context) (ok bool) {
	ok, _ = v.repository.ExistById(ctx, userId)
	return
}

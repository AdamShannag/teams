package user

import (
	"context"
	"fmt"
	"team-service/repository/ent/member"
	"team-service/validation/common"
	"team-service/validation/violation"
)

func (v *Validation) validateUser(userId string, ctx context.Context) (bool, violation.Violation) {
	if err := common.IsEmptyString(userId); err != nil {
		return false, violation.FieldViolation("userId", err)
	}
	if err := v.isExistUser(userId, ctx); err != nil {
		return false, violation.FieldViolation("userId", err)
	}
	return true, violation.Violation{}
}

func (v *Validation) isExistUser(userId string, ctx context.Context) error {
	if ok := v.client.Member.
		Query().
		Where(member.ID(userId)).
		ExistX(ctx); !ok {
		return fmt.Errorf("user [%s] dose not exist", userId)
	}
	return nil
}

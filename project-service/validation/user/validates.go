package user

import (
	"context"
	"fmt"
	"team-service/validation/common"
	"team-service/validation/violation"
)

func (v *Validator) validateUser(userId string, ctx context.Context) (bool, violation.Violation) {
	if err := common.IsEmptyString(userId); err != nil {
		return false, violation.FieldViolation("userId", err)
	}
	if err := v.isExistUser(userId, ctx); err != nil {
		return false, violation.FieldViolation("userId", err)
	}
	return true, violation.Violation{}
}

func (v *Validator) isExistUser(userId string, ctx context.Context) error {
	ok, err := v.repository.ExistById(ctx, userId)
	if err != nil {
		return err
	}
	if !ok {
		return fmt.Errorf("user [%s] dose not exist", userId)
	}
	return nil
}

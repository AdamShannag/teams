package create

import (
	"context"
	"errors"
	"fmt"
	"team-service/repository/ent/team"
	"team-service/validation/common"
	"team-service/validation/violation"
)

func (v *Validator) validateName(name string, ctx context.Context) (violations []violation.Violation) {
	if err := common.IsEmptyString(name); err != nil {
		return []violation.Violation{violation.FieldViolation("name", err)}
	}
	if err := v.existByName(name, ctx); err != nil {
		return []violation.Violation{violation.FieldViolation("name", err)}
	}
	return violations
}

func (v *Validator) existByName(name string, ctx context.Context) error {
	ok, err := v.repository.ExistByNameAndStatusNot(ctx, name, team.StatusDELETED)
	if err != nil {
		return err
	}
	if ok {
		return errors.New(fmt.Sprintf("team name [%s] is exist", name))
	}

	return nil
}

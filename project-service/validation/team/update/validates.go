package update

import (
	"context"
	"errors"
	"fmt"
	"team-service/repository/ent/team"
	"team-service/validation/common"
	"team-service/validation/violation"
)

func (v *Validator) validateTeamId(teamId *string, ctx context.Context) (violations []violation.Violation) {
	if err := common.IsNilString(teamId); err != nil {
		return []violation.Violation{violation.FieldViolation("teamId", err)}
	}
	if err := common.IsEmptyString(*teamId); err != nil {
		return []violation.Violation{violation.FieldViolation("teamId", err)}
	}
	if err := v.existById(*teamId, ctx); err != nil {
		return []violation.Violation{violation.FieldViolation("teamId", err)}
	}
	return violations
}

func (v *Validator) validateName(name string, ctx context.Context) (violations []violation.Violation) {
	if err := common.IsEmptyString(name); err != nil {
		return []violation.Violation{violation.FieldViolation("name", err)}
	}
	if err := v.existByName(name, ctx); err != nil {
		return []violation.Violation{violation.FieldViolation("name", err)}
	}
	return violations
}

func (v *Validator) validateStatus(status team.Status) (violations []violation.Violation) {
	if err := team.StatusValidator(status); err != nil {
		return []violation.Violation{violation.FieldViolation("status", err)}
	}
	return violations
}

func (v *Validator) existById(id string, ctx context.Context) error {
	ok, err := v.repository.ExistByIdAndStatusNot(ctx, id, team.StatusDELETED)
	if err != nil {
		return err
	}
	if !ok {
		return errors.New(fmt.Sprintf("project id [%s] dose not exist", id))
	}
	return nil
}

func (v *Validator) existByName(name string, ctx context.Context) error {
	ok, err := v.repository.ExistByNameAndStatusNot(ctx, name, team.StatusDELETED)
	if err != nil {
		return err
	}
	if ok {
		return errors.New(fmt.Sprintf("project name [%s] is exist", name))
	}
	return nil
}

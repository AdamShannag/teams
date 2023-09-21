package create_team_validation

import (
	"context"
	"errors"
	"fmt"
	"team-service/repository/ent/team"
	"team-service/validation/common"
	"team-service/validation/violation"
)

func (v *Validation) validateCreatedBy(createdBy string) []violation.Violation {
	if err := common.IsEmptyString(createdBy); err != nil {
		return []violation.Violation{violation.FieldViolation("createdBy", err)}
	}
	return []violation.Violation{}
}

func (v *Validation) validateName(name string, ctx context.Context) (violations []violation.Violation) {
	if err := common.IsEmptyString(name); err != nil {
		return []violation.Violation{violation.FieldViolation("name", err)}
	}
	if err := v.existByName(name, ctx); err != nil {
		return []violation.Violation{violation.FieldViolation("name", err)}
	}
	return violations
}

func (v *Validation) existByName(name string, ctx context.Context) error {
	if exist := v.client.Team.
		Query().
		Where(team.NameEQ(name)).
		ExistX(ctx); exist {
		return errors.New(fmt.Sprintf("team name [%s] is exist", name))
	}
	return nil
}

package update_team_validation

import (
	"context"
	"errors"
	"fmt"
	"team-service/repository/ent/team"
	"team-service/validation/common"
	"team-service/validation/violation"
)

func (v *Validation) validateTeamId(teamId string, ctx context.Context) (violations []violation.Violation) {
	if err := common.IsEmptyString(teamId); err != nil {
		return []violation.Violation{violation.FieldViolation("teamId", err)}
	}
	if err := v.existById(teamId, ctx); err != nil {
		return []violation.Violation{violation.FieldViolation("teamId", err)}
	}
	return violations
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

func (v *Validation) validateStatus(status team.Status) (violations []violation.Violation) {
	if err := team.StatusValidator(status); err != nil {
		return []violation.Violation{violation.FieldViolation("status", err)}
	}
	return violations
}

func (v *Validation) existById(id string, ctx context.Context) error {
	if exist := v.client.Team.
		Query().
		Where(team.ID(id)).
		ExistX(ctx); !exist {
		return errors.New(fmt.Sprintf("team id [%s] dose not exist", id))
	}
	return nil
}

func (v *Validation) existByName(name string, ctx context.Context) error {
	if exist := v.client.Team.
		Query().
		Where(team.Name(name)).
		ExistX(ctx); exist {
		return errors.New(fmt.Sprintf("team name [%s] is exist", name))
	}
	return nil
}

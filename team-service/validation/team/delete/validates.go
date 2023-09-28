package delete

import (
	"context"
	"errors"
	"fmt"
	"team-service/repository/ent/team"
	"team-service/validation/common"
	"team-service/validation/violation"
)

func (v *Validator) validateTeamIds(teamIds []string, ctx context.Context) (violations []violation.Violation) {
	if err := common.IsEmpty(teamIds); err != nil {
		return []violation.Violation{violation.FieldViolation("teamIds", err)}
	}
	for _, teamId := range teamIds {
		if err := v.existById(teamId, ctx); err != nil {
			violations = append(violations, violation.FieldViolation("teamId", err))
		}
	}
	return violations
}

func (v *Validator) existById(id string, ctx context.Context) error {
	ok, err := v.repository.ExistByIdAndStatusNot(ctx, id, team.StatusDELETED)
	if err != nil {
		return err
	}
	if !ok {
		return errors.New(fmt.Sprintf("team id [%s] dose not exist", id))
	}
	return nil
}

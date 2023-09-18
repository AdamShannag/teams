package validation

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-playground/validator"
	"team-service/repository/ent/team"
	team2 "team-service/resource/team"
)

func (v Validation) Validate(request team2.Request, ctx context.Context) error {
	s := validator.New()
	return s.Struct(request)
	//if err := v.IsEmpty(request.Name); err != nil {
	//	violations = append(violations, violation.FieldViolation("name", err))
	//	return
	//}
	//if err := v.ExistByName(request.Name, ctx); err != nil {
	//	violations = append(violations, violation.FieldViolation("name", err))
	//	return
	//}
	//return
}

func (v *Validation) IsEmpty(value string) error {
	if len(value) == 0 {
		return errors.New("not be empty")
	}
	return nil
}

func (v *Validation) ExistByName(name string, ctx context.Context) error {

	if exist := v.existByName(name, ctx); exist {
		return errors.New(fmt.Sprintf("team name [%s] is exist", name))
	}
	return nil
}

func (v *Validation) existByName(name string, ctx context.Context) bool {
	return v.client.Team.
		Query().
		Where(team.NameEQ(name)).
		ExistX(ctx)
}

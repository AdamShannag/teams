package user

import (
	"context"
	userrep "team-service/repository/user"
	"team-service/validation/violation"
)

type Validator struct {
	repository userrep.Repository
}

func NewValidator(repository userrep.Repository) *Validator {
	return &Validator{repository: repository}
}

func (v *Validator) Validate(userId string, ctx context.Context) (violations []violation.Violation) {
	if ok, vio := v.validateUser(userId, ctx); !ok {
		violations = append(violations, vio)
	}
	return
}

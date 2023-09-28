package user

import (
	"context"
	memberrep "team-service/repository/member"
	"team-service/validation/violation"
)

type Validator struct {
	repository memberrep.Repository
}

func NewValidator(repository memberrep.Repository) *Validator {
	return &Validator{repository: repository}
}

func (v *Validator) Validate(userId string, ctx context.Context) (violations []violation.Violation) {
	if ok, vio := v.validateUser(userId, ctx); !ok {
		violations = append(violations, vio)
	}
	return
}

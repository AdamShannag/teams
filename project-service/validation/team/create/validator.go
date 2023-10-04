package create

import (
	"context"
	teamrep "project-service/repository/project"
	"team-service/resource/team"
	"team-service/validation/violation"
)

type Validator struct {
	repository teamrep.Repository
}

func NewValidator(repository teamrep.Repository) *Validator {
	return &Validator{repository: repository}
}

func (v *Validator) Validate(request team.Request, ctx context.Context) (violations []violation.Violation) {
	return v.validateName(request.Name, ctx)
}

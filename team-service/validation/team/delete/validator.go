package delete

import (
	"context"
	teamrep "team-service/repository/team"
	"team-service/resource/team"
	"team-service/validation/violation"
)

type Validator struct {
	repository teamrep.Repository
}

func NewValidator(repository teamrep.Repository) *Validator {
	return &Validator{repository: repository}
}

func (v *Validator) Validate(request team.DeleteRequest, ctx context.Context) (violations []violation.Violation) {
	return v.validateTeamIds(request.TeamIds, ctx)
}
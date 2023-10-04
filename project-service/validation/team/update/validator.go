package update

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

func (v *Validator) Validate(request team.UpdateRequest, ctx context.Context) (violations []violation.Violation) {
	violations = append(violations, v.validateTeamId(request.TeamId, ctx)...)
	if request.Name != nil {
		violations = append(violations, v.validateName(*request.Name, ctx)...)
	}
	if request.Status != nil {
		violations = append(violations, v.validateStatus(*request.Status)...)
	}
	return
}

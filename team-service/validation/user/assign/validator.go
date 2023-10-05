package assign

import (
	"context"
	teamrep "team-service/repository/team"
	userrepo "team-service/repository/user"
	"team-service/resource/user"
	"team-service/validation/violation"
)

type Validator struct {
	userRepo userrepo.Repository
	teamRepo teamrep.Repository
}

func NewValidator(userRepo userrepo.Repository, teamRepo teamrep.Repository) *Validator {
	return &Validator{
		userRepo: userRepo,
		teamRepo: teamRepo,
	}
}

func (v *Validator) Validate(request user.AssignRequest, ctx context.Context) (violations []violation.Violation) {
	if request.IsAssign() {
		violations = append(violations, v.validateTeamId(request.TeamId, ctx)...)
	}
	violations = append(violations, v.validateUsers(request, ctx)...)
	return
}

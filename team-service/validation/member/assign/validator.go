package assign

import (
	"context"
	memberrepo "team-service/repository/member"
	teamrep "team-service/repository/team"
	"team-service/resource/member"
	"team-service/validation/violation"
)

type Validator struct {
	memberRepo memberrepo.Repository
	teamRepo   teamrep.Repository
}

func NewValidator(memberRepo memberrepo.Repository, teamRepo teamrep.Repository) *Validator {
	return &Validator{
		memberRepo: memberRepo,
		teamRepo:   teamRepo,
	}
}

func (v *Validator) Validate(request member.AssignRequest, ctx context.Context) (violations []violation.Violation) {
	if request.IsAssign() {
		violations = append(violations, v.validateTeamId(request.TeamId, ctx)...)
	}
	violations = append(violations, v.validateMembers(request, ctx)...)
	return
}

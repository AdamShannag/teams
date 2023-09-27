package assign

import (
	"context"
	"team-service/repository/ent"
	"team-service/resource/member"
	"team-service/validation/violation"
)

type Validator struct {
	client *ent.Client
}

func NewValidator(client *ent.Client) *Validator {
	return &Validator{
		client: client,
	}
}

func (v *Validator) Validate(request member.AssignRequest, ctx context.Context) (violations []violation.Violation) {
	if request.IsAssign() {
		violations = append(violations, v.validateTeamId(request.TeamId, ctx)...)
	}
	violations = append(violations, v.validateMembers(request, ctx)...)
	return
}

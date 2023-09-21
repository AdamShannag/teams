package update_team_validation

import (
	"context"
	"team-service/repository/ent"
	"team-service/resource/team"
	"team-service/validation/violation"
)

type Validation struct {
	client *ent.Client
}

func NewValidation(client *ent.Client) *Validation {
	return &Validation{
		client: client,
	}
}

func (v *Validation) Validate(request team.UpdateRequest, ctx context.Context) (violations []violation.Violation) {
	violations = append(violations, v.validateTeamId(request.TeamId, ctx)...)
	violations = append(violations, v.validateName(request.Name, ctx)...)
	violations = append(violations, v.validateStatus(request.Status)...)
	return
}

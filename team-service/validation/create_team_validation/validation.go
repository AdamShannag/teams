package create_team_validation

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

func (v *Validation) Validate(request team.Request, ctx context.Context) (violations []violation.Violation) {
	violations = append(violations, v.validateName(request.Name, ctx)...)
	violations = append(violations, v.validateCreatedBy(request.CreatedBy)...)
	return
}

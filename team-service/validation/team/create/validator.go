package create

import (
	"context"
	"team-service/repository/ent"
	"team-service/resource/team"
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

func (v *Validator) Validate(request team.Request, ctx context.Context) (violations []violation.Violation) {
	return v.validateName(request.Name, ctx)
}

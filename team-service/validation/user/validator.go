package user

import (
	"context"
	"team-service/repository/ent"
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

func (v *Validation) Validate(userId string, ctx context.Context) (violations []violation.Violation) {
	if ok, vio := v.validateUser(userId, ctx); !ok {
		violations = append(violations, vio)
	}
	return
}

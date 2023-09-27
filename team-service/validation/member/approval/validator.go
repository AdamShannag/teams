package approval

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

func (v *Validator) Validate(request member.ApprovalRequest, ctx context.Context) []violation.Violation {
	return v.validateAssignApproval(request.Members, ctx)
}

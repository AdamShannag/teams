package approval

import (
	"context"
	memberrep "team-service/repository/member"
	"team-service/resource/member"
	"team-service/validation/violation"
)

type Validator struct {
	repository memberrep.Repository
}

func NewValidator(repository memberrep.Repository) *Validator {
	return &Validator{repository: repository}
}

func (v *Validator) Validate(request member.ApprovalRequest, ctx context.Context) []violation.Violation {
	return v.validateAssignApproval(request.Members, ctx)
}

package approval

import (
	"context"
	userrep "team-service/repository/user"
	"team-service/resource/user"
	"team-service/validation/violation"
)

type Validator struct {
	repository userrep.Repository
}

func NewValidator(repository userrep.Repository) *Validator {
	return &Validator{repository: repository}
}

func (v *Validator) Validate(request user.ApprovalRequest, ctx context.Context) []violation.Violation {
	return v.validateAssignApproval(request.Users, ctx)
}

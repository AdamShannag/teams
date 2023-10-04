package member

import (
	"context"
	resource "project-service/resource/member"
	"team-service/repository/member"
	"team-service/service/log"
	"team-service/validation/validator"
	"team-service/validation/violation"
)

// Service instance for member's domain.
// Any operation done to any of object within this domain should use this service.
type Service interface {
	Assign(context.Context, *resource.AssignRequest, string) []violation.Violation
	AssignApproval(context.Context, *resource.ApprovalRequest, string) []violation.Violation
}

// beside embedding the struct, you can also declare the function directly on this struct.
// the advantage of embedding the struct is it allows spreading the implementation across multiple files.
type service struct {
	assign
	approval
}
type commonDependencies struct {
	repository    member.Repository
	userValidator validator.Validator[string]
	log           log.Assign
}

var _ Service = (*service)(nil)

// NewService Member service.
func NewService(
	repository member.Repository,
	userValidation validator.Validator[string],
	assignValidation validator.Validator[resource.AssignRequest],
	approvalValidation validator.Validator[resource.ApprovalRequest],
	log log.Service,

) Service {
	commonDependencies := commonDependencies{
		repository:    repository,
		userValidator: userValidation,
		log:           log.Assign,
	}
	return service{
		assign:   assign{commonDependencies: commonDependencies, validator: assignValidation},
		approval: approval{commonDependencies: commonDependencies, validator: approvalValidation},
	}
}

package user

import (
	"context"
	"team-service/repository/user"
	resource "team-service/resource/user"
	"team-service/service/log"
	"team-service/validation/validator"
	"team-service/validation/violation"
)

// Service instance for user's domain.
// Any operation done to any of object within this domain should use this service.
type Service interface {
	Assign(context.Context, *resource.AssignRequest, string) []violation.Violation
	AssignApproval(context.Context, *resource.ApprovalRequest, string) []violation.Violation
	Create(context.Context, string) error
}

// beside embedding the struct, you can also declare the function directly on this struct.
// the advantage of embedding the struct is it allows spreading the implementation across multiple files.
type service struct {
	assign
	approval
	create
}
type commonDependencies struct {
	repository    user.Repository
	userValidator validator.Validator[string]
	log           log.Assign
}

var _ Service = (*service)(nil)

// NewService User service.
func NewService(
	repository user.Repository,
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

func NewServiceCreate(
	repository user.Repository,
	log log.Service,
) Service {
	return service{
		create: create{repository: repository, log: log.Create},
	}
}

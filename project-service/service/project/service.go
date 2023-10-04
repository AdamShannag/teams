package project

import (
	"context"
	mapper "project-service/mapper/project"
	project2 "project-service/query/project"
	teamrep "project-service/repository/project"
	"project-service/resource/project"
	"project-service/service/log"
	"project-service/validation/validator"
	"project-service/validation/violation"
)

// Service instance for project's domain.
// Any operation done to any of object within this domain should use this service.
type Service interface {
	List(context.Context, project2.Query) (*project.ListResource, error)
	Get(context.Context, string) (*project.Resource, error)
	Create(context.Context, *project.Request, string) (*project.Resource, []violation.Violation)
	Update(context.Context, *project.UpdateRequest) (*project.Resource, []violation.Violation)
	Delete(context.Context, project.DeleteRequest) []violation.Violation
}

// beside embedding the struct, you can also declare the function directly on this struct.
// the advantage of embedding the struct is it allows spreading the implementation across multiple files.
type service struct {
	get
	create
	update
	delete
	list
}

type commonDependencies struct {
	repository teamrep.Repository
	mapper     mapper.Mapper
	log        log.Service
}

var _ Service = (*service)(nil)

// NewService Project service.
func NewService(
	teamRepository teamrep.Repository,
	mapper mapper.Mapper,
	log log.Service,
	userValidation validator.Validator[string],
	createValidation validator.Validator[project.Request],
	updateValidation validator.Validator[project.UpdateRequest],
	deleteValidator validator.Validator[project.DeleteRequest],
) Service {
	dependencies := commonDependencies{
		repository: teamRepository,
		mapper:     mapper,
	}
	return service{
		get:    get{commonDependencies: dependencies, log: log.Retrieve},
		list:   list{commonDependencies: dependencies, log: log.List},
		create: create{commonDependencies: dependencies, validator: createValidation, userValidator: userValidation, log: log.Create},
		update: update{commonDependencies: dependencies, validator: updateValidation, log: log.Update},
		delete: delete{commonDependencies: dependencies, validator: deleteValidator, log: log.Delete},
	}
}

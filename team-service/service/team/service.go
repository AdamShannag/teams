package team

import (
	"context"
	filter "team-service/filter/team"
	mapper "team-service/mapper/team"
	teamrep "team-service/repository/team"
	"team-service/resource/team"
	"team-service/service/log"
	page "team-service/service/pagination"
	"team-service/service/sorting"
	"team-service/validation/validator"
	"team-service/validation/violation"
)

// Service instance for team's domain.
// Any operation done to any of object within this domain should use this service.
type Service interface {
	List(context.Context, *page.Pagination, *filter.Filter, *sorting.Sort) (*team.ListResource, error)
	Get(context.Context, string) (*team.Resource, error)
	Create(context.Context, *team.Request, string) (*team.Resource, []violation.Violation)
	Update(context.Context, *team.UpdateRequest) (*team.Resource, []violation.Violation)
	Delete(context.Context, team.DeleteRequest) error
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

// NewService Team service.
func NewService(
	teamRepository teamrep.Repository,
	mapper mapper.Mapper,
	log log.Service,
	userValidation validator.Validator[string],
	createValidation validator.Validator[team.Request],
	updateValidation validator.Validator[team.UpdateRequest],
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
		delete: delete{commonDependencies: dependencies, log: log.Delete},
	}
}

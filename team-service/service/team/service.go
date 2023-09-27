package team

import (
	"context"
	"github.com/rs/zerolog"
	filter "team-service/filter/team"
	mapper "team-service/mapper/team"
	teamrep "team-service/repository/team"
	"team-service/resource/team"
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
	upd
	del
}

type commonDependencies struct {
	repository teamrep.Repository
	mapper     mapper.Mapper
	log        zerolog.Logger
}

var _ Service = (*service)(nil)

// NewService Team service.
func NewService(
	teamRepository teamrep.Repository,
	mapper mapper.Mapper,
	log zerolog.Logger,
	userValidation validator.Validator[string],
	createValidation validator.Validator[team.Request],
	updateValidation validator.Validator[team.UpdateRequest],
) Service {
	dependencies := commonDependencies{
		repository: teamRepository,
		mapper:     mapper,
		log:        log,
	}
	return service{
		get:    get{commonDependencies: dependencies},
		create: create{commonDependencies: dependencies, validator: createValidation, userValidator: userValidation},
		upd:    upd{commonDependencies: dependencies, validator: updateValidation},
		del:    del{commonDependencies: dependencies},
	}
}

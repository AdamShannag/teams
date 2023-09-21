package team

import (
	"context"
	"team-service/filter/teamfilter"
	mapper "team-service/mapper/team"
	"team-service/repository/ent"
	"team-service/resource/team"
	srv "team-service/service"
	"team-service/validation/create_team_validation"
	"team-service/validation/update_team_validation"
	"team-service/validation/violation"
)

// Service instance for team's domain.
// Any operation done to any of object within this domain should use this service.
type Service interface {
	List(context.Context, *srv.Pagination, *teamfilter.Filter) (*team.ListResource, error)
	Get(context.Context, string) (*team.Resource, error)
	Create(context.Context, *team.Request) (*team.Resource, []violation.Violation)
	Update(context.Context, *team.UpdateRequest) (*team.Resource, []violation.Violation)
	Delete(context.Context, string) error
}

// beside embedding the struct, you can also declare the function directly on this struct.
// the advantage of embedding the struct is it allows spreading the implementation across multiple files.
type service struct {
	get
	list
	create
	update
	delete
}

var _ Service = (*service)(nil)

// NewService Team service.
func NewService(
	client *ent.Client,
	mapper *mapper.Mapper,
	createValidation *create_team_validation.Validation,
	updateValidation *update_team_validation.Validation,
) Service {
	return service{
		get:    get{client: client, mapper: *mapper},
		list:   list{client: client, mapper: *mapper},
		create: create{client: client, mapper: *mapper, validation: createValidation},
		update: update{client: client, mapper: *mapper, validation: updateValidation},
		delete: delete{client: client},
	}
}

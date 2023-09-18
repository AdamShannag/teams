package team

import (
	"context"
	mapper "team-service/mapper/team"
	"team-service/repository/ent"
	"team-service/resource/team"
	srv "team-service/service"
	"team-service/validation"
)

// Service instance for team's domain.
// Any operation done to any of object within this domain should use this service.
type Service interface {
	List(context.Context, *srv.Pagination, *Filter) (*team.Resources, error)
	Get(context.Context, string) (*team.Resource, error)
	Create(context.Context, *team.Request) (*team.Resource, error)
	Update(context.Context, *team.UpdateRequest) (*team.Resource, error)
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
func NewService(client *ent.Client, mapper *mapper.Mapper, validation *validation.Validation) Service {
	return service{
		get:    get{client: client, mapper: *mapper},
		list:   list{client: client, mapper: *mapper},
		create: create{client: client, mapper: *mapper, validation: validation},
		update: update{client: client, mapper: *mapper, validation: validation},
		delete: delete{client: client},
	}
}

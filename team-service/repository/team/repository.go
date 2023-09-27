package team

import (
	"context"
	filter "team-service/filter/team"
	"team-service/repository/ent"
	"team-service/resource/team"
	page "team-service/service/pagination"
	"team-service/service/sorting"
)

// Repository represents a repository
type Repository interface {
	GetAll(context.Context, *page.Pagination, *filter.Filter, *sorting.Sort) ([]*ent.Team, error)
	Get(context.Context, string) (*ent.Team, error)
	GetAllAvailable(context.Context, *page.Pagination, *filter.Filter, *sorting.Sort) ([]*ent.Team, error)
	GetAvailable(context.Context, string) (*ent.Team, error)
	Save(context.Context, *ent.Team) (*ent.Team, error)
	Update(context.Context, *team.UpdateRequest) (*ent.Team, error)
	DeleteAll(context.Context, []string) error
	GetSize(context.Context) (int, error)
}

// Team represents a team
type repository struct {
	retrieve
	save
	update
	delete
}

var _ Repository = (*repository)(nil)

// NewRepository creates a new repository
func NewRepository(client ent.TeamClient) Repository {
	return repository{
		retrieve: retrieve{client: &client},
		save:     save{client: &client},
		update:   update{client: &client},
		delete:   delete{client: &client},
	}
}

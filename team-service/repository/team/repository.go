package team

import (
	"context"
	filter "team-service/filter/team"
	"team-service/repository/ent"
	team2 "team-service/repository/ent/team"
	"team-service/resource/team"
	page "team-service/service/pagination"
	"team-service/service/sorting"
)

// Repository represents a repository
type Repository interface {
	GetAll(context.Context, *page.Pagination, *filter.Filter, *sorting.Sort) ([]*ent.Team, error)
	GetAllWithStatusNot(context.Context, *page.Pagination, *filter.Filter, *sorting.Sort, team2.Status) ([]*ent.Team, error)
	Get(context.Context, string) (*ent.Team, error)
	GetAvailable(context.Context, string) (*ent.Team, error)
	Save(context.Context, *ent.Team) (*ent.Team, error)
	Update(context.Context, *team.UpdateRequest) (*ent.Team, error)
	DeleteAll(context.Context, []string) error
	GetSize(context.Context) (int, error)
	ExistByName(context.Context, string) (bool, error)
	ExistByNameAndStatusNot(context.Context, string, team2.Status) (bool, error)
	ExistByIdAndStatusNot(context.Context, string, team2.Status) (bool, error)
}

// Team represents a team
type repository struct {
	client *ent.TeamClient
}

var _ Repository = (*repository)(nil)

// NewRepository creates a new repository
func NewRepository(client ent.TeamClient) Repository {
	return repository{&client}
}

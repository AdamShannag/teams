package project

import (
	"context"
	filter "project-service/filter/project"
	"project-service/repository/ent"
	project2 "project-service/repository/ent/project"
	"project-service/resource/project"
	page "project-service/service/pagination"
	"project-service/service/sorting"
)

// Repository represents a repository
type Repository interface {
	GetAll(context.Context, *page.Pagination, *filter.Filter, *sorting.Sort) ([]*ent.Project, error)
	GetAllWithStatusNot(context.Context, *page.Pagination, *filter.Filter, *sorting.Sort, project2.Status) ([]*ent.Project, error)
	Get(context.Context, string) (*ent.Project, error)
	Save(context.Context, *ent.Project) (*ent.Project, error)
	Update(context.Context, *project.UpdateRequest) (*ent.Project, error)
	DeleteAll(context.Context, []string) error
	GetSize(context.Context) (int, error)
	ExistByName(context.Context, string) (bool, error)
	ExistByNameAndStatusNot(context.Context, string, project2.Status) (bool, error)
	ExistByIdAndStatusNot(context.Context, string, project2.Status) (bool, error)
}

// Project represents a project
type repository struct {
	client *ent.ProjectClient
}

var _ Repository = (*repository)(nil)

// NewRepository creates a new repository
func NewRepository(client ent.ProjectClient) Repository {
	return repository{&client}
}

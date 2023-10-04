package project

import (
	"context"
	filter "project-service/filter/project"
	"project-service/repository/ent"
	"project-service/repository/ent/project"
	page "project-service/service/pagination"
	"project-service/service/sorting"
)

func (r repository) Get(ctx context.Context, projectId string) (*ent.Project, error) {
	return r.client.Get(ctx, projectId)
}

func (r repository) GetAll(ctx context.Context, pagination *page.Pagination, filter *filter.Filter, sort *sorting.Sort) ([]*ent.Project, error) {
	return r.client.
		Query().
		Where(filter.Predicate...).
		Limit(pagination.Size).
		Offset(pagination.Page).
		Order(sort.Order).
		All(ctx)
}

func (r repository) GetAvailable(ctx context.Context, projectId string) (*ent.Project, error) {
	return r.client.Query().Where(project.ID(projectId), project.StatusEQ(project.StatusAVAILABLE)).Only(ctx)
}

func (r repository) GetAllWithStatusNot(ctx context.Context, pagination *page.Pagination, filter *filter.Filter, sort *sorting.Sort, status project.Status) ([]*ent.Project, error) {
	return r.client.
		Query().
		Where(project.StatusNEQ(status)).
		Where(filter.Predicate...).
		Limit(pagination.Size).
		Offset(pagination.Page).
		Order(sort.Order).
		All(ctx)
}

func (r repository) GetSize(ctx context.Context) (int, error) {
	return r.client.Query().Count(ctx)
}

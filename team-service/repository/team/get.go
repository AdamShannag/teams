package team

import (
	"context"
	filter "team-service/filter/team"
	"team-service/repository/ent"
	"team-service/repository/ent/team"
	page "team-service/service/pagination"
	"team-service/service/sorting"
)

func (r repository) Get(ctx context.Context, teamedId string) (*ent.Team, error) {
	return r.client.Get(ctx, teamedId)
}

func (r repository) GetAll(ctx context.Context, pagination *page.Pagination, filter *filter.Filter, sort *sorting.Sort) ([]*ent.Team, error) {
	return r.client.
		Query().
		Where(filter.Predicate...).
		Limit(pagination.Size).
		Offset(pagination.Page).
		Order(sort.Order).
		All(ctx)
}

func (r repository) GetAvailable(ctx context.Context, teamedId string) (*ent.Team, error) {
	return r.client.Query().Where(team.ID(teamedId), team.StatusEQ(team.StatusAVAILABLE)).Only(ctx)
}

func (r repository) GetAllWithStatusNot(ctx context.Context, pagination *page.Pagination, filter *filter.Filter, sort *sorting.Sort, status team.Status) ([]*ent.Team, error) {
	return r.client.
		Query().
		Where(team.StatusNEQ(status)).
		Where(filter.Predicate...).
		Limit(pagination.Size).
		Offset(pagination.Page).
		Order(sort.Order).
		All(ctx)
}

func (r repository) GetSize(ctx context.Context) (int, error) {
	return r.client.Query().Count(ctx)
}

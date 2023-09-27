package team

import (
	"context"
	filter "team-service/filter/team"
	"team-service/repository/ent"
	"team-service/repository/ent/team"
	page "team-service/service/pagination"
	"team-service/service/sorting"
)

type retrieve struct {
	client *ent.TeamClient
}

func (r retrieve) Get(ctx context.Context, teamedId string) (*ent.Team, error) {
	return r.client.Get(ctx, teamedId)
}

func (r retrieve) GetAll(ctx context.Context, pagination *page.Pagination, filter *filter.Filter, sort *sorting.Sort) ([]*ent.Team, error) {
	return r.client.
		Query().
		Where(filter.Predicate...).
		Limit(pagination.Size).
		Offset(pagination.Page).
		Order(sort.Order).
		All(ctx)
}

func (r retrieve) GetAvailable(ctx context.Context, teamedId string) (*ent.Team, error) {
	return r.client.Query().Where(team.ID(teamedId), team.StatusEQ(team.StatusAVAILABLE)).Only(ctx)
}

func (r retrieve) GetAllAvailable(ctx context.Context, pagination *page.Pagination, filter *filter.Filter, sort *sorting.Sort) ([]*ent.Team, error) {
	return r.client.
		Query().
		Where(filter.Predicate...).
		Where(team.StatusEQ(team.StatusAVAILABLE)).
		Limit(pagination.Size).
		Offset(pagination.Page).
		Order(sort.Order).
		All(ctx)
}

func (r retrieve) GetSize(ctx context.Context) (int, error) {
	return r.client.Query().Count(ctx)
}

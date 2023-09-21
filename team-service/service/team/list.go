package team

import (
	"context"
	"team-service/filter/teamfilter"
	mapper "team-service/mapper/team"
	"team-service/repository/ent"
	"team-service/resource/team"
	srv "team-service/service"
)

type list struct {
	client *ent.Client
	mapper mapper.Mapper
}

func (s list) List(ctx context.Context, page *srv.Pagination, filter *teamfilter.Filter) (*team.ListResource, error) {
	var resources = []team.Resource{}

	query := s.client.Team.Query().Where(filter.Predicate...)

	teams, err := query.Limit(page.Size).Offset(page.Page).All(ctx)
	if err != nil {
		return nil, err
	}

	totalItems := s.getTotalItems(ctx)
	sizeItems := len(teams)

	for i, t := range teams {
		resource := s.mapper.ToResource(t)
		resource.SeqNo = i + 1 + page.Page
		resources = append(resources, resource)
	}

	return &team.ListResource{
		Content:    resources,
		Pagination: *page.GetResource(sizeItems, totalItems),
	}, err
}

func (s list) getTotalItems(ctx context.Context) int {
	return s.client.Team.Query().CountX(ctx)
}

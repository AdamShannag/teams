package team

import (
	"context"
	mapper "team-service/mapper/team"
	"team-service/repository/ent"
	"team-service/resource/team"
	srv "team-service/service"
)

// Filter for search.
type Filter struct {
	team.Resource
}

type list struct {
	client *ent.Client
	mapper mapper.Mapper
}

func (s list) List(ctx context.Context, page *srv.Pagination, filter *Filter) (*team.Resources, error) {
	var resources = []team.Resource{}

	teams, err := s.client.Team.Query().Limit(page.Size).Offset(page.Page).Order(ent.Asc("created_by")).All(ctx)
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

	return &team.Resources{
		Content:    resources,
		Pagination: *page.GetResource(sizeItems, totalItems),
	}, err
}

func (s list) getTotalItems(ctx context.Context) int {
	return s.client.Team.Query().CountX(ctx)
}

package team

import (
	"context"
	"team-service/repository/ent"
	team3 "team-service/repository/ent/team"
	"team-service/resource/team"
	"team-service/service/log"
	team2 "team-service/service/query/team"
)

type list struct {
	commonDependencies
	log log.List
}

func (s list) List(ctx context.Context, query team2.Query) (*team.ListResource, error) {
	var resources = []team.Resource{}

	page := &query.Pagination

	teams, err := s.getAll(ctx, query)

	if err != nil {
		s.log.Failed("team", err)
		return nil, err
	}

	totalItems, _ := s.repository.GetSize(ctx)
	sizeItems := len(teams)

	for i, t := range teams {
		resource := s.mapper.ToResource(t)
		resource.SeqNo = i + 1 + page.Page
		resources = append(resources, resource)
	}

	s.log.Success("team", sizeItems)
	return &team.ListResource{
		Content:    resources,
		Pagination: *page.GetResource(sizeItems, totalItems),
	}, nil
}

func (s list) getAll(ctx context.Context, query team2.Query) ([]*ent.Team, error) {
	page := &query.Pagination
	filter := &query.Filter
	sort := &query.Sort
	if query.IsAll {
		return s.repository.GetAll(ctx, page, filter, sort)
	} else {
		return s.repository.GetAllWithStatusNot(ctx, page, filter, sort, team3.StatusDELETED)
	}
}

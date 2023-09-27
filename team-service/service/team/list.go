package team

import (
	"context"
	filter "team-service/filter/team"
	"team-service/resource/team"
	"team-service/service/log"
	page "team-service/service/pagination"
	"team-service/service/sorting"
)

type list struct {
	commonDependencies
	log log.List
}

func (s list) List(ctx context.Context, page *page.Pagination, filter *filter.Filter, sort *sorting.Sort) (*team.ListResource, error) {
	var resources = []team.Resource{}

	teams, err := s.repository.GetAll(ctx, page, filter, sort)
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

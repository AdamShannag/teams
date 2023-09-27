package team

import (
	"context"
	"team-service/constant/message"
	filter "team-service/filter/team"
	"team-service/resource/team"
	page "team-service/service/pagination"
	"team-service/service/sorting"
)

type get struct {
	commonDependencies
}

func (s get) Get(ctx context.Context, teamedId string) (*team.Resource, error) {
	result, err := s.repository.Get(ctx, teamedId)

	if err != nil {
		s.log.Error().Err(err).Msgf(message.RETRIEVED_FAILED, "team", teamedId)
		return nil, err
	}

	resource := s.mapper.ToResource(result)

	s.log.Info().Msgf(message.RETRIEVED_SUCCESSFULLY, "team", resource.ID)
	return &resource, nil
}

func (s get) List(ctx context.Context, page *page.Pagination, filter *filter.Filter, sort *sorting.Sort) (*team.ListResource, error) {
	var resources = []team.Resource{}

	teams, err := s.repository.GetAll(ctx, page, filter, sort)
	if err != nil {
		s.log.Error().Err(err).Msgf(message.RETRIEVED_LIST_FAILED, "teams")
		return nil, err
	}

	totalItems, _ := s.repository.GetSize(ctx)
	sizeItems := len(teams)

	for i, t := range teams {
		resource := s.mapper.ToResource(t)
		resource.SeqNo = i + 1 + page.Page
		resources = append(resources, resource)
	}

	s.log.Info().Msgf(message.RETRIEVED_LIST_SUCCESSFULLY, "teams", sizeItems)
	return &team.ListResource{
		Content:    resources,
		Pagination: *page.GetResource(sizeItems, totalItems),
	}, nil
}

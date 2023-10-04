package project

import (
	"context"
	project2 "project-service/query/project"
	"project-service/repository/ent"
	project3 "project-service/repository/ent/project"
	"project-service/resource/project"
	"project-service/service/log"
)

type list struct {
	commonDependencies
	log log.List
}

func (s list) List(ctx context.Context, query project2.Query) (*project.ListResource, error) {
	var resources = []project.Resource{}

	page := &query.Pagination

	teams, err := s.getAll(ctx, query)

	if err != nil {
		s.log.Failed("project", err)
		return nil, err
	}

	totalItems, _ := s.repository.GetSize(ctx)
	sizeItems := len(teams)

	for i, t := range teams {
		resource := s.mapper.ToResource(t)
		resource.SeqNo = i + 1 + page.Page
		resources = append(resources, resource)
	}

	s.log.Success("project", sizeItems)
	return &project.ListResource{
		Content:    resources,
		Pagination: *page.GetResource(sizeItems, totalItems),
	}, nil
}

func (s list) getAll(ctx context.Context, query project2.Query) ([]*ent.Project, error) {
	page := &query.Pagination
	filter := &query.Filter
	sort := &query.Sort
	if query.IsAll {
		return s.repository.GetAll(ctx, page, filter, sort)
	} else {
		return s.repository.GetAllWithStatusNot(ctx, page, filter, sort, project3.StatusDELETED)
	}
}

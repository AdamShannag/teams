package project

import (
	"context"
	"project-service/resource/project"
	"project-service/service/log"
)

type get struct {
	commonDependencies
	log log.Retrieve
}

func (s get) Get(ctx context.Context, teamId string) (*project.Resource, error) {
	result, err := s.repository.Get(ctx, teamId)

	if err != nil {
		s.log.Failed("project", err, teamId)
		return nil, err
	}

	resource := s.mapper.ToResource(result)

	s.log.Success("project", resource.ID)
	return &resource, nil
}

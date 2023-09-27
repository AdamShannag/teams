package team

import (
	"context"
	"team-service/resource/team"
	"team-service/service/log"
)

type get struct {
	commonDependencies
	log log.Retrieve
}

func (s get) Get(ctx context.Context, teamId string) (*team.Resource, error) {
	result, err := s.repository.Get(ctx, teamId)

	if err != nil {
		s.log.Failed("team", err, teamId)
		return nil, err
	}

	resource := s.mapper.ToResource(result)

	s.log.Success("team", resource.ID)
	return &resource, nil
}

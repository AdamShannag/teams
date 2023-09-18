package team

import (
	"context"
	mapper "team-service/mapper/team"
	"team-service/repository/ent"
	"team-service/resource/team"
)

type get struct {
	client *ent.Client
	mapper mapper.Mapper
}

func (s get) Get(ctx context.Context, teamedId string) (*team.Resource, error) {
	t, err := s.client.Team.Get(ctx, teamedId)

	if err != nil {
		return nil, err
	}

	resource := s.mapper.ToResource(t)

	return &resource, nil
}

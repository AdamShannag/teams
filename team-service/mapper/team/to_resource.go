package team

import (
	"team-service/repository/ent"
	"team-service/resource/team"
)

type toResource struct{}

func (m toResource) ToResource(repo *ent.Team) team.Resource {
	return team.Resource{
		Team: *repo,
	}
}

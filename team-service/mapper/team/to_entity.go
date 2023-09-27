package team

import (
	"team-service/repository/ent"
	"team-service/resource/team"
)

type requestToEntity struct{}

func (m requestToEntity) ToEntity(request team.Request) ent.Team {
	return ent.Team{
		Name:        request.Name,
		Description: request.Description,
	}
}

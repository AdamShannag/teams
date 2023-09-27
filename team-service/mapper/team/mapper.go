package team

import (
	"team-service/repository/ent"
	"team-service/resource/team"
)

type Mapper interface {
	ToResource(*ent.Team) team.Resource
	ToEntity(team.Request) ent.Team
}

type mapper struct {
	toResource
	requestToEntity
}

func NewMapper() Mapper {
	return mapper{
		toResource:      toResource{},
		requestToEntity: requestToEntity{},
	}
}

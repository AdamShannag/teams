package team

import (
	"team-service/repository/ent"
	"team-service/resource/team"
)

type Mapper interface {
	ToResource(*ent.Team) team.Resource
}

type mapper struct {
	toResource
}

func NewMapper() Mapper {
	return mapper{
		toResource: toResource{},
	}
}

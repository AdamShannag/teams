package project

import (
	"project-service/repository/ent"
	"project-service/resource/project"
)

type Mapper interface {
	ToResource(*ent.Project) project.Resource
	ToEntity(project.Request) ent.Project
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

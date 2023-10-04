package project

import (
	"project-service/repository/ent"
	"project-service/resource/project"
)

type requestToEntity struct{}

func (m requestToEntity) ToEntity(request project.Request) ent.Project {
	return ent.Project{}
}

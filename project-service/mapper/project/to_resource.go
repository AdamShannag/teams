package project

import (
	"project-service/repository/ent"
	"project-service/resource/project"
)

type toResource struct{}

func (m toResource) ToResource(repo *ent.Project) project.Resource {
	return project.Resource{
		Project: *repo,
	}
}

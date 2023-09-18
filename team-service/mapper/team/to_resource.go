package team

import (
	"team-service/repository/ent"
	"team-service/resource/team"
)

type toResource struct{}

func (m toResource) ToResource(repo *ent.Team) team.Resource {
	return team.Resource{
		ID:          repo.ID,
		Name:        repo.Name,
		Description: repo.Description,
		Status:      repo.Status,
		CreatedBy:   repo.CreatedBy,
		CreatedAt:   repo.CreatedAt,
		UpdatedAt:   repo.UpdatedAt,
	}
}

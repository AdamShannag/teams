package project

import "project-service/repository/ent/project"

type Request struct {
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
}

type UpdateRequest struct {
	ProjectId   *string         `json:"projectId,omitempty"`
	Name        *string         `json:"name,omitempty"`
	Description *string         `json:"description,omitempty"`
	Status      *project.Status `json:"status,omitempty"`
}

type DeleteRequest struct {
	ProjectIds []string `json:"projectIds"`
}

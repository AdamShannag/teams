package team

import (
	"team-service/repository/ent/team"
)

type Request struct {
	Name        string `validate:"required" json:"name"`
	Description string `json:"description,omitempty"`
	CreatedBy   string `json:"created_by"`
}

type UpdateRequest struct {
	ID          string      `json:"id"`
	Name        string      `json:"name"`
	Description string      `json:"description,omitempty"`
	Status      team.Status `json:"status"`
}

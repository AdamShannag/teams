package team

import (
	"team-service/repository/ent/team"
)

type Request struct {
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
}

type UpdateRequest struct {
	TeamId      *string      `json:"teamId,omitempty"`
	Name        *string      `json:"name,omitempty"`
	Description *string      `json:"description,omitempty"`
	Status      *team.Status `json:"status,omitempty"`
}

type DeleteRequest struct {
	TeamIds []string `json:"teamIds"`
}

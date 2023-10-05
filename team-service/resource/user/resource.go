package user

import (
	"team-service/repository/ent/user"
)

type AssignResource struct {
	Status user.Status `json:"status"`
	TeamId string      `json:"teamId"`
	UserID string      `json:"userID"`
	Users  []string    `json:"users"`
}

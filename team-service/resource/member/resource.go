package member

import (
	"team-service/repository/ent/member"
)

type AssignResource struct {
	Status  member.Status `json:"status"`
	TeamId  string        `json:"teamId"`
	UserID  string        `json:"userID"`
	Members []string      `json:"members"`
}

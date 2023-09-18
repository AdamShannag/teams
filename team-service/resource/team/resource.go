package team

import (
	"github.com/mohammadyaseen2/pagination/model"
	"team-service/repository/ent/team"
	"time"
)

type Resource struct {
	ID          string      `json:"id"`
	SeqNo       int         `json:"seqNo,omitempty"`
	Name        string      `json:"name"`
	Description string      `json:"description,omitempty"`
	Status      team.Status `json:"status"`
	CreatedBy   string      `json:"createdBy"`
	CreatedAt   time.Time   `json:"createdAt"`
	UpdatedAt   time.Time   `json:"updatedAt"`
}

type Resources struct {
	Content    []Resource     `json:"content"`
	Pagination model.Resource `json:"pagination,omitempty"`
}

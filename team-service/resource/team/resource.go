package team

import (
	"github.com/mohammadyaseen2/pagination/model"
	"team-service/repository/ent"
)

type Resource struct {
	SeqNo int `json:"seqNo,omitempty"`
	ent.Team
}

type ListResource struct {
	Content    []Resource     `json:"content"`
	Pagination model.Resource `json:"pagination,omitempty"`
}

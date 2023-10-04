package project

import (
	"github.com/mohammadyaseen2/pagination/model"
	"project-service/repository/ent"
)

type Resource struct {
	SeqNo int `json:"seqNo,omitempty"`
	ent.Project
}

type ListResource struct {
	Content    []Resource     `json:"content"`
	Pagination model.Resource `json:"pagination,omitempty"`
}

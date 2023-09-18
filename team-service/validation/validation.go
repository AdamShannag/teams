package validation

import (
	"context"
	"github.com/go-playground/validator"
	"team-service/repository/ent"
)

type Validation struct {
	client    *ent.Client
	validator *validator.Validate
}

func NewValidation(client *ent.Client) *Validation {
	return &Validation{
		client: client,
	}
}

type TeamValidation interface {
	ExistByName(string, context.Context) error
}

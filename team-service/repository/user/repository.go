package user

import (
	"context"
	"team-service/repository/ent"
	user2 "team-service/repository/ent/user"
	"team-service/resource/user"
)

// Repository represents a repository
type Repository interface {
	Save(context.Context, string) (*ent.User, error)
	Assign(context.Context, user.AssignResource) error
	UnAssign(context.Context, user.AssignResource) error
	Approve(context.Context, user.AssignResource) error
	Reject(context.Context, user.AssignResource) error
	ExistById(context.Context, string) (bool, error)
	ExistByIdAndStatus(context.Context, string, user2.Status) (bool, error)
}

// User represents a user
type repository struct {
	client *ent.UserClient
}

var _ Repository = (*repository)(nil)

// NewRepository creates a new repository
func NewRepository(client ent.UserClient) Repository {
	return repository{&client}
}

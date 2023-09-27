package member

import (
	"context"
	"team-service/repository/ent"
	"team-service/resource/member"
)

// Repository represents a repository
type Repository interface {
	Assign(context.Context, member.AssignResource) error
	UnAssign(context.Context, member.AssignResource) error
	Approve(context.Context, member.AssignResource) error
	Reject(context.Context, member.AssignResource) error
}

// Team represents a team
type repository struct {
	client *ent.MemberClient
}

var _ Repository = (*repository)(nil)

// NewRepository creates a new repository
func NewRepository(client ent.MemberClient) Repository {
	return repository{client: &client}
}

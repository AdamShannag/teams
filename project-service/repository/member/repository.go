package member

import (
	"context"
	member2 "project-service/repository/ent/member"
	"team-service/repository/ent"
	"team-service/resource/member"
)

// Repository represents a repository
type Repository interface {
	Assign(context.Context, member.AssignResource) error
	UnAssign(context.Context, member.AssignResource) error
	Approve(context.Context, member.AssignResource) error
	Reject(context.Context, member.AssignResource) error
	ExistById(context.Context, string) (bool, error)
	ExistByIdAndStatus(context.Context, string, member2.Status) (bool, error)
}

// Member represents a member
type repository struct {
	client *ent.MemberClient
}

var _ Repository = (*repository)(nil)

// NewRepository creates a new repository
func NewRepository(client ent.MemberClient) Repository {
	return repository{&client}
}

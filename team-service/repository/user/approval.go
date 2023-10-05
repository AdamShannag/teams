package user

import (
	"context"
	"team-service/repository/ent/user"
	resource "team-service/resource/user"
)

func (r repository) Approve(ctx context.Context, resource resource.AssignResource) error {
	return r.client.
		Update().
		Where(user.IDIn(resource.Users...)).
		SetStatus(resource.Status).
		SetApprovedBy(resource.UserID).
		Exec(ctx)
}

func (r repository) Reject(ctx context.Context, resource resource.AssignResource) error {
	return r.UnAssign(ctx, resource)
}

package member

import (
	"context"
	"team-service/repository/ent/member"
	resource "team-service/resource/member"
)

func (r repository) Approve(ctx context.Context, resource resource.AssignResource) error {
	return r.client.
		Update().
		Where(member.IDIn(resource.Members...)).
		SetStatus(resource.Status).
		SetApprovedBy(resource.UserID).
		Exec(ctx)
}

func (r repository) Reject(ctx context.Context, resource resource.AssignResource) error {
	return r.UnAssign(ctx, resource)
}

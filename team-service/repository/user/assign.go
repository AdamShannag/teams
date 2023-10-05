package user

import (
	"context"
	"team-service/repository/ent/user"
	resource "team-service/resource/user"
)

func (r repository) Assign(ctx context.Context, resource resource.AssignResource) error {
	return r.client.
		Update().
		Where(user.IDIn(resource.Users...)).
		SetStatus(resource.Status).
		SetTeamID(resource.TeamId).
		SetAssignedBy(resource.UserID).
		Exec(ctx)
}

func (r repository) UnAssign(ctx context.Context, resource resource.AssignResource) error {
	return r.client.
		Update().
		Where(user.IDIn(resource.Users...)).
		SetStatus(resource.Status).
		ClearTeamID().
		ClearAssignedBy().
		ClearApprovedBy().
		Exec(ctx)
}

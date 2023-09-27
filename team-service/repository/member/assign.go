package member

import (
	"context"
	"team-service/repository/ent/member"
	resource "team-service/resource/member"
)

func (r repository) Assign(ctx context.Context, resource resource.AssignResource) error {
	return r.client.
		Update().
		Where(member.IDIn(resource.Members...)).
		SetStatus(resource.Status).
		SetTeamID(resource.TeamId).
		SetAssignedBy(resource.UserID).
		Exec(ctx)
}

func (r repository) UnAssign(ctx context.Context, resource resource.AssignResource) error {
	return r.client.
		Update().
		Where(member.IDIn(resource.Members...)).
		SetStatus(resource.Status).
		ClearTeamID().
		ClearAssignedBy().
		ClearApprovedBy().
		Exec(ctx)
}

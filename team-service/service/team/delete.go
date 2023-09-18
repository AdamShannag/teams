package team

import (
	"context"
	"team-service/repository/ent"
	t "team-service/repository/ent/team"
)

type delete struct {
	client *ent.Client
}

func (h delete) Delete(ctx context.Context, teamedId string) error {
	err := h.client.Team.
		UpdateOneID(teamedId).
		SetStatus(t.StatusDELETED).
		Exec(ctx)

	if err != nil {
		return err
	}
	return nil
}

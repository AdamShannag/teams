package team

import (
	"context"
	"team-service/repository/ent"
	t "team-service/repository/ent/team"
)

type delete struct {
	client *ent.TeamClient
}

func (r delete) DeleteAll(ctx context.Context, teamedIds []string) error {
	return r.client.
		Update().
		Where(t.IDIn(teamedIds...)).
		SetStatus(t.StatusDELETED).
		Exec(ctx)
}

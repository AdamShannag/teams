package team

import (
	"context"
	t "team-service/repository/ent/team"
)

func (r repository) DeleteAll(ctx context.Context, teamedIds []string) error {
	return r.client.
		Update().
		Where(t.IDIn(teamedIds...)).
		SetStatus(t.StatusDELETED).
		Exec(ctx)
}

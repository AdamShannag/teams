package project

import (
	"context"
	t "project-service/repository/ent/project"
)

func (r repository) DeleteAll(ctx context.Context, projectIds []string) error {
	return r.client.
		Update().
		Where(t.IDIn(projectIds...)).
		SetStatus(t.StatusDELETED).
		Exec(ctx)
}

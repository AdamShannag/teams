package team

import (
	"context"
	"team-service/repository/ent/team"
)

func (r repository) ExistByName(ctx context.Context, name string) (bool, error) {
	return r.client.
		Query().
		Where(team.NameEQ(name)).
		Exist(ctx)
}

func (r repository) ExistByNameAndStatusNot(ctx context.Context, name string, status team.Status) (bool, error) {
	return r.client.
		Query().
		Where(team.NameEQ(name)).
		Where(team.StatusNEQ(status)).
		Exist(ctx)
}

func (r repository) ExistByIdAndStatusNot(ctx context.Context, id string, status team.Status) (bool, error) {
	return r.client.
		Query().
		Where(team.ID(id)).
		Where(team.StatusNEQ(status)).
		Exist(ctx)
}

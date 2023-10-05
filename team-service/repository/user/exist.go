package user

import (
	"context"
	"team-service/repository/ent/user"
)

func (r repository) ExistByIdAndStatus(ctx context.Context, id string, status user.Status) (bool, error) {
	return r.client.
		Query().
		Where(user.ID(id)).
		Where(user.StatusEQ(status)).
		Exist(ctx)
}

func (r repository) ExistById(ctx context.Context, id string) (bool, error) {
	return r.client.
		Query().
		Where(user.ID(id)).
		Exist(ctx)
}

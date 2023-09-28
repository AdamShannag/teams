package member

import (
	"context"
	"team-service/repository/ent/member"
)

func (r repository) ExistByIdAndStatus(ctx context.Context, id string, status member.Status) (bool, error) {
	return r.client.
		Query().
		Where(member.ID(id)).
		Where(member.StatusEQ(status)).
		Exist(ctx)
}

func (r repository) ExistById(ctx context.Context, id string) (bool, error) {
	return r.client.
		Query().
		Where(member.ID(id)).
		Exist(ctx)
}

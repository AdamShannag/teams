package user

import (
	"context"
	"team-service/repository/ent"
	"team-service/repository/ent/user"
)

func (r repository) Save(ctx context.Context, userId string) (*ent.User, error) {
	return r.client.
		Create().
		SetID(userId).
		SetStatus(user.StatusFREE).
		Save(ctx)
}

package team

import (
	"context"
	"team-service/repository/ent"
)

type save struct {
	client *ent.TeamClient
}

func (r save) Save(ctx context.Context, entity *ent.Team) (*ent.Team, error) {
	return r.client.
		Create().
		SetID(entity.ID).
		SetStatus(entity.Status).
		SetName(entity.Name).
		SetDescription(entity.Description).
		SetCreatedBy(entity.CreatedBy).
		Save(ctx)
}

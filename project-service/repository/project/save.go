package project

import (
	"context"
	"project-service/repository/ent"
)

func (r repository) Save(ctx context.Context, entity *ent.Project) (*ent.Project, error) {
	return r.client.
		Create().
		SetID(entity.ID).
		SetStatus(entity.Status).
		SetName(entity.Name).
		SetDescription(entity.Description).
		SetCreatedBy(entity.CreatedBy).
		Save(ctx)
}

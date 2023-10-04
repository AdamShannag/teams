package project

import (
	"context"
	"project-service/repository/ent/project"
)

func (r repository) ExistByName(ctx context.Context, name string) (bool, error) {
	return r.client.
		Query().
		Where(project.NameEQ(name)).
		Exist(ctx)
}

func (r repository) ExistByNameAndStatusNot(ctx context.Context, name string, status project.Status) (bool, error) {
	return r.client.
		Query().
		Where(project.NameEQ(name)).
		Where(project.StatusNEQ(status)).
		Exist(ctx)
}

func (r repository) ExistByIdAndStatusNot(ctx context.Context, id string, status project.Status) (bool, error) {
	return r.client.
		Query().
		Where(project.ID(id)).
		Where(project.StatusNEQ(status)).
		Exist(ctx)
}

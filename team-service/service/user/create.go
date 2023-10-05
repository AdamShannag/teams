package user

import (
	"context"
	"team-service/repository/user"
	"team-service/service/log"
)

type create struct {
	repository user.Repository
	log        log.Create
}

func (s create) Create(ctx context.Context, userId string) error {
	_, err := s.repository.Save(ctx, userId)

	if err != nil {
		s.log.Failed("user", err)
		return err
	}

	s.log.Success("user", userId)
	return nil
}

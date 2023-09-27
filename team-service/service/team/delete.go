package team

import (
	"context"
	"team-service/constant/message"
	"team-service/resource/team"
)

type del struct {
	commonDependencies
}

func (s del) Delete(ctx context.Context, request team.DeleteRequest) (err error) {
	err = s.repository.DeleteAll(ctx, request.TeamIds)

	if err != nil {
		s.log.Error().Err(err).Msgf(message.DELETED_FAILED, "teams", request.TeamIds)
		return
	}
	s.log.Info().Msgf(message.DELETED_SUCCESSFULLY, "teams", request.TeamIds)
	return
}

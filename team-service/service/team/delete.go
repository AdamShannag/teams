package team

import (
	"context"
	"team-service/resource/team"
	"team-service/service/log"
)

type delete struct {
	commonDependencies
	log log.Delete
}

func (s delete) Delete(ctx context.Context, request team.DeleteRequest) (err error) {
	err = s.repository.DeleteAll(ctx, request.TeamIds)

	if err != nil {
		s.log.Failed("teams", err, request.TeamIds...)
		return
	}
	s.log.Success("teams", request.TeamIds)
	return
}

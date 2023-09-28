package teams

import (
	"github.com/go-chi/chi/v5"
	"net/http"
	"team-service/constant"
	team2 "team-service/service/query/team"
)

func (t *Teams) GetTeams(w http.ResponseWriter, r *http.Request) {
	var (
		ctx   = r.Context()
		query = team2.NewQuery(r)
	)

	result, err := t.service.List(ctx, *query)
	if err != nil {
		t.Error(w, err)
		return
	}

	t.Succeed(w, result)
}

func (t *Teams) GetTeam(w http.ResponseWriter, r *http.Request) {
	var (
		ctx    = r.Context()
		teamId = chi.URLParam(r, constant.TEAM_ID)
	)
	result, err := t.service.Get(ctx, teamId)
	if err != nil {
		t.Error(w, err)
		return
	}

	t.Succeed(w, result)
}

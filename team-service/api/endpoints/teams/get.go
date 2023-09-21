package teams

import (
	"github.com/go-chi/chi/v5"
	"net/http"
	"team-service/filter/teamfilter"
	"team-service/service"
)

const (
	TEAM_ID = "teamId"
)

func (t *Teams) GetTeams(w http.ResponseWriter, r *http.Request) {
	var (
		ctx        = r.Context()
		query      = r.URL.Query()
		filter     = teamfilter.NewFilter(query)
		pagination = service.NewPagination(
			query.Get("page"),
			query.Get("size"),
		)
	)

	result, err := t.teams.List(ctx, pagination, filter)
	if err != nil {
		t.Error(w, err)
		return
	}

	t.Succeed(w, result)
}

func (t *Teams) GetTeam(w http.ResponseWriter, r *http.Request) {
	var (
		ctx    = r.Context()
		teamId = chi.URLParam(r, TEAM_ID)
	)
	result, err := t.teams.Get(ctx, teamId)
	if err != nil {
		t.Error(w, err)
		return
	}

	t.Succeed(w, result)
}

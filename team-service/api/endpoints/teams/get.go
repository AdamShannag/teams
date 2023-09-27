package teams

import (
	"github.com/go-chi/chi/v5"
	"net/http"
	"team-service/constant/key"
	"team-service/filter/team"
	page "team-service/service/pagination"
	"team-service/service/sorting"
)

func (t *Teams) GetTeams(w http.ResponseWriter, r *http.Request) {
	var (
		ctx        = r.Context()
		query      = r.URL.Query()
		filter     = team.NewFilter(query)
		sort       = sorting.NewSort(query)
		pagination = page.NewPagination(
			query.Get("page"),
			query.Get("size"),
		)
	)

	result, err := t.service.List(ctx, pagination, filter, sort)
	if err != nil {
		t.Error(w, err)
		return
	}

	t.Succeed(w, result)
}

func (t *Teams) GetTeam(w http.ResponseWriter, r *http.Request) {
	var (
		ctx    = r.Context()
		teamId = chi.URLParam(r, key.TEAM_ID)
	)
	result, err := t.service.Get(ctx, teamId)
	if err != nil {
		t.Error(w, err)
		return
	}

	t.Succeed(w, result)
}

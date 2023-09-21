package teams

import (
	"github.com/go-chi/chi/v5"
	"net/http"
)

func (t *Teams) DeleteTeam(w http.ResponseWriter, r *http.Request) {
	var (
		ctx    = r.Context()
		teamId = chi.URLParam(r, TEAM_ID)
	)

	if err := t.teams.Delete(ctx, teamId); err != nil {
		t.Error(w, err)
		return
	}

	t.Deleted(w)
}

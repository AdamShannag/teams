package teams

import (
	"net/http"

	"github.com/AdamShannag/toolkit/v2"
	"github.com/go-chi/chi/v5"
)

const (
	TEAM_ID = "teamId"
)

func (t *Teams) GetTeams(w http.ResponseWriter, r *http.Request) {
	t.tools.WriteJSON(w, http.StatusOK,
		toolkit.JSONResponse{
			Error:   false,
			Message: "success",
			Data:    "some teams list object",
		})
}

func (t *Teams) GetTeam(w http.ResponseWriter, r *http.Request) {
	teamId := chi.URLParam(r, TEAM_ID)

	t.tools.WriteJSON(w, http.StatusOK,
		toolkit.JSONResponse{
			Error:   false,
			Message: "found " + teamId,
			Data:    "some team object",
		})
}

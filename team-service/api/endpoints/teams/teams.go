package teams

import (
	"team-service/pkg/logger"

	"github.com/AdamShannag/toolkit/v2"
	"github.com/go-chi/chi/v5"
	"github.com/rs/zerolog"
)

type Teams struct {
	*chi.Mux
	l     zerolog.Logger
	tools *toolkit.Tools
}

func NewTeams(extras ...any) Teams {
	h := Teams{
		Mux:   chi.NewMux(),
		l:     logger.Get(),
		tools: &toolkit.Tools{},
	}

	h.Get("/", h.GetTeams)
	h.Get("/{teamId}", h.GetTeam)

	return h
}

package teams

import (
	"team-service/api/handler"
	"team-service/pkg/logger"
	"team-service/service/team"

	"github.com/AdamShannag/toolkit/v2"
	"github.com/go-chi/chi/v5"
	"github.com/rs/zerolog"
)

type Teams struct {
	*chi.Mux
	l zerolog.Logger
	*toolkit.Tools
	*handler.Handler
	teams team.Service
}

func NewTeams(teams team.Service) Teams {
	h := Teams{
		Mux:     chi.NewMux(),
		l:       logger.Get(),
		Tools:   &toolkit.Tools{},
		Handler: handler.NewHandler(&toolkit.Tools{}),
		teams:   teams,
	}

	h.Get("/", h.GetTeams)
	h.Get("/{teamId}", h.GetTeam)
	h.Post("/", h.Create)
	h.Put("/", h.Update)
	h.Delete("/{teamId}", h.DeleteTeam)

	return h
}

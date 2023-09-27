package teams

import (
	"team-service/api/handler"
	"team-service/pkg/logger"
	"team-service/service/team"

	"github.com/AdamShannag/toolkit/v2"
	"github.com/go-chi/chi/v5"
)

type Teams struct {
	*chi.Mux
	*toolkit.Tools
	*handler.Handler
	service team.Service
}

func NewTeams(teams team.Service) Teams {
	h := Teams{
		Mux:     chi.NewMux(),
		Tools:   &toolkit.Tools{},
		Handler: handler.NewHandler(&toolkit.Tools{}, logger.Get()),
		service: teams,
	}

	// with users info (gRpc)
	h.Get("/", h.GetTeams)
	// with users info (gRpc)
	h.Get("/{teamId}", h.GetTeam)
	h.Post("/", h.Create)
	h.Put("/", h.Update)
	h.Delete("/", h.DeleteTeam)

	return h
}

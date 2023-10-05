package teams

import (
	"github.com/nats-io/nats.go/jetstream"
	"github.com/rs/zerolog"
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
	jetstream.JetStream
	service team.Service
	log     zerolog.Logger
}

func NewTeams(teams team.Service, js jetstream.JetStream) Teams {
	log := logger.Get()
	h := Teams{
		Mux:       chi.NewMux(),
		Tools:     &toolkit.Tools{},
		JetStream: js,
		Handler:   handler.NewHandler(&toolkit.Tools{}, log),
		service:   teams,
		log:       log,
	}

	h.Get("/", h.GetTeams)
	h.Get("/{teamId}", h.GetTeam)
	h.Post("/", h.Create)
	h.Put("/", h.Update)
	h.Delete("/", h.DeleteTeam)

	return h
}

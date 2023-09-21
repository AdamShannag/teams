package users

import (
	"user-service/pkg/koj/kmid"
	"user-service/pkg/logger"
	"user-service/service/userservice"

	"github.com/AdamShannag/toolkit/v2"
	"github.com/go-chi/chi/v5"
	"github.com/nats-io/nats.go/jetstream"
	"github.com/rs/zerolog"
)

type Users struct {
	*chi.Mux
	jetstream.JetStream
	*toolkit.Tools
	zerolog.Logger
	keycloakUserService userservice.KeycloakUserService
}

func NewUsers(js jetstream.JetStream, keycloakUserService userservice.KeycloakUserService) Users {
	h := Users{
		Mux:                 chi.NewMux(),
		JetStream:           js,
		Tools:               &toolkit.Tools{},
		Logger:              logger.Get(),
		keycloakUserService: keycloakUserService,
	}

	h.Group(func(r chi.Router) {
		r.Use(kmid.Roles("admin"))
		r.Post("/", h.Create)
	})

	return h
}

package user

import (
	"github.com/AdamShannag/toolkit/v2"
	"github.com/go-chi/chi/v5"
	"team-service/api/handler"
	"team-service/pkg/logger"
	"team-service/service/user"
)

type User struct {
	*chi.Mux
	*toolkit.Tools
	*handler.Handler
	service user.Service
}

func NewUser(service user.Service) User {
	h := User{
		Mux:     chi.NewMux(),
		Tools:   &toolkit.Tools{},
		Handler: handler.NewHandler(&toolkit.Tools{}, logger.Get()),
		service: service,
	}

	h.Post("/assign", h.Assign)
	h.Put("/assign", h.AssignApproval)

	return h
}

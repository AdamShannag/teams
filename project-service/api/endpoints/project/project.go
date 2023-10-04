package project

import (
	"github.com/AdamShannag/toolkit/v2"
	"github.com/go-chi/chi/v5"
	"project-service/api/handler"
	"project-service/pkg/logger"
	"project-service/service/project"
)

type Project struct {
	*chi.Mux
	*toolkit.Tools
	*handler.Handler
	service project.Service
}

func NewProject(service project.Service) Project {
	h := Project{
		Mux:     chi.NewMux(),
		Tools:   &toolkit.Tools{},
		Handler: handler.NewHandler(&toolkit.Tools{}, logger.Get()),
		service: service,
	}

	h.Get("/", h.GetProjects)
	h.Get("/{projectId}", h.GetProject)
	h.Post("/", h.Create)
	h.Put("/", h.Update)
	h.Delete("/", h.DeleteProject)

	return h
}

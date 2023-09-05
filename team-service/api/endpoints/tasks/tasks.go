package tasks

import (
	"team-service/pkg/logger"

	"github.com/AdamShannag/toolkit/v2"
	"github.com/go-chi/chi/v5"
	"github.com/rs/zerolog"
)

type Tasks struct {
	*chi.Mux
	l     zerolog.Logger
	tools *toolkit.Tools
}

func NewTasks(extras ...any) Tasks {
	h := Tasks{
		Mux:   chi.NewMux(),
		l:     logger.Get(),
		tools: &toolkit.Tools{},
	}

	h.Get("/", h.GetTasks)
	h.Get("/{taskId}", h.GetTask)

	return h
}

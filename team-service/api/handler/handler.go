package handler

import (
	"github.com/AdamShannag/toolkit/v2"
	"net/http"
	"team-service/validation/violation"
)

type Handler struct {
	tools *toolkit.Tools
}

func NewHandler(tools *toolkit.Tools) *Handler {
	return &Handler{
		tools: tools,
	}
}

func (h *Handler) Render(w http.ResponseWriter, body interface{}, message string, status ...int) {
	statusCode := http.StatusOK
	if len(status) > 0 {
		statusCode = status[0]
	}
	if _, ok := body.(error); ok {
		h.tools.ErrorJSON(w, body.(error))
		return
	}

	h.tools.WriteJSON(w, statusCode, toolkit.JSONResponse{
		Error:   false,
		Message: message,
		Data:    body,
	})

}

func (h *Handler) Succeed(w http.ResponseWriter, body interface{}) {
	h.Render(w, body, "success")
}

func (h *Handler) Created(w http.ResponseWriter, body interface{}) {
	h.Render(w, body, "created", http.StatusCreated)
}

func (h *Handler) Updated(w http.ResponseWriter, body interface{}) {
	h.Render(w, body, "updated")
}

func (h *Handler) Deleted(w http.ResponseWriter) {
	h.Render(w, nil, "deleted")
}

func (h *Handler) Error(w http.ResponseWriter, err error) {
	h.Render(w, err, "")
}

func (h *Handler) ErrorViolation(w http.ResponseWriter, violations []violation.Violation) {
	h.Render(w, violations, "error", http.StatusBadRequest)
}

package handler

import (
	"fmt"
	"github.com/AdamShannag/toolkit/v2"
	"github.com/rs/zerolog"
	"net/http"
	"project-service/constant"
	"project-service/validation/violation"
)

type Handler struct {
	*toolkit.Tools
	log zerolog.Logger
}

func NewHandler(tools *toolkit.Tools, log zerolog.Logger) *Handler {
	return &Handler{
		Tools: tools,
		log:   log,
	}
}

func (h *Handler) Render(w http.ResponseWriter, body interface{}, message string, status ...int) {
	statusCode := http.StatusOK
	if len(status) > 0 {
		statusCode = status[0]
	}

	if _, ok := body.(error); ok {
		_ = h.ErrorJSON(w, body.(error))
		return
	}

	_ = h.WriteJSON(w, statusCode, toolkit.JSONResponse{
		Error:   statusCode >= 300,
		Message: message,
		Data:    body,
	})

}

func (h *Handler) SucceedF(w http.ResponseWriter, body string, a ...any) {
	message := fmt.Sprintf(body, a...)
	h.Succeed(w, message)
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
	h.ErrorViolation(w, violation.FieldViolation("noField", err))
}

func (h *Handler) ErrorParsing(w http.ResponseWriter, err error) {
	h.log.Error().Err(err).Msgf(constant.FAILED_PARSING_MESSAGE)
	h.ErrorViolation(w, violation.FieldViolation("request", err))
}

func (h *Handler) ErrorViolation(w http.ResponseWriter, vio violation.Violation) {
	h.ErrorViolations(w, []violation.Violation{vio})
}

func (h *Handler) ErrorViolations(w http.ResponseWriter, violations []violation.Violation) {
	h.Render(w, violations, "error", http.StatusBadRequest)
}

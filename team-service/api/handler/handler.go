package handler

import (
	"fmt"
	"github.com/AdamShannag/toolkit/v2"
	"github.com/rs/zerolog"
	"net/http"
	"team-service/constant"
	"team-service/validation/violation"
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

	response := Response{
		Message:  message,
		Severity: SUCCESS,
		Summary:  "Success",
		Payload:  body,
	}

	if statusCode >= 300 {
		response.Severity = FAILED
		response.Summary = "Failed"
	}

	_ = h.WriteJSON(w, statusCode, response)
}

func (h *Handler) SuccessF(w http.ResponseWriter, message string, a ...any) {
	m := fmt.Sprintf(message, a...)
	h.Success(w, nil, m)
}

func (h *Handler) Success(w http.ResponseWriter, body interface{}, message ...string) {
	h.Render(w, body, getMessage(message))
}

func (h *Handler) Created(w http.ResponseWriter, body interface{}, entity string) {
	message := fmt.Sprintf(MSG_CREATED, entity)
	h.Render(w, body, message, http.StatusCreated)
}

func (h *Handler) Updated(w http.ResponseWriter, body interface{}, entity string) {
	message := fmt.Sprintf(MSG_UPDATED, entity)
	h.Render(w, body, message)
}

func (h *Handler) Deleted(w http.ResponseWriter, entity string) {
	message := fmt.Sprintf(MSG_DELETED, entity)
	h.Render(w, nil, message)
}

func (h *Handler) Error(w http.ResponseWriter, err error) {
	h.ErrorViolation(w, violation.FieldViolation("noField", err), err.Error())
}

func (h *Handler) ErrorParsing(w http.ResponseWriter, err error) {
	message := constant.FAILED_PARSING_MESSAGE
	h.log.Error().Err(err).Msgf(message)
	h.ErrorViolation(w, violation.FieldViolation("request", err), message)
}

func (h *Handler) ErrorViolation(w http.ResponseWriter, vio violation.Violation, message string) {
	h.ErrorViolations(w, []violation.Violation{vio}, message)
}

func (h *Handler) ErrorViolations(w http.ResponseWriter, violations []violation.Violation, message string) {
	h.Render(w, violations, message, http.StatusBadRequest)
}

func getMessage(message []string) string {
	if len(message) == 0 {
		return ""
	}
	return message[0]
}

package tasks

import (
	"net/http"

	"github.com/AdamShannag/toolkit/v2"
	"github.com/go-chi/chi/v5"
)

const (
	TASK_ID string = "taskId"
)

func (t *Tasks) GetTasks(w http.ResponseWriter, r *http.Request) {
	t.tools.WriteJSON(w, http.StatusOK,
		toolkit.JSONResponse{
			Error:   false,
			Message: "success",
			Data:    "some tasks list object",
		})
}

func (t *Tasks) GetTask(w http.ResponseWriter, r *http.Request) {
	taskId := chi.URLParam(r, TASK_ID)

	var taskRequest TaskRequest

	err := t.tools.ReadJSON(w, r, &taskRequest)
	if err != nil {
		t.l.Error().Msg("Team Id not found")
		t.tools.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}

	t.tools.WriteJSON(w, http.StatusOK,
		toolkit.JSONResponse{
			Error:   false,
			Message: "found " + taskId,
			Data:    "some task object",
		})
}

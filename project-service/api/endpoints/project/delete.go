package project

import (
	"net/http"
	"project-service/resource/project"
)

func (t *Project) DeleteProject(w http.ResponseWriter, r *http.Request) {
	var (
		ctx     = r.Context()
		request project.DeleteRequest
	)

	if err := t.ReadJSON(w, r, &request); err != nil {
		t.ErrorParsing(w, err)
		return
	}

	if violations := t.service.Delete(ctx, request); violations != nil {
		t.ErrorViolations(w, violations)
		return
	}

	t.Deleted(w)
}

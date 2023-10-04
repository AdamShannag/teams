package project

import (
	"net/http"
	"project-service/resource/project"
)

func (t *Project) Update(w http.ResponseWriter, r *http.Request) {
	var (
		ctx     = r.Context()
		request project.UpdateRequest
	)

	if err := t.ReadJSON(w, r, &request); err != nil {
		t.ErrorParsing(w, err)
		return
	}

	updated, violations := t.service.Update(ctx, &request)

	if violations != nil {
		t.ErrorViolations(w, violations)
		return
	}

	t.Updated(w, updated)
}

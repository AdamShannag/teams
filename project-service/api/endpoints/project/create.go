package project

import (
	"net/http"
	"project-service/constant"
	"project-service/resource/project"
)

func (t *Project) Create(w http.ResponseWriter, r *http.Request) {
	var (
		ctx       = r.Context()
		projectId = r.Header.Get(constant.PROJECT_ID)
		request   project.Request
	)

	if err := t.ReadJSON(w, r, &request); err != nil {
		t.ErrorParsing(w, err)
		return
	}

	created, violations := t.service.Create(ctx, &request, projectId)

	if violations != nil {
		t.ErrorViolations(w, violations)
		return
	}

	t.Created(w, created)
}

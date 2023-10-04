package project

import (
	"github.com/go-chi/chi/v5"
	"net/http"
	"project-service/constant"
	"project-service/query/project"
)

func (t *Project) GetProjects(w http.ResponseWriter, r *http.Request) {
	var (
		ctx   = r.Context()
		query = project.NewQuery(r)
	)

	result, err := t.service.List(ctx, *query)
	if err != nil {
		t.Error(w, err)
		return
	}

	t.Succeed(w, result)
}

func (t *Project) GetProject(w http.ResponseWriter, r *http.Request) {
	var (
		ctx       = r.Context()
		projectId = chi.URLParam(r, constant.PROJECT_ID)
	)
	result, err := t.service.Get(ctx, projectId)
	if err != nil {
		t.Error(w, err)
		return
	}

	t.Succeed(w, result)
}

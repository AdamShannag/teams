package teams

import (
	"net/http"
	"team-service/resource/team"
)

func (t *Teams) Create(w http.ResponseWriter, r *http.Request) {
	var (
		ctx     = r.Context()
		request team.Request
	)

	if err := t.ReadJSON(w, r, &request); err != nil {
		t.Error(w, err)
		return
	}

	created, violations := t.teams.Create(ctx, &request)

	if violations != nil {
		t.ErrorViolation(w, violations)
		return
	}

	t.Created(w, created)
}

package teams

import (
	"net/http"
	"team-service/resource/team"
)

func (t *Teams) Update(w http.ResponseWriter, r *http.Request) {
	var (
		ctx     = r.Context()
		request team.UpdateRequest
	)

	if err := t.ReadJSON(w, r, &request); err != nil {
		t.Error(w, err)
		return
	}

	updated, violations := t.teams.Update(ctx, &request)

	if violations != nil {
		t.ErrorViolation(w, violations)
		return
	}

	t.Updated(w, updated)
}

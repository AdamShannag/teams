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

	if err := t.tools.ReadJSON(w, r, &request); err != nil {
		t.handler.Error(w, err)
		return
	}

	created, err := t.teams.Create(ctx, &request)

	if err != nil {
		t.handler.Render(w, err, "error", http.StatusBadRequest)
		return
	}

	t.handler.Created(w, created)
}

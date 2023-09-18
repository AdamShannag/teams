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

	if err := t.tools.ReadJSON(w, r, &request); err != nil {
		t.handler.Error(w, err)
		return
	}

	updated, err := t.teams.Update(ctx, &request)

	if err != nil {
		t.handler.Error(w, err)
		return
	}

	t.handler.Updated(w, updated)
}

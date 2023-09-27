package teams

import (
	"net/http"
	"team-service/resource/team"
)

func (t *Teams) DeleteTeam(w http.ResponseWriter, r *http.Request) {
	var (
		ctx     = r.Context()
		request team.DeleteRequest
	)

	if err := t.ReadJSON(w, r, &request); err != nil {
		t.ErrorParsing(w, err)
		return
	}

	if err := t.service.Delete(ctx, request); err != nil {
		t.Error(w, err)
		return
	}

	t.Deleted(w)
}

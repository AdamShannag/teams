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

	if violations := t.service.Delete(ctx, request); violations != nil {
		t.ErrorViolations(w, violations)
		return
	}

	t.Deleted(w)
}

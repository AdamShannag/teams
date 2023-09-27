package teams

import (
	"net/http"
	"team-service/constant"
	"team-service/resource/team"
)

func (t *Teams) Create(w http.ResponseWriter, r *http.Request) {
	var (
		ctx     = r.Context()
		userId  = r.Header.Get(constant.USER_ID)
		request team.Request
	)

	if err := t.ReadJSON(w, r, &request); err != nil {
		t.ErrorParsing(w, err)
		return
	}

	created, violations := t.service.Create(ctx, &request, userId)

	if violations != nil {
		t.ErrorViolations(w, violations)
		return
	}

	t.Created(w, created)
}

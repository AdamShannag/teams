package teams

import (
	"github.com/nats-io/nats.go"
	"net/http"
	"team-service/config"
	"team-service/pkg/koj/kmid"
	"team-service/pkg/nts"
	"team-service/resource/team"
)

func (t *Teams) Create(w http.ResponseWriter, r *http.Request) {
	var (
		ctx     = r.Context()
		userId  = r.Context().Value(kmid.USER_ID_KEY).(string)
		request team.Request
	)

	if err := t.ReadJSON(w, r, &request); err != nil {
		t.ErrorParsing(w, err)
		return
	}

	created, violations := t.service.Create(ctx, &request, userId)

	if violations != nil {
		t.ErrorViolations(w, violations, "error while creating team")
		return
	}

	t.publish(r, created)

	t.Created(w, created, "team")
}

func (t *Teams) publish(r *http.Request, created *team.Resource) {
	_, err := nts.Publish(r.Context(), t, &nats.Msg{
		Subject: config.Get().TeamsSubjectNew,
		Data:    []byte(created.ID),
	})

	if err != nil {
		t.log.Error().Err(err).Msg("an error has occurred while publish")
	}
}

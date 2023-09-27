package member

import (
	"net/http"
	"team-service/constant/key"
	"team-service/resource/member"
)

func (t *Member) Assign(w http.ResponseWriter, r *http.Request) {
	var (
		ctx     = r.Context()
		userId  = r.Header.Get(key.USER_ID)
		request member.AssignRequest
	)
	if err := t.ReadJSON(w, r, &request); err != nil {
		t.ErrorParsing(w, err)
		return
	}

	violations := t.service.Assign(ctx, &request, userId)

	if violations != nil {
		t.ErrorViolations(w, violations)
		return
	}

	t.SucceedF(w, "%sassigned %s to team [%s]", getPrefixAssignMessage(request.IsAssign()), request.Members, request.TeamId)
}

func (t *Member) AssignApproval(w http.ResponseWriter, r *http.Request) {
	var (
		ctx     = r.Context()
		userId  = r.Header.Get(key.USER_ID)
		request member.ApprovalRequest
	)

	if err := t.ReadJSON(w, r, &request); err != nil {
		t.ErrorParsing(w, err)
		return
	}

	violations := t.service.AssignApproval(ctx, &request, userId)

	if violations != nil {
		t.ErrorViolations(w, violations)
		return
	}

	t.SucceedF(w, "%s assign %s to team", getPrefixApprovalMessage(request.IsApprove()), request.Members)
}

func getPrefixAssignMessage(assign bool) string {
	if assign {
		return ""
	}
	return "un"
}

func getPrefixApprovalMessage(approve bool) string {
	if approve {
		return "approved"
	}
	return "rejected"
}

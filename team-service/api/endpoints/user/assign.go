package user

import (
	"fmt"
	"net/http"
	"team-service/pkg/koj/kmid"
	"team-service/resource/user"
)

func (t *User) Assign(w http.ResponseWriter, r *http.Request) {
	var (
		ctx     = r.Context()
		userId  = r.Context().Value(kmid.USER_ID_KEY).(string)
		request user.AssignRequest
	)
	if err := t.ReadJSON(w, r, &request); err != nil {
		t.ErrorParsing(w, err)
		return
	}

	violations := t.service.Assign(ctx, &request, userId)
	prefixAssignMessage := getPrefixAssignMessage(request.IsAssign())

	if violations != nil {
		t.ErrorViolations(w, violations, fmt.Sprintf("error while %sassigning member", prefixAssignMessage))
		return
	}

	t.SuccessF(w, "%sassigned user to team", prefixAssignMessage)
}

func (t *User) AssignApproval(w http.ResponseWriter, r *http.Request) {
	var (
		ctx     = r.Context()
		userId  = r.Context().Value(kmid.USER_ID_KEY).(string)
		request user.ApprovalRequest
	)

	if err := t.ReadJSON(w, r, &request); err != nil {
		t.ErrorParsing(w, err)
		return
	}

	violations := t.service.AssignApproval(ctx, &request, userId)
	prefixApprovalMessage := getPrefixApprovalMessage(request.IsApprove())

	if violations != nil {
		t.ErrorViolations(w, violations, fmt.Sprintf("error while approving the assignment"))
		return
	}

	t.SuccessF(w, "%s assign user to team", prefixApprovalMessage)
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

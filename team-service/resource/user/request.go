package user

type AssignRequest struct {
	Assign *bool    `json:"assign,omitempty"`
	TeamId string   `json:"teamId"`
	Users  []string `json:"users"`
}

type ApprovalRequest struct {
	Approve *bool    `json:"approve,omitempty"`
	Users   []string `json:"users"`
}

func (a *AssignRequest) IsAssign() bool {
	return a.Assign == nil || *a.Assign
}

func (r *ApprovalRequest) IsApprove() bool {
	return r.Approve == nil || *r.Approve
}

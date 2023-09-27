package member

type AssignRequest struct {
	Assign  *bool    `json:"assign,omitempty"`
	TeamId  string   `json:"teamId"`
	Members []string `json:"members"`
}

type ApprovalRequest struct {
	Approve *bool    `json:"approve,omitempty"`
	Members []string `json:"members"`
}

func (a *AssignRequest) IsAssign() bool {
	return a.Assign == nil || *a.Assign
}

func (r *ApprovalRequest) IsApprove() bool {
	return r.Approve == nil || *r.Approve
}

package handler

type Response struct {
	Message  string      `json:"message,omitempty"`
	Severity Severity    `json:"severity,omitempty"`
	Summary  string      `json:"summary,omitempty"`
	Payload  interface{} `json:"payload,omitempty"`
}

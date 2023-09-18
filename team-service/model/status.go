package model

type Status string

const (
	NEW     Status = "NEW"
	DELETED Status = "DELETED"
	ACTIVE  Status = "ACTIVE"
)

func AllStatus() []Status {
	return []Status{NEW, DELETED, ACTIVE}
}

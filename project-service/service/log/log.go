package log

import "github.com/rs/zerolog"

type Log interface {
	Success(string, ...any)
	Failed(string, error, ...string)
}

type Service struct {
	Create
	Retrieve
	Update
	Delete
	List
	Assign
}

func NewLogService(log zerolog.Logger) Service {
	return Service{
		Create:   Create{log},
		Retrieve: Retrieve{log},
		Update:   Update{log},
		Delete:   Delete{log},
		List:     List{log},
		Assign:   Assign{log},
	}
}

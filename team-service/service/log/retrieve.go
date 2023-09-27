package log

import "github.com/rs/zerolog"

type Retrieve struct {
	l zerolog.Logger
}

func (l *Retrieve) Success(entity string, reference ...any) {
	l.l.Info().Msgf(RETRIEVED_SUCCESSFULLY, entity, reference)
}

func (l *Retrieve) Failed(entity string, err error, reference ...string) {
	l.l.Error().Err(err).Msgf(RETRIEVED_FAILED, entity, reference)
}

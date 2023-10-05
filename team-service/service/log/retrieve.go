package log

import "github.com/rs/zerolog"

type Retrieve struct {
	l zerolog.Logger
}

func (l *Retrieve) Success(entity string, reference ...any) {
	l.l.Info().Any("uuid", reference[0]).Msgf(RETRIEVED_SUCCESSFULLY, entity)
}

func (l *Retrieve) Failed(entity string, err error, reference ...string) {
	l.l.Error().Err(err).Str("uuid", reference[0]).Msgf(RETRIEVED_FAILED, entity)
}

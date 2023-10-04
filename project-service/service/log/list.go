package log

import "github.com/rs/zerolog"

type List struct {
	l zerolog.Logger
}

func (l *List) Success(entity string, reference ...any) {
	l.l.Info().Msgf(RETRIEVED_LIST_SUCCESSFULLY, entity, reference)
}

func (l *List) Failed(entity string, err error, reference ...string) {
	l.l.Error().Err(err).Msgf(RETRIEVED_LIST_FAILED, entity, reference)
}

package log

import "github.com/rs/zerolog"

type List struct {
	l zerolog.Logger
}

func (l *List) Success(entity string, reference ...any) {
	l.l.Info().Any("size", reference[0]).Msgf(RETRIEVED_LIST_SUCCESSFULLY, entity)
}

func (l *List) Failed(entity string, err error, _ ...string) {
	l.l.Error().Err(err).Msgf(RETRIEVED_LIST_FAILED, entity)
}

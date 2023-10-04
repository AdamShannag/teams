package log

import "github.com/rs/zerolog"

type Update struct {
	l zerolog.Logger
}

func (l *Update) Success(entity string, reference ...any) {
	l.l.Info().Msgf(UPDATED_SUCCESSFULLY, entity, reference)
}

func (l *Update) Failed(entity string, err error, reference ...string) {
	l.l.Error().Err(err).Msgf(UPDATED_FAILED, entity, reference)
}

package log

import "github.com/rs/zerolog"

type Update struct {
	l zerolog.Logger
}

func (l *Update) Success(entity string, reference ...any) {
	l.l.Info().Any("uuid", reference[0]).Msgf(UPDATED_SUCCESSFULLY, entity)
}

func (l *Update) Failed(entity string, err error, reference ...string) {
	l.l.Error().Err(err).Str("uuid", reference[0]).Msgf(UPDATED_FAILED, entity)
}

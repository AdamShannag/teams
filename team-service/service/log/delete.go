package log

import "github.com/rs/zerolog"

type Delete struct {
	l zerolog.Logger
}

func (l *Delete) Success(entity string, reference ...any) {
	l.l.Info().Any("uuid's", reference).Msgf(DELETED_SUCCESSFULLY, entity, reference)
}

func (l *Delete) Failed(entity string, err error, reference ...string) {
	l.l.Error().Err(err).Any("uuid's", reference).Msgf(DELETED_FAILED, entity, reference)
}

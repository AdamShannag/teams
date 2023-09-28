package log

import "github.com/rs/zerolog"

type Create struct {
	l zerolog.Logger
}

func (l *Create) Success(entity string, reference ...any) {
	l.l.Info().Msgf(CREATED_SUCCESSFULLY, entity, reference)
}

func (l *Create) Failed(entity string, err error, _ ...string) {
	l.l.Error().Err(err).Msgf(CREATED_FAILED, entity)
}

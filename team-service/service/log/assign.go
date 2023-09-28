package log

import "github.com/rs/zerolog"

type Assign struct {
	l zerolog.Logger
}

func (l *Assign) Success(message string, _ ...any) {
	l.l.Info().Msgf(message)
}

func (l *Assign) Failed(message string, err error, _ ...string) {
	l.l.Error().Err(err).Msgf(message)
}

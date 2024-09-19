package logger

import (
	"github.com/rs/zerolog"
)

func SetupZeroLogger() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
}

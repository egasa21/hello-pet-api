package logger

import (
	"github.com/egasa21/hello-pet-api/helpers"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"
	"strings"
)

var logger zerolog.Logger

func init() {
	logger = log.With().Timestamp().CallerWithSkipFrameCount(3).Stack().Logger()
}

func SetLogLevel(level string) {
	var l zerolog.Level

	switch strings.ToLower(level) {
	case "error":
		l = zerolog.ErrorLevel
	case "warn":
		l = zerolog.ErrorLevel
	case "info":
		l = zerolog.ErrorLevel
	case "debug":
		l = zerolog.ErrorLevel
	default:
		l = zerolog.InfoLevel
	}
	zerolog.SetGlobalLevel(l)
}

func Info(message string, args ...interface{}) {
	log.Info().Msgf(message, args...)
}

func DebugJson(message string, data interface{}) {
	log.Info().RawJSON(message, helpers.ServeJSON(data)).Msg("")
}

func Debug(message string, args ...interface{}) {
	logger.Debug().Msgf(message, args...)
}

func Warn(message string, args ...interface{}) {
	logger.Warn().Msgf(message, args...)
}

func Error(message string, args ...interface{}) {
	logger.Error().Msgf(message, args...)
}

func Fatal(message string, args ...interface{}) {
	logger.Fatal().Msgf(message, args...)
	os.Exit(1)
}

func Log(message string, args ...interface{}) {
	if len(args) == 0 {
		log.Info().Msg(message)
	} else {
		log.Info().Msgf(message, args...)
	}
}

package logger

import (
	"os"
	"strconv"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

var GlobalLogLevel zerolog.Level

func getLogLevel(logLevel string) zerolog.Level {
	switch logLevel {
	case "debug":
		return zerolog.DebugLevel
	case "info":
		return zerolog.InfoLevel
	case "warn":
		return zerolog.WarnLevel
	case "error":
		return zerolog.ErrorLevel
	case "trace":
		return zerolog.TraceLevel
	default:
		return zerolog.TraceLevel
	}
}

func InitLogger(globalLogLevel string) {
	GlobalLogLevel = getLogLevel(globalLogLevel)

	zerolog.SetGlobalLevel(GlobalLogLevel)
	zerolog.TimestampFieldName = "t"
	if GlobalLogLevel == zerolog.DebugLevel {
		multi := zerolog.MultiLevelWriter(os.Stdout)
		log.Logger = zerolog.New(multi).With().Caller().Timestamp().Logger().Output(zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339})
	} else if os.Getenv("GIN_MODE") != "release" {
		log.Logger = log.With().Caller().Logger()
	}

	log.Info().Msg("Global level set to " + strconv.FormatInt(int64(GlobalLogLevel), 10))
}

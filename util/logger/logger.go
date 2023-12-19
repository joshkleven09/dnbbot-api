package logger

import (
	"github.com/rs/zerolog"
	"github.com/spf13/viper"
	"os"
	"strings"
)

//todo eventually handle gin logs

func New() *zerolog.Logger {
	logLevel := zerolog.InfoLevel

	if strings.ToLower(viper.GetString("logger.level")) == "debug" {
		logLevel = zerolog.TraceLevel
	}

	zerolog.SetGlobalLevel(logLevel)
	logger := zerolog.New(os.Stdout).With().Timestamp().Caller().Logger()

	if strings.ToLower(viper.GetString("env")) == "local" {
		logger = logger.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	}

	return &logger
}

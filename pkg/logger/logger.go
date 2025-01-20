package logger

import (
	"github.com/nathaliaguayos/msg-receiver/config"
	"github.com/nathaliaguayos/msg-receiver/internal/version"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/pkgerrors"
	"os"
)

func New(config *config.Config) *zerolog.Logger {
	configureZerolog(config)
	log := zerolog.New(os.Stdout).With().Timestamp().
		Str("service", config.ServiceName).
		Str("version", version.VERSION).
		Logger()
	return &log
}

func configureZerolog(config *config.Config) {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	lvl, err := zerolog.ParseLevel(config.LogLevel)
	if err != nil {
		panic("invalid log level")
	}
	zerolog.SetGlobalLevel(lvl)
}

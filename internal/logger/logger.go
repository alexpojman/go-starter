package logger

import (
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

const (
	Development string = "development"
	Production  string = "production"
)

func InitLogger(environment string) {
	logger := zerolog.New(os.Stdout).With().Timestamp().Logger()

	if environment == Development {
		log.Logger = logger.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	} else {
		log.Logger = logger
	}

}

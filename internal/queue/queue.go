package queue

import (
	"github.com/rs/zerolog"
)

func Queue(logger zerolog.Logger) {
	logger.Info().Msg("Adding to queue")
}

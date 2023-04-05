package log

import (
	"io"
	"sync"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/rs/zerolog/pkgerrors"
)

var once sync.Once

func New() {
	once.Do(func() {
		zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack
		zerolog.TimeFieldFormat = time.RFC3339Nano

		log.Logger = zerolog.New(Console()).
			Level(zerolog.Level(zerolog.InfoLevel)).
			With().
			Timestamp().
			Caller().
			Logger()
	})
}

func SetOutput(writers ...io.Writer) {
	log.Logger = log.Logger.Output(zerolog.MultiLevelWriter(writers...))
}

func SetLogLevel(level int) {
	log.Logger = log.Logger.Level(zerolog.Level(level))
}

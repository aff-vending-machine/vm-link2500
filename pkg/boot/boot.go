package boot

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/rs/zerolog/log"
)

type (
	TerminateFn = func(context.Context)
	CloseFn     = func() error
)

type bootContext struct {
	osSignal     chan os.Signal
	terminateFn  []TerminateFn
	isTerminated bool
}

// Create global context to use in anywhere
var bootCtx = &bootContext{
	osSignal:     make(chan os.Signal, 1),
	terminateFn:  make([]TerminateFn, 0),
	isTerminated: false,
}

func Init(conf interface{}) {
	signal.Notify(
		bootCtx.osSignal,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT,
	)

	go bootCtx.waitForSignal()
}

func AddTerminateFn(fn TerminateFn) {
	bootCtx.terminateFn = append(bootCtx.terminateFn, fn)
}

func AddCloseFn(close CloseFn) {
	AddTerminateFn(func(context.Context) {
		if err := close(); err != nil {
			log.Error().Err(err).Msg("close service failed")
		}
	})
}

func TerminateWhenError(err error, clues ...interface{}) {
	if err != nil {
		log.Error().CallerSkipFrame(1).Err(err).Any("clues", clues).Msg("critical error")

		bootCtx.osSignal <- syscall.SIGABRT
		bootCtx.isTerminated = true
	}
}

func Serve() {
	for {
		if bootCtx.isTerminated {
			break
		}
		// Sleep for 1 second.
		time.Sleep(time.Second)
	}
}

func (b *bootContext) waitForSignal() {
	signal := <-b.osSignal

	for _, fn := range b.terminateFn {
		fn(context.TODO())
	}

	b.isTerminated = true
	code := handleSignal(signal)
	os.Exit(code)
}

func handleSignal(signal os.Signal) int {
	// Print the signal name and description.
	log.Info().Str("signal", signal.String()).Msgf("terminate signal received")

	switch signal {
	case syscall.SIGABRT:
		return 1
	}

	return 0
}

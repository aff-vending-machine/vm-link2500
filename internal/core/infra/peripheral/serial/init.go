package serial

import (
	"context"
	"fmt"
	"sync"
	"vm-link2500/pkg/helpers/errs"

	"github.com/rs/zerolog/log"
	"github.com/tarm/serial"
)

type SerialPort struct {
	*serial.Port
	mtx  sync.Mutex
	done chan struct{}
}

func OpenPort(config *serial.Config) (*SerialPort, error) {
	port, err := serial.OpenPort(config)
	if err != nil {
		return nil, err
	}

	return &SerialPort{port, sync.Mutex{}, make(chan struct{})}, nil
}

func (s *SerialPort) Close() {
	close(s.done)
	if s.Port != nil {
		if err := s.Port.Close(); errs.NoMsg(err, "file already closed") {
			log.Error().Err(err).Msg("failed to close port")
			return
		}
	}

	s.Port = nil
}

func (s *SerialPort) Write(ctx context.Context, payload []byte) error {
	s.mtx.Lock()
	defer s.mtx.Unlock()

	select {
	case <-ctx.Done():
		// Cancelled from outside, close the port
		if s.Port != nil {
			s.Port.Close()
		}
		if err := ctx.Err(); err != nil {
			return fmt.Errorf(err.Error())
		}
		return fmt.Errorf("cancelled")

	case <-s.done:
		// Done channel is closed
		return fmt.Errorf("channel closed")

	default:
		_, err := s.Port.Write(payload)
		return err
	}
}

func (s *SerialPort) Read(ctx context.Context, payload []byte) (int, error) {
	s.mtx.Lock()
	defer s.mtx.Unlock()

	readCtx, cancelled := context.WithCancel(context.Background())
	go func() {
		select {
		case <-ctx.Done():
			// Cancelled from outside, close the port
			if s.Port != nil {
				s.Port.Close()
			}
			if err := ctx.Err(); err != nil {
				log.Warn().Str("reason", err.Error()).Msg("port is closed")
				return
			}
			log.Info().Msg("port is closed by user")

		case <-s.done:
			// Done channel is closed
			log.Warn().Str("reason", "channel is closed").Msg("port is closed")

		case <-readCtx.Done():
			// Read context is completed
			log.Info().Msg("read serial data is completed")
		}
	}()

	n, err := s.Port.Read(payload)

	cancelled()

	return n, err
}

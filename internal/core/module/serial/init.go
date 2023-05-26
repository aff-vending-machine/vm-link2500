package serial

import (
	"context"
	"fmt"
	"sync"
	"vm-link2500/pkg/utils/errs"

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
	if err := s.Port.Close(); errs.Not(err, "file already closed") {
		log.Error().Err(err).Msg("failed to close port")
		return
	}

	s.Port = nil
}

func (s *SerialPort) Write(ctx context.Context, payload []byte) error {
	s.mtx.Lock()
	defer s.mtx.Unlock()

	select {
	case <-ctx.Done():
		// Cancelled, close the port
		s.Close()
		return fmt.Errorf("cancelled")

	case <-s.done:
		// Done channel is closed, close the port
		s.Close()
		return fmt.Errorf("cancelled")

	default:
		_, err := s.Port.Write(payload)
		return err
	}
}

func (s *SerialPort) Read(ctx context.Context, payload []byte) (int, error) {
	s.mtx.Lock()
	defer s.mtx.Unlock()

	go func() {
		<-ctx.Done()
		s.Port.Close()
		log.Warn().Str("reason", ctx.Err().Error()).Msg("port is closed")
	}()

	return s.Port.Read(payload)
}

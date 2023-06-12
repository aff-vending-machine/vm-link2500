package serial

import (
	"context"
	"fmt"
	"sync"

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

func (s *SerialPort) Close() error {
	s.mtx.Lock()
	defer s.mtx.Unlock()

	close(s.done)

	if s.Port != nil {
		if err := s.Port.Close(); err != nil {
			log.Error().Err(err).Msg("failed to close port")
			return fmt.Errorf("failed to close port: %w", err)
		}
		s.Port = nil
	}

	return nil
}

func (s *SerialPort) Write(ctx context.Context, payload []byte) error {
	s.mtx.Lock()
	defer s.mtx.Unlock()

	select {
	case <-ctx.Done():
		// Cancelled from outside
		return fmt.Errorf("write operation cancelled: %w", ctx.Err())

	case <-s.done:
		// Done channel is closed
		return fmt.Errorf("write operation failed because the done channel is closed")

	default:
		// Neither ctx.Done() nor s.done is ready, continue to the Write operation
		_, err := s.Port.Write(payload)
		if err != nil {
			return fmt.Errorf("failed to write payload: %w", err)
		}
		return nil
	}
}

func (s *SerialPort) Read(ctx context.Context, payload []byte) (int, error) {
	readCtx, cancel := context.WithCancel(ctx)
	defer cancel()

	s.mtx.Lock()
	defer s.mtx.Unlock()

	// Create channels to handle read result
	res := make(chan int, 1)
	errs := make(chan error, 1)

	// Start a goroutine to read data
	go func() {
		n, err := s.Port.Read(payload)
		if err != nil {
			errs <- err
		} else {
			res <- n
		}
		cancel() // Cancel the readCtx when we're done
	}()

	// Wait for either the read to complete or the context to be done
	select {
	case <-readCtx.Done():
		return 0, nil
	case n := <-res:
		return n, nil
	case err := <-errs:
		return 0, err
	}
}

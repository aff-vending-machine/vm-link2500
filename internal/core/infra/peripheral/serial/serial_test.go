package serial

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/tarm/serial"
)

const Port = "/dev/ttyACM0"

// TestOpenPort tests the OpenPort function.
func TestOpenPort(t *testing.T) {
	_, err := OpenPort(&serial.Config{Name: Port, Baud: 9600})
	assert.NoError(t, err)
}

// TestClose tests the Close method.
func TestClose(t *testing.T) {
	port, _ := OpenPort(&serial.Config{Name: Port, Baud: 9600})
	err := port.Close()
	assert.NoError(t, err)
}

// TestWrite tests the Write method.
func TestWrite(t *testing.T) {
	port, _ := OpenPort(&serial.Config{Name: Port, Baud: 9600})

	// Test writing some data
	err := port.Write(context.Background(), []byte("testdata"))
	assert.NoError(t, err)

	// Test writing with a cancelled context
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	err = port.Write(ctx, []byte("testdata"))
	assert.Error(t, err)

	// Test writing after closing
	port.Close()
	err = port.Write(context.Background(), []byte("testdata"))
	assert.Error(t, err)
}

// TestRead tests the Read method.
func TestRead(t *testing.T) {
	port, _ := OpenPort(&serial.Config{Name: Port, Baud: 9600})

	// Test reading some data
	buf := make([]byte, 10)
	_, err := port.Read(context.Background(), buf)
	assert.NoError(t, err)

	// Test reading with a cancelled context
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	_, err = port.Read(ctx, buf)
	assert.Error(t, err)

	// Test reading after closing
	port.Close()
	_, err = port.Read(context.Background(), buf)
	assert.Error(t, err)
}

// TestConcurrency tests concurrent access to the port.
func TestConcurrency(t *testing.T) {
	port, _ := OpenPort(&serial.Config{Name: Port, Baud: 9600})

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			default:
				_ = port.Write(ctx, []byte("testdata"))
			}
		}
	}()

	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			default:
				buf := make([]byte, 10)
				_, _ = port.Read(ctx, buf)
			}
		}
	}()

	<-ctx.Done()
	assert.NoError(t, port.Close())
}

package link2500

import (
	"time"

	"vm-link2500/configs"

	"github.com/tarm/serial"
)

type serialImpl struct {
	config *serial.Config
}

func New(conf configs.Link2500Config) *serialImpl {
	config := &serial.Config{
		Name:        conf.Port,
		Baud:        9600,
		ReadTimeout: time.Duration(conf.TimeoutInSec) * time.Second,
		Size:        8,
	}

	return &serialImpl{
		config: config,
	}
}

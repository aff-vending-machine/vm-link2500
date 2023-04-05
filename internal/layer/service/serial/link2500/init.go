package link2500

import (
	"time"

	"github.com/aff-vending-machine/vm-link2500/config"
	"github.com/tarm/serial"
)

type serialImpl struct {
	config *serial.Config
}

func New(conf config.Link2500Config) *serialImpl {
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

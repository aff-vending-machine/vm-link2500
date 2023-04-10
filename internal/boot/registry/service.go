package registry

import (
	"github.com/aff-vending-machine/vm-link2500/internal/layer/service/serial"
	"github.com/aff-vending-machine/vm-link2500/internal/layer/service/serial/link2500"
)

// Interface Adapter layers (driven)
type Service struct {
	Serial SerialService
}

type SerialService struct {
	Link2500 serial.Link2500
}

func NewService(module Module) Service {
	return Service{
		SerialService{
			link2500.New(module.Config.Link2500),
		},
	}
}

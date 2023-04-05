package registry

import (
	"github.com/aff-vending-machine/vm-link2500/config"
	"github.com/aff-vending-machine/vm-link2500/internal/layer/service/serial"
	"github.com/aff-vending-machine/vm-link2500/internal/layer/service/serial/link2500"
)

// Interface Adapter layers (driven)
type AppDriven struct {
	Serial SerialDriven
}

type SerialDriven struct {
	Link2500 serial.Link2500
}

func NewAppDriven(cfg config.BootConfig) AppDriven {
	return AppDriven{
		SerialDriven{
			link2500.New(cfg.Link2500),
		},
	}
}

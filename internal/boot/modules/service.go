package modules

import (
	"vm-link2500/internal/layer/service/serial"
)

// Interface Adapter layers (driven)
type Service struct {
	Serial SerialService
}

type SerialService struct {
	Link2500 serial.Link2500
}

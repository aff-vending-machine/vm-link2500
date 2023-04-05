package link2500

import (
	"github.com/aff-vending-machine/vm-link2500/internal/layer/service/serial"
)

type usecaseImpl struct {
	link2500 serial.Link2500
}

func New(s serial.Link2500) *usecaseImpl {
	return &usecaseImpl{s}
}

package registry

import (
	"github.com/aff-vending-machine/vm-link2500/internal/layer/usecase"
	"github.com/aff-vending-machine/vm-link2500/internal/layer/usecase/link2500"
	"github.com/aff-vending-machine/vm-link2500/internal/layer/wrapper/link2500_wrapper"
)

// Usecase layers
type Usecase struct {
	Link2500 usecase.Link2500
}

func NewUsecase(adapter Service) Usecase {
	return Usecase{
		link2500_wrapper.New(link2500.New(
			adapter.Serial.Link2500,
		)),
	}
}

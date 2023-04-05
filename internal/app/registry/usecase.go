package registry

import (
	"github.com/aff-vending-machine/vm-link2500/internal/layer/usecase"
	"github.com/aff-vending-machine/vm-link2500/internal/layer/usecase/link2500"
	"github.com/aff-vending-machine/vm-link2500/internal/layer/usecase/link2500_wrapper"
)

// Usecase layers
type AppUsecase struct {
	Link2500 usecase.Link2500
}

func NewAppUsecase(adapter AppDriven) AppUsecase {
	return AppUsecase{
		link2500_wrapper.New(
			link2500.New(
				adapter.Serial.Link2500,
			),
		),
	}
}

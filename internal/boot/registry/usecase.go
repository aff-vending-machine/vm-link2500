package registry

import (
	"vm-link2500/internal/layer/usecase"
	"vm-link2500/internal/layer/usecase/link2500"
)

// Usecase layers
type Usecase struct {
	Link2500 usecase.Link2500
}

func NewUsecase(adapter Service) Usecase {
	return Usecase{
		link2500.New(
			adapter.Serial.Link2500,
		),
	}
}

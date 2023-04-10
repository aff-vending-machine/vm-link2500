package link2500_wrapper

import (
	"github.com/aff-vending-machine/vm-link2500/internal/layer/usecase"
)

type wrapperImpl struct {
	usecase usecase.Link2500
}

func New(uc usecase.Link2500) *wrapperImpl {
	return &wrapperImpl{uc}
}

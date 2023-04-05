package registry

import (
	"github.com/aff-vending-machine/vm-link2500/internal/layer/transport/http"
	"github.com/aff-vending-machine/vm-link2500/internal/layer/transport/http/link2500"
)

// Interface Adapter layers (driver)
type AppDriver struct {
	HTTP AppHTTPDriver
}

type AppHTTPDriver struct {
	Link2500 http.Link2500
}

func NewAppDriver(uc AppUsecase) AppDriver {
	return AppDriver{
		AppHTTPDriver{
			link2500.New(uc.Link2500),
		},
	}
}

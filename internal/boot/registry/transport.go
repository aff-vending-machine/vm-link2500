package registry

import (
	"github.com/aff-vending-machine/vm-link2500/internal/layer/transport/http"
	"github.com/aff-vending-machine/vm-link2500/internal/layer/transport/http/link2500"
)

// Interface Adapter layers (driver)
type Transport struct {
	HTTP HTTPTransport
}

type HTTPTransport struct {
	Link2500 http.Link2500
}

func NewTransport(uc Usecase) Transport {
	return Transport{
		HTTPTransport{
			link2500.New(uc.Link2500),
		},
	}
}

package registry

import (
	"vm-link2500/internal/layer/transport/http"
	"vm-link2500/internal/layer/transport/http/link2500"
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

package modules

import (
	"vm-link2500/internal/layer/transport/http"
)

// Interface Adapter layers (driver)
type Transport struct {
	HTTP HTTPTransport
}

type HTTPTransport struct {
	Link2500 http.Link2500
}

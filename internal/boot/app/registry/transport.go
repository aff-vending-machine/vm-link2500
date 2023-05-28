package registry

import (
	"vm-link2500/internal/boot/modules"
	"vm-link2500/internal/layer/transport/http/link2500"
)

func NewTransport(uc modules.Usecase) modules.Transport {
	return modules.Transport{
		HTTP: modules.HTTPTransport{
			Link2500: link2500.New(uc.Link2500),
		},
	}
}

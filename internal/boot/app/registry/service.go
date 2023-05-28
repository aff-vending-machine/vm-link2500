package registry

import (
	"vm-link2500/internal/boot/modules"
	"vm-link2500/internal/layer/service/serial/link2500"
)

func NewService(infra modules.Infrastructure) modules.Service {
	return modules.Service{
		Serial: modules.SerialService{
			Link2500: link2500.New(infra.Config.Link2500),
		},
	}
}

package registry

import (
	"github.com/aff-vending-machine/vm-link2500/config"
	"github.com/aff-vending-machine/vm-link2500/internal/core/infra/fiber"
)

// Interface Infrastructure
type AppInfrastructure struct {
	Fiber *fiber.Wrapper
}

// Setup infrastructure
func NewAppInfrastructure(cfg config.BootConfig) AppInfrastructure {
	return AppInfrastructure{
		fiber.New(cfg.Fiber),
	}
}

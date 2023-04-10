package registry

import (
	"github.com/aff-vending-machine/vm-link2500/config"
	"github.com/aff-vending-machine/vm-link2500/internal/core/module/fiber"
)

// Infrastructure
type Module struct {
	Config config.BootConfig
	Fiber  *fiber.Wrapper
}

func NewModule(cfg config.BootConfig) Module {
	return Module{
		Config: cfg,
		Fiber:  fiber.New(cfg.Fiber),
	}
}

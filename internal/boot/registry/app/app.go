package app

import (
	"github.com/aff-vending-machine/vm-link2500/config"
	"github.com/aff-vending-machine/vm-link2500/internal/boot/registry"
	"github.com/aff-vending-machine/vm-link2500/internal/boot/router/fiber"
	"github.com/rs/zerolog/log"
)

func Run(cfg config.BootConfig) {
	log.Debug().Msg("init application")

	var (
		module    = registry.NewModule(cfg)
		service   = registry.NewService(module)
		usecase   = registry.NewUsecase(service)
		transport = registry.NewTransport(usecase)
	)

	fiber.New(module.Fiber).Serve(transport.HTTP)

	log.Debug().Msg("start application")
}

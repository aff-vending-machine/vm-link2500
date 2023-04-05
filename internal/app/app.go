package app

import (
	"github.com/aff-vending-machine/vm-link2500/config"
	"github.com/aff-vending-machine/vm-link2500/internal/app/registry"
	"github.com/aff-vending-machine/vm-link2500/internal/app/router/fiber"
	"github.com/rs/zerolog/log"
)

func Run(cfg config.BootConfig) {
	var (
		infra     = registry.NewAppInfrastructure(cfg)
		service   = registry.NewAppDriven(cfg)
		usecase   = registry.NewAppUsecase(service)
		transport = registry.NewAppDriver(usecase)
	)

	fiber.New(infra.Fiber).Serve(transport.HTTP)

	log.Debug().Msg("start application")
}

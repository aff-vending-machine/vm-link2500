package app

import (
	"vm-link2500/config"
	"vm-link2500/internal/boot/registry"
	"vm-link2500/internal/boot/router/gin"

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

	gin.New(module.Gin).Serve(transport.HTTP)

	log.Debug().Msg("start application")
}

package registry

import (
	"vm-link2500/config"
	"vm-link2500/internal/core/module/gin"
)

// Infrastructure
type Module struct {
	Config config.BootConfig
	Gin    *gin.App
}

func NewModule(cfg config.BootConfig) Module {
	return Module{
		Config: cfg,
		Gin:    gin.New(cfg.Gin),
	}
}

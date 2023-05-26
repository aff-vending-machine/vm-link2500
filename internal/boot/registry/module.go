package registry

import (
	"vm-link2500/config"
	"vm-link2500/internal/core/infra/network/gin"
)

// Infrastructure
type Module struct {
	Config config.Config
	Gin    *gin.App
}

func NewModule(cfg config.Config) Module {
	return Module{
		Config: cfg,
		Gin:    gin.New(cfg.Gin),
	}
}

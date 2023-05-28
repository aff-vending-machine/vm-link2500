package registry

import (
	"vm-link2500/config"
	"vm-link2500/internal/boot/modules"
	"vm-link2500/internal/core/infra/network/gin"
)

// Infrastructure
func NewInfrastructure(cfg config.Config) modules.Infrastructure {
	return modules.Infrastructure{
		Config: cfg,
		Gin:    gin.New(cfg.Gin),
	}
}

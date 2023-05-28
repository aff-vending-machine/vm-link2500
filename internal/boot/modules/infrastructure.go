package modules

import (
	"vm-link2500/config"
	"vm-link2500/internal/core/infra/network/gin"
)

// Infrastructure
type Infrastructure struct {
	Config config.Config
	Gin    *gin.App
}

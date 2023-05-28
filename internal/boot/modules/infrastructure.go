package modules

import (
	"vm-link2500/configs"
	"vm-link2500/internal/core/infra/network/gin"
)

// Infrastructure
type Infrastructure struct {
	Config configs.Config
	Gin    *gin.App
}

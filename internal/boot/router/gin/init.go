package gin

import (
	"vm-link2500/internal/core/module/gin"
)

type server struct {
	*gin.App
}

func New(client *gin.App) *server {
	return &server{
		client,
	}
}

package gin

import (
	"fmt"
	"io"
	"os"

	"vm-link2500/configs"
	"vm-link2500/pkg/log"

	"github.com/gin-gonic/gin"
)

type App struct {
	*gin.Engine
	Address string
}

func New(cfg configs.GinConfig) *App {
	// Create a new Zerolog logger with a human-readable console formatter
	logger := log.ColorConsole()

	gin.SetMode(gin.ReleaseMode)
	gin.DefaultWriter = io.MultiWriter(logger, os.Stdout)
	gin.DefaultErrorWriter = io.MultiWriter(logger, os.Stderr)

	app := gin.Default()
	return &App{app, fmt.Sprintf(":%d", cfg.Port)}
}

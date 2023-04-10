package fiber

import (
	"fmt"

	"github.com/aff-vending-machine/vm-link2500/config"
	"github.com/aff-vending-machine/vm-link2500/internal/core/module/fiber/middleware"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/requestid"
)

type Wrapper struct {
	App     *fiber.App
	Address string
}

func New(cfg config.FiberConfig) *Wrapper {
	app := fiber.New(fiber.Config{
		DisableStartupMessage: true,
		Prefork:               cfg.Prefork,
		CaseSensitive:         cfg.CaseSensitive,
		StrictRouting:         cfg.StrictRouting,
		ServerHeader:          cfg.ServerHeader,
		AppName:               cfg.AppName,
	})

	app.Use(requestid.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "*",
		AllowMethods:     "POST",
		AllowHeaders:     "accept,content-type",
		AllowCredentials: true,
		MaxAge:           1728000,
	}))
	app.Use(middleware.NewLogger())
	// app.Use(csrf.New())

	return &Wrapper{
		App:     app,
		Address: fmt.Sprintf(":%d", cfg.Port),
	}
}

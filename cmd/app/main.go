package main

import (
	"vm-link2500/configs"
	"vm-link2500/internal/boot/app"
	"vm-link2500/pkg/boot"
	"vm-link2500/pkg/log"
)

func init() {
	log.New()
}

func main() {
	// Create boot with configuration
	conf := configs.Init("env/app")
	boot.Init(conf)
	defer boot.Serve()

	initLog(conf)
	conf.Preview()

	// Run main application
	app.Run(conf)
}

func initLog(conf configs.Config) {
	log.SetOutput(log.ColorConsole())
	log.SetLogLevel(conf.App.LogLevel)
}

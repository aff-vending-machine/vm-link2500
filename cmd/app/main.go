package main

import (
	"vm-link2500/config"
	"vm-link2500/internal/boot/registry/app"
	"vm-link2500/pkg/boot"
	"vm-link2500/pkg/log"
)

func init() {
	log.New()
}

func main() {
	// Create boot with configuration
	conf := config.Init("env/app")
	boot.Init(conf)
	defer boot.Serve()

	initLog(conf)

	// Run main application
	app.Run(conf)
}

func initLog(conf config.BootConfig) {
	log.SetOutput(log.ColorConsole())
	log.SetLogLevel(conf.App.LogLevel)
}

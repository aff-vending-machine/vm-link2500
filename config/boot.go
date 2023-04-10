package config

type BootConfig struct {
	App      AppConfig      `mapstructure:"APP"`
	Fiber    FiberConfig    `mapstructure:"FIBER"`
	Link2500 Link2500Config `mapstructure:"LINK2500"`
}

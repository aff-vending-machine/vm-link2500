package config

type BootConfig struct {
	App      AppConfig      `mapstructure:"APP"`
	Gin      GinConfig      `mapstructure:"GIN"`
	Link2500 Link2500Config `mapstructure:"LINK2500"`
}

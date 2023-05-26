package config

type GinConfig struct {
	Port int `default:"8082" mapstructure:"PORT"`
}

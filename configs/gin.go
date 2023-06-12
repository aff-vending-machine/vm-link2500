package configs

type GinConfig struct {
	Port int `default:"3002" mapstructure:"PORT"`
}

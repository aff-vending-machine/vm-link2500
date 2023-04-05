package config

type AppConfig struct {
	ENV      string `default:"production" mapstructure:"ENV"`
	LogLevel int    `default:"0" mapstructure:"LOG_LEVEL"`
}

package config

type FiberConfig struct {
	Prefork       bool   `default:"false" mapstructure:"PREFORK"`
	CaseSensitive bool   `default:"false" mapstructure:"CASE_SENSITIVE"`
	StrictRouting bool   `default:"false" mapstructure:"STRICT_ROUTING"`
	ServerHeader  string `default:"" mapstructure:"SERVER_HEADER"`
	AppName       string `default:"link2500" mapstructure:"APP_NAME"`
	Port          int    `default:"8082" mapstructure:"PORT"`
}

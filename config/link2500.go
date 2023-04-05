package config

type Link2500Config struct {
	Port         string `default:"/dev/ttyACM0" mapstructure:"PORT"`
	TimeoutInSec int    `default:"120" mapstructure:"TIMEOUT_IN_SEC"`
}

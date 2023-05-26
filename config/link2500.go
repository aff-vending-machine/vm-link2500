package config

type Link2500Config struct {
	Port         string `default:"/dev/link2500" mapstructure:"PORT"`
	TimeoutInSec int    `default:"10" mapstructure:"TIMEOUT_IN_SEC"`
}

package config

type Link2500Config struct {
	Port         string `default:"/dev/ACM0" mapstructure:"PORT"`
	TimeoutInSec int    `default:"15" mapstructure:"TIMEOUT_IN_SEC"`
}

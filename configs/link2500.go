package configs

type Link2500Config struct {
	Port         string `default:"/dev/ttyACM0" mapstructure:"PORT"`
	TimeoutInSec int    `default:"15" mapstructure:"TIMEOUT_IN_SEC"`
}

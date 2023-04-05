package config

import (
	"os"
	"reflect"
	"strings"

	"github.com/creasty/defaults"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

// Init creates a new config service.
func Init(fallback string) BootConfig {
	filename := fallback
	if value, ok := os.LookupEnv("APP_ENV"); ok {
		filename = value
	}

	var out BootConfig
	if err := defaults.Set(&out); err != nil {
		log.Panic().Err(err).Msgf("set default config failed")
	}

	v := viper.New()
	v.SetConfigName(filename)
	v.AddConfigPath(".")
	v.SetConfigType("yaml")
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_", "-", "_"))

	// read config file, ignore error when not found.
	if err := v.ReadInConfig(); err != nil {
		log.Warn().Err(err).Msgf("read config file failed")
	}

	// set environment variables as overrides.
	bindEnvs(v, BootConfig{})
	v.AutomaticEnv()

	// set config form file or env
	if err := v.Unmarshal(&out); err != nil {
		log.Warn().Err(err).Msgf("unmarshal config failed")
	}

	preview(out)

	return out
}

func bindEnvs(vp *viper.Viper, iface interface{}, parts ...string) {
	ifv := reflect.ValueOf(iface)
	ift := reflect.TypeOf(iface)
	for i := 0; i < ift.NumField(); i++ {
		v := ifv.Field(i)
		t := ift.Field(i)
		tv, ok := t.Tag.Lookup("mapstructure")
		if !ok {
			continue
		}
		switch v.Kind() {
		case reflect.Struct:
			bindEnvs(vp, v.Interface(), append(parts, tv)...)
		default:
			vp.BindEnv(strings.Join(append(parts, tv), "."))
		}
	}
}

func preview(conf BootConfig) {
	log.Debug().Interface("App", conf.App).Msg("configuration")
	log.Debug().Interface("Fiber", conf.Fiber).Msg("configuration")
	log.Debug().Interface("SmartEDC", conf.Link2500).Msg("configuration")
}

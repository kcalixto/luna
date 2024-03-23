package config

import "github.com/spf13/viper"

type Config struct {
	Environment string
}

func New() *Config {
	env := viper.GetString("env")

	return &Config{
		Environment: env,
	}
}

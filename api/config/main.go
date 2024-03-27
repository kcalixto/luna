package config

import "os"

type Config struct {
	Environment string
	Database    Database
}

type Database struct {
	SingleTableName string
}

func New() *Config {
	env := os.Getenv("ENV")

	return &Config{
		Environment: env,
		Database: Database{
			SingleTableName: os.Getenv("SINGLE_TABLE_NAME"),
		},
	}
}

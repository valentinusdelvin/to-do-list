package config

import (
	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
)

type Config struct {
	DbUser     string `env:"DB_USER"`
	DbPassword string `env:"DB_PASSWORD"`
	DbName     string `env:"DB_NAME"`
}

var cfg Config

func LoadConfig() Config {
	return cfg
}

func NewConfig() Config {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}

	err = env.Parse(&cfg)
	if err != nil {
		panic(err)
	}
	return cfg
}

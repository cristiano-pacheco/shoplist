package config

import (
	"log"

	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
)

type Config struct {
	HTTPPort    uint   `env:"HTTP_PORT"`
	Environment string `env:"ENVIROMENT"`
	Cors        string `env:"CORS"`
	DB          DB     `envPrefix:"DB_"`
}

const EnvProduction = "production"
const EnvDevelopment = "development"
const EnvStaging = "staging"

var _global Config

func Init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	env.Parse(&_global)
}

func GetConfig() Config {
	return _global
}

func (c *Config) IsProduction() bool {
	return c.Environment == "production"
}

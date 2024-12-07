package config

import (
	"log"

	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
)

type Config struct {
	Environment string    `env:"ENVIROMENT"`
	Cors        string    `env:"CORS"`
	JWT         JWT       `envPrefix:"JWT_"`
	DB          DB        `envPrefix:"DB_"`
	MAIL        MAIL      `envPrefix:"MAIL_"`
	HTTPPort    uint      `env:"HTTP_PORT"`
	Telemetry   Telemetry `envPrefix:"TELEMETRY_"`
	App         App       `envPrefix:"APP_"`
	Log         Log       `envPrefix:"LOG_"`
	Faktory     Faktory   `envPrefix:"FAKTORY_"`
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

	err = env.Parse(&_global)
	if err != nil {
		log.Fatalf("%+v\n", err)
	}
}

func GetConfig() Config {
	return _global
}

func (c *Config) IsProduction() bool {
	return c.Environment == "production"
}

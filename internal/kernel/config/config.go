package config

import (
	"log"
	"strings"

	"github.com/spf13/viper"
)

type Config struct {
	Environment string    `mapstructure:"ENVIRONMENT"`
	HTTPPort    uint      `mapstructure:"HTTP_PORT"`
	CORS        CORS      `mapstructure:",squash"`
	JWT         JWT       `mapstructure:",squash"`
	DB          DB        `mapstructure:",squash"`
	MAIL        MAIL      `mapstructure:",squash"`
	Telemetry   Telemetry `mapstructure:",squash"`
	App         App       `mapstructure:",squash"`
	Log         Log       `mapstructure:",squash"`
	RabbitMQ    RabbitMQ  `mapstructure:",squash"`
}

const EnvProduction = "production"
const EnvDevelopment = "development"
const EnvStaging = "staging"

var _global Config

func Init() {
	v := viper.New()

	// Allow environment variables to override config file settings
	v.AutomaticEnv()
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	// Configure to read from .env file
	v.SetConfigName(".env")
	v.SetConfigType("env")
	v.AddConfigPath(".")

	// Read the config file (must exist)
	if err := v.ReadInConfig(); err != nil {
		panic(err)
	}

	log.Printf("Using config file: %s", v.ConfigFileUsed())

	// Unmarshal the config into our struct
	if err := v.Unmarshal(&_global); err != nil {
		log.Fatalf("Failed to unmarshal config: %+v", err)
	}
}

func GetConfig() Config {
	return _global
}

func (c *Config) IsProduction() bool {
	return c.Environment == "production"
}

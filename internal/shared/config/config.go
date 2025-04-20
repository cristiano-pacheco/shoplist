package config

import (
	"log"
	"strings"

	"github.com/spf13/viper"
)

type Config struct {
	Environment string    `mapstructure:"ENVIROMENT"`
	Cors        string    `mapstructure:"CORS"`
	JWT         JWT       `mapstructure:",squash"`
	DB          DB        `mapstructure:",squash"`
	MAIL        MAIL      `mapstructure:",squash"`
	HTTPPort    uint      `mapstructure:"HTTP_PORT"`
	Telemetry   Telemetry `mapstructure:",squash"`
	App         App       `mapstructure:",squash"`
	Log         Log       `mapstructure:",squash"`
	RabbitMQ    RabbitMQ  `mapstructure:",squash"`
}

const EnvProduction = "production"
const EnvDevelopment = "development"
const EnvStaging = "staging"

var _global Config

func Init(envFile ...string) {
	v := viper.New()

	// Set defaults (can be added as needed)
	setDefaults(v)

	// Allow environment variables to override config file settings
	v.AutomaticEnv()
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	// If specific config file provided, use it
	if len(envFile) > 0 && envFile[0] != "" {
		v.SetConfigFile(envFile[0])
	} else {
		// Otherwise look for config in standard locations
		v.SetConfigName(".env")      // name of config file (without extension)
		v.SetConfigType("env")       // REQUIRED if the config file does not have the extension in the name
		v.AddConfigPath(".")         // look for config in the working directory
		v.AddConfigPath("./config/") // look for config in ./config/ directory

		// Add support for other config formats
		v.SetConfigName("config") // also look for config.yml, config.json, etc
		v.AddConfigPath(".")
		v.AddConfigPath("./config/")
	}

	// Try to read the config file, but don't fail if it doesn't exist
	if err := v.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Println("No config file found, using environment variables and defaults")
		} else {
			log.Printf("Warning: Error reading config file: %v", err)
		}
	} else {
		log.Printf("Using config file: %s", v.ConfigFileUsed())
	}

	// Unmarshal the config into our struct
	if err := v.Unmarshal(&_global); err != nil {
		log.Fatalf("Failed to unmarshal config: %+v", err)
	}
}

func setDefaults(v *viper.Viper) {
	// DB defaults
	v.SetDefault("DB_HOST", "localhost")
	v.SetDefault("DB_NAME", "shoplist")
	v.SetDefault("DB_USER", "postgres")
	v.SetDefault("DB_PASSWORD", "postgres")
	v.SetDefault("DB_PORT", 5432)
	v.SetDefault("DB_MAX_OPEN_CONNECTIONS", 10)
	v.SetDefault("DB_MAX_IDLE_CONNECTIONS", 10)
	v.SetDefault("DB_SSL_MODE", false)
	v.SetDefault("DB_PREPARE_STMT", false)
	v.SetDefault("DB_ENABLE_LOGS", false)

	// RabbitMQ defaults
	v.SetDefault("RABBITMQ_HOST", "localhost")
	v.SetDefault("RABBITMQ_PORT", "5672")
	v.SetDefault("RABBITMQ_USERNAME", "guest")
	v.SetDefault("RABBITMQ_PASSWORD", "guest")
	v.SetDefault("RABBITMQ_VHOST", "/")

	// Add other defaults as needed
}

func GetConfig() Config {
	return _global
}

func (c *Config) IsProduction() bool {
	return c.Environment == "production"
}

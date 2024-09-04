package config

import (
	"fmt"
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

func GenerateGormDatabaseDSN(cfg Config) string {
	sslMode := "enabled"
	if !cfg.DB.SSLMode {
		sslMode = "disable"
	}

	dsn := fmt.Sprintf(
		"host=%s "+
			"user=%s "+
			"password=%s "+
			"dbname=%s "+
			"port=%d "+
			"sslmode=%s "+
			"TimeZone=UTC",
		cfg.DB.DBHost,
		cfg.DB.DBUser,
		cfg.DB.DBPassword,
		cfg.DB.DBName,
		cfg.DB.DBPort,
		sslMode,
	)

	return dsn
}

func GeneratePostgresDatabaseDSN(cfg Config) string {
	return fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s?sslmode=disable&TimeZone=UTC",
		cfg.DB.DBUser,
		cfg.DB.DBPassword,
		cfg.DB.DBHost,
		cfg.DB.DBPort,
		cfg.DB.DBName,
	)
}

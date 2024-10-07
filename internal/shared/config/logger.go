package config

type Log struct {
	IsEnabled bool   `env:"ENABLED"`
	LogLevel  string `env:"LEVEL"`
}

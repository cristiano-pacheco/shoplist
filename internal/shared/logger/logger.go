package logger

import (
	"github.com/cristiano-pacheco/go-modulith/internal/shared/config"
	"github.com/cristiano-pacheco/go-modulith/pkg/logger"
)

type LoggerI interface {
	logger.LoggerI
}

func New(config config.Config) LoggerI {
	logConfig := logger.LoggerConfig{
		IsEnabled: config.Log.IsEnabled,
		LogLevel:  logger.LogLevel(config.Log.LogLevel),
	}
	return logger.New(logConfig)
}

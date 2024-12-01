package logger

import (
	"github.com/cristiano-pacheco/shoplist/internal/shared/config"
	"github.com/cristiano-pacheco/shoplist/pkg/logger"
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

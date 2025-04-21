package logger

import (
	"github.com/cristiano-pacheco/shoplist/internal/kernel/config"
	"github.com/cristiano-pacheco/shoplist/pkg/logger"
)

type Logger interface {
	logger.Logger
}

func New(config config.Config) Logger {
	logConfig := logger.LoggerConfig{
		IsEnabled: config.Log.IsEnabled,
		LogLevel:  logger.LogLevel(config.Log.LogLevel),
	}
	return logger.New(logConfig)
}

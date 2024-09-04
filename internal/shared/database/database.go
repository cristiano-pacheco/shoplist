package database

import (
	"time"

	"github.com/cristiano-pacheco/go-modulith/internal/shared/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"
)

type DB struct {
	*gorm.DB
}

func New(cfg config.Config) *DB {
	return openConnection(cfg)
}

func openConnection(cfg config.Config) *DB {
	dsn := config.GenerateGormDatabaseDSN(cfg)
	gormConf := gorm.Config{}
	writer := stdoutWriter{}

	if cfg.DB.EnableLogs {
		newLogger := gormLogger.New(writer, gormLogger.Config{
			SlowThreshold: time.Second,
			LogLevel:      gormLogger.Info,
			Colorful:      false,
		})
		gormConf.Logger = newLogger
	}

	pgconfig := postgres.Config{DSN: dsn}
	db, err := gorm.Open(postgres.New(pgconfig), &gormConf)
	if err != nil {
		panic(err)
	}

	return &DB{db}
}

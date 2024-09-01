package database

import (
	"fmt"
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
	dsn := generateDSN(cfg)
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

func generateDSN(cfg config.Config) string {
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

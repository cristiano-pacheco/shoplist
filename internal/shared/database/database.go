package database

import (
	"github.com/cristiano-pacheco/go-modulith/internal/shared/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DB struct {
	*gorm.DB
}

func New(cfg config.Config) *DB {
	return openConnection(cfg)
}

func NewFromGorm(db *gorm.DB) *DB {
	return &DB{db}
}

func openConnection(cfg config.Config) *DB {
	dsn := config.GenerateGormDatabaseDSN(cfg)
	gormConf := gorm.Config{}

	if cfg.DB.EnableLogs {
		gormConf.Logger = logger.Default.LogMode(logger.Info)
	}

	pgconfig := postgres.Config{DSN: dsn}
	db, err := gorm.Open(postgres.New(pgconfig), &gormConf)
	if err != nil {
		panic(err)
	}

	return &DB{db}
}

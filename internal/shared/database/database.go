package database

import (
	"log"
	"os"
	"time"

	"github.com/cristiano-pacheco/go-modulith/internal/shared/config"
	pkg_logger "github.com/cristiano-pacheco/go-modulith/pkg/logger"
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
	logLevel := pkg_logger.ParseLogLevel(pkg_logger.LogLevel(cfg.Log.LogLevel))
	gormLogLevel := logger.LogLevel(logLevel)
	if !cfg.Log.IsEnabled {
		gormLogLevel = logger.Silent
	}

	loggerConfig := logger.Config{
		SlowThreshold:             200 * time.Millisecond, // Slow SQL threshold
		LogLevel:                  gormLogLevel,           // Log level
		IgnoreRecordNotFoundError: true,                   // Ignore ErrRecordNotFound error for logger
		ParameterizedQueries:      true,                   // Don't include params in the SQL log
		Colorful:                  false,                  // Disable color
	}

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		loggerConfig,
	)

	if cfg.DB.EnableLogs {
		gormConf.Logger = newLogger
	}

	pgconfig := postgres.Config{DSN: dsn}
	db, err := gorm.Open(postgres.New(pgconfig), &gormConf)
	if err != nil {
		panic(err)
	}

	return &DB{db}
}

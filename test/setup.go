package test

import (
	"context"
	"log"
	"path/filepath"
	"runtime"

	"github.com/cristiano-pacheco/shoplist/internal/modules/identity"
	"github.com/cristiano-pacheco/shoplist/internal/modules/list"
	"github.com/cristiano-pacheco/shoplist/internal/shared"
	"github.com/cristiano-pacheco/shoplist/internal/shared/config"
	"github.com/cristiano-pacheco/shoplist/internal/shared/database"
	"github.com/cristiano-pacheco/shoplist/internal/shared/otel"
	database_pkg "github.com/cristiano-pacheco/shoplist/pkg/database"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"go.uber.org/fx"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type IntegrationTest struct {
	DB      *gorm.DB
	Config  config.Config
	BaseURL string
}

func Setup() *IntegrationTest {
	_, currentFile, _, _ := runtime.Caller(0)
	projectRoot := filepath.Join(filepath.Dir(currentFile), "../")
	envFile := filepath.Join(projectRoot, ".env")
	config.Init(envFile)

	cfg := config.GetConfig()
	otel.Init(cfg)

	app := fx.New(
		shared.Module,
		identity.Module,
		list.Module,
	)

	app.Start(context.Background())

	db := database.New(cfg)
	return &IntegrationTest{
		DB:      db.DB,
		Config:  cfg,
		BaseURL: cfg.App.BaseURL,
	}
}

func (t *IntegrationTest) Close() {
	sqlDB, err := t.DB.DB()
	if err != nil {
		log.Fatalf("Error getting underlying *sql.DB: %v", err)
	}
	sqlDB.Close()
}

func (t *IntegrationTest) SetupDB() {
	dbName := t.Config.DB.Name

	// Create a new connection to the default 'postgres' database
	defaultDBConfig := database_pkg.DatabaseConfig{
		Host:               t.Config.DB.Host,
		User:               t.Config.DB.User,
		Password:           t.Config.DB.Password,
		Name:               "postgres", // Connect to the default 'postgres' database
		Port:               t.Config.DB.Port,
		MaxOpenConnections: t.Config.DB.MaxOpenConnections,
		MaxIdleConnections: t.Config.DB.MaxIdleConnections,
		SSLMode:            t.Config.DB.SSLMode,
		PrepareSTMT:        t.Config.DB.PrepareSTMT,
		EnableLogs:         t.Config.DB.EnableLogs,
	}

	defaultDSN := database_pkg.GeneratePostgresDatabaseDSN(defaultDBConfig)
	defaultDB, err := gorm.Open(postgres.Open(defaultDSN), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error connecting to default database: %v", err)
	}

	// Use the new connection to drop and create the database
	sqlDB, err := defaultDB.DB()
	if err != nil {
		log.Fatalf("Error getting underlying *sql.DB: %v", err)
	}
	defer sqlDB.Close()

	// Terminate all connections to the database before dropping it
	_, err = sqlDB.Exec(`
		SELECT pg_terminate_backend(pg_stat_activity.pid)
		FROM pg_stat_activity
		WHERE pg_stat_activity.datname = '` + dbName + `'
		AND pid <> pg_backend_pid()
	`)
	if err != nil {
		log.Printf("Warning: Error terminating existing connections: %v", err)
	}

	// Drop the database if it exists
	_, err = sqlDB.Exec("DROP DATABASE IF EXISTS " + dbName)
	if err != nil {
		log.Fatalf("Error dropping database: %v", err)
	}

	// Create the new database
	_, err = sqlDB.Exec("CREATE DATABASE " + dbName)
	if err != nil {
		log.Fatalf("Error creating database: %v", err)
	}

	dbConfig := database_pkg.DatabaseConfig{
		Host:               t.Config.DB.Host,
		User:               t.Config.DB.User,
		Password:           t.Config.DB.Password,
		Name:               t.Config.DB.Name,
		Port:               t.Config.DB.Port,
		MaxOpenConnections: t.Config.DB.MaxOpenConnections,
		MaxIdleConnections: t.Config.DB.MaxIdleConnections,
		SSLMode:            t.Config.DB.SSLMode,
		PrepareSTMT:        t.Config.DB.PrepareSTMT,
		EnableLogs:         t.Config.DB.EnableLogs,
	}

	dsn := database_pkg.GeneratePostgresDatabaseDSN(dbConfig)

	_, currentFile, _, _ := runtime.Caller(0)
	projectRoot := filepath.Join(filepath.Dir(currentFile), "../")
	migrationsPath := filepath.Join(projectRoot, "migrations")
	m, err := migrate.New("file://"+migrationsPath, dsn)
	if err != nil {
		log.Fatal(err)
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatal(err)
	}
}

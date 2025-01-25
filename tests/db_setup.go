package tests

import (
	"fmt"
	"log"

	"github.com/cristiano-pacheco/shoplist/internal/shared/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func setupTestDB(cfg config.Config) error {
	// Connect to default postgres database to create/drop test database
	dsn := fmt.Sprintf("host=%s user=%s password=%s port=%d dbname=postgres sslmode=disable",
		cfg.DB.Host, cfg.DB.User, cfg.DB.Password, cfg.DB.Port)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("failed to connect to postgres: %v", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		return fmt.Errorf("failed to get underlying *sql.DB: %v", err)
	}
	defer sqlDB.Close()

	// Drop test database if it exists
	db.Exec(fmt.Sprintf("DROP DATABASE IF EXISTS %s", cfg.DB.Name))

	// Create test database
	err = db.Exec(fmt.Sprintf("CREATE DATABASE %s", cfg.DB.Name)).Error
	if err != nil {
		return fmt.Errorf("failed to create test database: %v", err)
	}

	log.Printf("Created test database: %s", cfg.DB.Name)
	return nil
}

func teardownTestDB(cfg config.Config) error {
	// Connect to default postgres database to drop test database
	dsn := fmt.Sprintf("host=%s user=%s password=%s port=%d dbname=postgres sslmode=disable",
		cfg.DB.Host, cfg.DB.User, cfg.DB.Password, cfg.DB.Port)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("failed to connect to postgres: %v", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		return fmt.Errorf("failed to get underlying *sql.DB: %v", err)
	}
	defer sqlDB.Close()

	// Drop test database
	err = db.Exec(fmt.Sprintf("DROP DATABASE IF EXISTS %s", cfg.DB.Name)).Error
	if err != nil {
		return fmt.Errorf("failed to drop test database: %v", err)
	}

	log.Printf("Dropped test database: %s", cfg.DB.Name)
	return nil
}

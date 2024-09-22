package service

import (
	"fmt"
	"log"
	"os"

	"github.com/cristiano-pacheco/go-modulith/internal/shared/config"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

type ExecuteDBMigrationsService struct {
	cfg config.Config
}

func ExecuteMigrations(cfg config.Config) {
	service := &ExecuteDBMigrationsService{cfg}
	service.execute()
}

func (s *ExecuteDBMigrationsService) execute() {
	dsn := config.GeneratePostgresDatabaseDSN(s.cfg)

	m, err := migrate.New("file://migrations", dsn)
	if err != nil {
		log.Fatal(err)
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatal(err)
	}

	fmt.Println("Migrations executed successfully")
	os.Exit(0)
}

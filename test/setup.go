package test

import (
	"context"
	"log"

	"github.com/cristiano-pacheco/shoplist/internal/modules/identity"
	"github.com/cristiano-pacheco/shoplist/internal/modules/list"
	"github.com/cristiano-pacheco/shoplist/internal/shared"
	"github.com/cristiano-pacheco/shoplist/internal/shared/config"
	"github.com/cristiano-pacheco/shoplist/internal/shared/database"
	"github.com/cristiano-pacheco/shoplist/internal/shared/otel"
	"go.uber.org/fx"
	"gorm.io/gorm"
)

type IntegrationTest struct {
	DB     *gorm.DB
	Config config.Config
}

func Setup() *IntegrationTest {
	envFile := "../../.env"
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
		DB:     db.DB,
		Config: cfg,
	}
}

func (t *IntegrationTest) Close() {
	sqlDB, err := t.DB.DB()
	if err != nil {
		log.Fatalf("Error getting underlying *sql.DB: %v", err)
	}
	sqlDB.Close()
}

package tests

import (
	"testing"

	"github.com/cristiano-pacheco/shoplist/internal/shared/database"
	"github.com/cristiano-pacheco/shoplist/internal/shared/http/httpserver"
	"github.com/stretchr/testify/assert"
	"go.uber.org/fx"
)

func TestExample(t *testing.T) {
	var (
		server *httpserver.Server
		db     *database.ShoplistDB
	)

	app := NewTestApp(t,
		fx.Populate(&server, &db),
	)

	// Start the application
	app.Start()
	defer app.Stop()

	// Test the server
	assert.NotNil(t, server, "Server should not be nil")
	assert.NotNil(t, db, "Database should not be nil")

	// Test database connection
	sqlDB, err := db.DB.DB()
	assert.NoError(t, err, "Should be able to get underlying *sql.DB")
	assert.NoError(t, sqlDB.Ping(), "Should be able to ping database")
}

package tests

import (
	"testing"

	"github.com/cristiano-pacheco/shoplist/internal/shared/http/httpserver"
	"go.uber.org/fx"
)

func TestExample(t *testing.T) {
	var server *httpserver.Server

	app := NewTestApp(t,
		fx.Populate(&server), // Inject the HTTP server into our variable
	)

	// Start the application
	app.Start()
	defer app.Stop()

	// Now you can use the server or any other injected dependencies for your tests
	if server == nil {
		t.Error("Server should not be nil")
	}
}

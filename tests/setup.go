package tests

import (
	"context"
	"log"
	"os"
	"path/filepath"
	"testing"

	"github.com/cristiano-pacheco/shoplist/cmd"
	"github.com/cristiano-pacheco/shoplist/internal/modules/identity"
	"github.com/cristiano-pacheco/shoplist/internal/modules/list"
	"github.com/cristiano-pacheco/shoplist/internal/shared"
	"github.com/cristiano-pacheco/shoplist/internal/shared/config"
	"github.com/cristiano-pacheco/shoplist/internal/shared/otel"
	"go.uber.org/fx"
	"go.uber.org/fx/fxtest"
)

type TestApp struct {
	*fxtest.App
}

var testEnvFile string

func init() {
	// Try to find the project root and .env.test file
	wd, err := os.Getwd()
	if err != nil {
		log.Printf("Error getting working directory: %v", err)
		return
	}

	// Try to find the project root by looking for go.mod
	for {
		if _, err := os.Stat(filepath.Join(wd, "go.mod")); err == nil {
			testEnvFile = filepath.Join(wd, ".env.test")
			break
		}
		parent := filepath.Dir(wd)
		if parent == wd {
			log.Printf("Could not find project root")
			return
		}
		wd = parent
	}
}

func NewTestApp(t *testing.T, opts ...fx.Option) *TestApp {
	config.Init(testEnvFile)
	cfg := config.GetConfig()
	otel.Init(cfg)

	defaultOpts := []fx.Option{
		fx.Provide(func() *testing.T { return t }),
		shared.Module,
		identity.Module,
		list.Module,
	}

	opts = append(defaultOpts, opts...)

	app := fxtest.New(t, opts...)

	return &TestApp{App: app}
}

func (a *TestApp) Start() {
	if err := a.App.Start(context.Background()); err != nil {
		log.Fatalf("Failed to start test app: %v", err)
	}
}

func (a *TestApp) Stop() {
	if err := a.App.Stop(context.Background()); err != nil {
		log.Printf("Error stopping test app: %v", err)
	}

	if err := otel.Trace().Shutdown(context.Background()); err != nil {
		log.Printf("Error shutting down tracer provider: %v", err)
	}
}

// Setup is kept for backward compatibility
func Setup() {
	config.Init()
	cfg := config.GetConfig()
	otel.Init(cfg)

	defer func() {
		if err := otel.Trace().Shutdown(context.Background()); err != nil {
			log.Printf("Error shutting down tracer provider: %v", err)
		}
	}()

	cmd.Execute()
}

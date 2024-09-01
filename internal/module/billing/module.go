package billing

import (
	"github.com/cristiano-pacheco/go-modulith/internal/module/billing/handler"
	"github.com/cristiano-pacheco/go-modulith/internal/module/billing/repository"
	"github.com/cristiano-pacheco/go-modulith/internal/module/billing/router"
	"github.com/cristiano-pacheco/go-modulith/internal/module/billing/usecase"
	"go.uber.org/fx"
)

var Module = fx.Module(
	"billing",
	handler.Module,
	repository.Module,
	router.Module,
	usecase.Module,
)

package identity

import (
	"github.com/cristiano-pacheco/go-modulith/internal/module/identity/handler"
	"github.com/cristiano-pacheco/go-modulith/internal/module/identity/repository"
	"github.com/cristiano-pacheco/go-modulith/internal/module/identity/router"
	"github.com/cristiano-pacheco/go-modulith/internal/module/identity/service"
	"github.com/cristiano-pacheco/go-modulith/internal/module/identity/usecase"
	"go.uber.org/fx"
)

var Module = fx.Module(
	"identity",
	handler.Module,
	repository.Module,
	router.Module,
	usecase.Module,
	service.Module,
)

package identity

import (
	"github.com/cristiano-pacheco/go-modulith/internal/identity/handler"
	"github.com/cristiano-pacheco/go-modulith/internal/identity/repository"
	"github.com/cristiano-pacheco/go-modulith/internal/identity/router"
	"github.com/cristiano-pacheco/go-modulith/internal/identity/service"
	"github.com/cristiano-pacheco/go-modulith/internal/identity/usecase"
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

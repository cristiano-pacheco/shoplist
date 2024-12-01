package identity

import (
	"github.com/cristiano-pacheco/shoplist/internal/modules/identity/handler"
	"github.com/cristiano-pacheco/shoplist/internal/modules/identity/repository"
	"github.com/cristiano-pacheco/shoplist/internal/modules/identity/router"
	"github.com/cristiano-pacheco/shoplist/internal/modules/identity/service"
	"github.com/cristiano-pacheco/shoplist/internal/modules/identity/usecase"
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

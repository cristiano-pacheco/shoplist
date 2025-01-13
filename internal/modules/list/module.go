package list

import (
	"github.com/cristiano-pacheco/shoplist/internal/modules/list/handler"
	"github.com/cristiano-pacheco/shoplist/internal/modules/list/repository"
	"github.com/cristiano-pacheco/shoplist/internal/modules/list/router"
	"github.com/cristiano-pacheco/shoplist/internal/modules/list/usecase"
	"go.uber.org/fx"
)

var Module = fx.Module(
	"list",
	usecase.Module,
	handler.Module,
	router.Module,
	repository.Module,
)

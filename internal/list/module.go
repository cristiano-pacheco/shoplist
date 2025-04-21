package list

import (
	"github.com/cristiano-pacheco/shoplist/internal/list/handler"
	"github.com/cristiano-pacheco/shoplist/internal/list/repository"
	"github.com/cristiano-pacheco/shoplist/internal/list/router"
	"github.com/cristiano-pacheco/shoplist/internal/list/usecase"
	"go.uber.org/fx"
)

var Module = fx.Module(
	"list",
	usecase.Module,
	handler.Module,
	router.Module,
	repository.Module,
)

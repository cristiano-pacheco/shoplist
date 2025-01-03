package list

import (
	"github.com/cristiano-pacheco/shoplist/internal/modules/list/usecase"
	"go.uber.org/fx"
)

var Module = fx.Module(
	"list",
	usecase.Module,
)

package parser

import (
	"github.com/cristiano-pacheco/shoplist/internal/shared/parser/jwtparser"
	"go.uber.org/fx"
)

var Module = fx.Module(
	"shared/parser",
	fx.Provide(jwtparser.New),
)

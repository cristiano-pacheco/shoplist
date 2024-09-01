package shared

import (
	"github.com/cristiano-pacheco/go-modulith/internal/shared/module/config"
	"github.com/cristiano-pacheco/go-modulith/internal/shared/module/database"
	"github.com/cristiano-pacheco/go-modulith/internal/shared/module/validator"
	"go.uber.org/fx"
)

var Module = fx.Module(
	"shared",
	config.Module,
	database.Module,
	validator.Module,
)

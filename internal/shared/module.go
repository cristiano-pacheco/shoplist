package shared

import (
	"github.com/cristiano-pacheco/go-modulith/internal/shared/config"
	"github.com/cristiano-pacheco/go-modulith/internal/shared/database"
	"github.com/cristiano-pacheco/go-modulith/internal/shared/httpserver"
	"github.com/cristiano-pacheco/go-modulith/internal/shared/mapper/errormapper"
	"github.com/cristiano-pacheco/go-modulith/internal/shared/translator"
	"github.com/cristiano-pacheco/go-modulith/internal/shared/validator"
	"go.uber.org/fx"
)

var Module = fx.Module(
	"shared",
	config.Module,
	database.Module,
	validator.Module,
	translator.Module,
	httpserver.Module,
	errormapper.Module,
)

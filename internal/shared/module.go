package shared

import (
	"github.com/cristiano-pacheco/go-modulith/internal/shared/config"
	"github.com/cristiano-pacheco/go-modulith/internal/shared/database"
	"github.com/cristiano-pacheco/go-modulith/internal/shared/httpserver"
	"github.com/cristiano-pacheco/go-modulith/internal/shared/logger"
	"github.com/cristiano-pacheco/go-modulith/internal/shared/mailer"
	"github.com/cristiano-pacheco/go-modulith/internal/shared/mapper"
	"github.com/cristiano-pacheco/go-modulith/internal/shared/middleware"
	"github.com/cristiano-pacheco/go-modulith/internal/shared/parser"
	"github.com/cristiano-pacheco/go-modulith/internal/shared/registry"
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
	mapper.Module,
	logger.Module,
	middleware.Module,
	registry.Module,
	parser.Module,
	mailer.Module,
)

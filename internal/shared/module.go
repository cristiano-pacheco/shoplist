package shared

import (
	"github.com/cristiano-pacheco/shoplist/internal/shared/config"
	"github.com/cristiano-pacheco/shoplist/internal/shared/database"
	"github.com/cristiano-pacheco/shoplist/internal/shared/errs"
	"github.com/cristiano-pacheco/shoplist/internal/shared/http/httpserver"
	"github.com/cristiano-pacheco/shoplist/internal/shared/http/middleware"
	"github.com/cristiano-pacheco/shoplist/internal/shared/logger"
	"github.com/cristiano-pacheco/shoplist/internal/shared/mailer"
	"github.com/cristiano-pacheco/shoplist/internal/shared/parser"
	"github.com/cristiano-pacheco/shoplist/internal/shared/rabbitmq"
	"github.com/cristiano-pacheco/shoplist/internal/shared/registry"
	"github.com/cristiano-pacheco/shoplist/internal/shared/translator"
	"github.com/cristiano-pacheco/shoplist/internal/shared/validator"
	"go.uber.org/fx"
)

var Module = fx.Module(
	"shared",
	config.Module,
	database.Module,
	validator.Module,
	translator.Module,
	httpserver.Module,
	logger.Module,
	middleware.Module,
	registry.Module,
	parser.Module,
	mailer.Module,
	errs.Module,
	rabbitmq.Module,
)

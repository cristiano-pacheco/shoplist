package shared

import (
	"github.com/cristiano-pacheco/shoplist/internal/kernel/config"
	"github.com/cristiano-pacheco/shoplist/internal/kernel/database"
	"github.com/cristiano-pacheco/shoplist/internal/kernel/errs"
	"github.com/cristiano-pacheco/shoplist/internal/kernel/http/httpserver"
	"github.com/cristiano-pacheco/shoplist/internal/kernel/http/middleware"
	"github.com/cristiano-pacheco/shoplist/internal/kernel/jwt"
	"github.com/cristiano-pacheco/shoplist/internal/kernel/logger"
	"github.com/cristiano-pacheco/shoplist/internal/kernel/mailer"
	"github.com/cristiano-pacheco/shoplist/internal/kernel/rabbitmq"
	"github.com/cristiano-pacheco/shoplist/internal/kernel/registry"
	"github.com/cristiano-pacheco/shoplist/internal/kernel/translator"
	"github.com/cristiano-pacheco/shoplist/internal/kernel/validator"
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
	jwt.Module,
	mailer.Module,
	errs.Module,
	rabbitmq.Module,
)

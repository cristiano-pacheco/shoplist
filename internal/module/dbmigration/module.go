package dbmigration

import (
	"github.com/cristiano-pacheco/go-modulith/internal/module/dbmigration/service"
	"go.uber.org/fx"
)

var Module = fx.Module("dbmigration", fx.Invoke(service.ExecuteMigrations))

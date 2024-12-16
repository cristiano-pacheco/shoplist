package faktory

import (
	"github.com/cristiano-pacheco/shoplist/internal/shared/config"
	"github.com/cristiano-pacheco/shoplist/pkg/faktory"
)

type FaktoryFacade faktory.FaktoryFacade

func NewFaktoryFacade(cfg config.Config) FaktoryFacade {
	faktoryCfg := faktory.Config{
		Concurrency:     cfg.Faktory.Concurrency,
		PoolCapacity:    cfg.Faktory.PoolCapacity,
		ShutdownTimeout: cfg.Faktory.ShutdownTimeout,
	}
	return FaktoryFacade(faktory.NewFaktoryFacade(faktoryCfg))
}

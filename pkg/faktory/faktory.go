package faktory

import (
	"context"
	"time"

	client "github.com/contribsys/faktory/client"
	worker "github.com/contribsys/faktory_worker_go"
)

type FaktoryWorker struct {
	cfg     FactoryWorkerConfig
	manager *worker.Manager
}

func NewFaktoryWorker(cfg FactoryWorkerConfig) *FaktoryWorker {
	mgr := worker.NewManager()
	mgr.Concurrency = cfg.Concurrency
	mgr.ShutdownTimeout = time.Duration(cfg.ShutdownTimeout) * time.Second
	mgr.ProcessStrictPriorityQueues("critical", "default", "bulk")
	pool, err := client.NewPool(cfg.PoolCapacity)
	if err != nil {
		panic(err)
	}
	mgr.Pool = pool

	return &FaktoryWorker{cfg: cfg, manager: mgr}
}

func (w *FaktoryWorker) Register(jobs []Job) {
	for _, job := range jobs {
		w.manager.Register(job.Name, func(ctx context.Context, args ...interface{}) error {
			return job.Fn(ctx, args...)
		})
	}
}

func (w *FaktoryWorker) Start() error {
	err := w.manager.Run()
	if err != nil {
		return err
	}
	return nil
}

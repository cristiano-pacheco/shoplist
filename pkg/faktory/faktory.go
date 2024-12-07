package faktory

import (
	"context"
	"time"

	worker "github.com/contribsys/faktory_worker_go"
)

type FactoryWorker struct {
	cfg     FactoryWorkerConfig
	manager *worker.Manager
}

func NewFactoryWorker(cfg FactoryWorkerConfig) *FactoryWorker {
	mgr := worker.NewManager()
	mgr.Concurrency = cfg.Concurrency
	mgr.ShutdownTimeout = 25 * time.Second
	mgr.ProcessStrictPriorityQueues("critical", "default", "bulk")

	return &FactoryWorker{cfg: cfg, manager: mgr}
}

func (w *FactoryWorker) Register(jobs []Job) {
	for _, job := range jobs {
		w.manager.Register(job.Name, func(ctx context.Context, args ...interface{}) error {
			return job.Callback(ctx, args...)
		})
	}
}

func (w *FactoryWorker) Start() error {
	err := w.manager.Run()
	if err != nil {
		return err
	}

	return nil
}

package faktory

import (
	"context"
	"time"

	client "github.com/contribsys/faktory/client"
	worker "github.com/contribsys/faktory_worker_go"
)

type FaktoryFacade interface {
	Register(name string, fn JobFunc)
	Publish(job *Job) error
	Start() error
}

type faktoryFacade struct {
	cfg     Config
	manager *worker.Manager
	client  client.Client
}

func NewFaktoryFacade(cfg Config) FaktoryFacade {
	mgr := worker.NewManager()
	mgr.Concurrency = cfg.Concurrency
	mgr.ShutdownTimeout = time.Duration(cfg.ShutdownTimeout) * time.Second
	mgr.ProcessStrictPriorityQueues("critical", "default", "bulk")

	pool, err := client.NewPool(cfg.PoolCapacity)
	if err != nil {
		panic(err)
	}
	mgr.Pool = pool

	// Create a client for publishing jobs
	cl, err := client.NewPool(1)
	if err != nil {
		panic(err)
	}

	// Get a client from the pool
	publisher, err := cl.Get()
	if err != nil {
		panic(err)
	}

	return &faktoryFacade{
		cfg:     cfg,
		manager: mgr,
		client:  *publisher,
	}
}

func (f *faktoryFacade) Register(name string, fn JobFunc) {
	f.manager.Register(name, func(ctx context.Context, args ...interface{}) error {
		return fn(ctx, args...)
	})
}

func (f *faktoryFacade) Publish(job *Job) error {
	j := client.NewJob(job.Name, job.Args...)
	j.Queue = job.Queue
	j.Retry = &job.Retry

	return f.client.Push(j)
}

func (f *faktoryFacade) Start() error {
	return f.manager.Run()
}

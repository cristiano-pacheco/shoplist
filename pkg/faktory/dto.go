package faktory

import (
	"context"
)

type FactoryWorkerConfig struct {
	Concurrency     int
	PoolCapacity    int
	ShutdownTimeout int
}

type Job struct {
	Name string
	Fn   Fn
}

type Fn func(ctx context.Context, args ...interface{}) error

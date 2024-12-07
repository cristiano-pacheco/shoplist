package faktory

import "context"

type FactoryWorkerConfig struct {
	Concurrency int
	Jobs        []Job
}

type Job struct {
	Name     string
	Callback Callback
}

type Callback func(ctx context.Context, args ...interface{}) error

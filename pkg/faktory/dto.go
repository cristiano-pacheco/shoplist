package faktory

import (
	"context"
)

type Config struct {
	Concurrency     int
	PoolCapacity    int
	ShutdownTimeout int
}

// JobFunc represents a job callback function
type JobFunc func(ctx context.Context, args ...interface{}) error

// Job represents a Faktory job
type Job struct {
	Name  string
	Queue string
	Args  []interface{}
	Retry int
}

// NewJob creates a new job with default values
func NewJob(name string, args ...interface{}) *Job {
	return &Job{
		Name:  name,
		Queue: "default",
		Args:  args,
		Retry: 25,
	}
}

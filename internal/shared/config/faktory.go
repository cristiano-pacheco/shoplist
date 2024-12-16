package config

type Faktory struct {
	Provider        string `env:"PROVIDER"`
	Concurrency     int    `env:"CONCURRENCY" default:"10"`
	PoolCapacity    int    `env:"POOL_CAPACITY" default:"5"`
	ShutdownTimeout int    `env:"SHUTDOWN_TIMEOUT" default:"25"`
}

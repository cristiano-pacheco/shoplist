package config

type DB struct {
	DBHost             string `env:"HOST" envDefault:"localhost"`
	DBPort             uint   `env:"PORT" envDefault:"5432"`
	DBName             string `env:"NAME" envDefault:"controlweb"`
	DBUser             string `env:"USER" envDefault:"postgres"`
	DBPassword         string `env:"PASSWORD" envDefault:"postgres"`
	SSLMode            bool   `env:"SSL_MODE" envDefault:"false"`
	PrepareSTMT        bool   `env:"PREPARE_STMT" envDefault:"false"`
	EnableLogs         bool   `env:"ENABLE_LOGS" envDefault:"false"`
	MaxOpenConnections int    `env:"MAX_OPEN_CONNECTIONS" envDefault:"10"`
	MaxIdleConnections int    `env:"MAX_IDLE_CONNECTIONS" envDefault:"10"`
}

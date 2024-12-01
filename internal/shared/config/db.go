package config

type DB struct {
	Host               string `env:"HOST" envDefault:"localhost"`
	Name               string `env:"NAME" envDefault:"shoplist"`
	User               string `env:"USER" envDefault:"postgres"`
	Password           string `env:"PASSWORD" envDefault:"postgres"`
	Port               uint   `env:"PORT" envDefault:"5432"`
	MaxOpenConnections int    `env:"MAX_OPEN_CONNECTIONS" envDefault:"10"`
	MaxIdleConnections int    `env:"MAX_IDLE_CONNECTIONS" envDefault:"10"`
	SSLMode            bool   `env:"SSL_MODE" envDefault:"false"`
	PrepareSTMT        bool   `env:"PREPARE_STMT" envDefault:"false"`
	EnableLogs         bool   `env:"ENABLE_LOGS" envDefault:"false"`
}

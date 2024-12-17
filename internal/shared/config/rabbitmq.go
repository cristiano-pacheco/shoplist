package config

type RabbitMQ struct {
	Host     string `env:"HOST" envDefault:"localhost"`
	Port     string `env:"PORT" envDefault:"5672"`
	Username string `env:"USERNAME" envDefault:"guest"`
	Password string `env:"PASSWORD" envDefault:"guest"`
	VHost    string `env:"VHOST" envDefault:"/"`
}

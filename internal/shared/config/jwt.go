package config

type JWT struct {
	PrivateKey          string `env:"PRIVATE_KEY"`
	Issuer              string `env:"ISSUER"`
	ExpirationInSeconds int64  `env:"EXPIRATION_IN_SECONDS"`
}

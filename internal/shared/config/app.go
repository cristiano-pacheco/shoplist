package config

type App struct {
	Name    string `env:"NAME"`
	BaseURL string `env:"BASE_URL"`
	Version string `env:"VERSION"`
}

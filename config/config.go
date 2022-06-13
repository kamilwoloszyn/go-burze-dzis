package config

type Config struct {
	APIURL string `env:"API_URL"`
	APIKey string `env:"API_KEY"`
}

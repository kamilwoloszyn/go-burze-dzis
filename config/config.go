package config

type Config struct {
	APIKey  string `env:"API_KEY"`
	APIHost string `env:"API_HOST"`
}

func NewConfig(
	apiKey string,
	apiHost string,
) *Config {
	return &Config{
		APIKey:  apiKey,
		APIHost: apiHost,
	}
}

package config

// Config contains necessary params to start the client.
type Config struct {
	APIKey  string `env:"API_KEY"`
	APIHost string `env:"API_HOST"`
}

// NewConfig creates a Config from api key and host
func NewConfig(
	apiKey string,
	apiHost string,
) *Config {
	return &Config{
		APIKey:  apiKey,
		APIHost: apiHost,
	}
}

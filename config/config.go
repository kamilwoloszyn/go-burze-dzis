package config

type Config struct {
	APIURL string `env:"API_URL"`
	APIKey string `env:"API_KEY"`
}

func NewConfig(
	apiUrl string,
	apiKey string,
) *Config {
	return &Config{
		APIURL: apiUrl,
		APIKey: apiKey,
	}
}

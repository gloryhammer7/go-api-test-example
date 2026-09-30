package config

import "os"

type Config struct {
	BaseURL string
}

func Load() Config {
	baseURL := os.Getenv("API_BASE_URL")

	if baseURL == "" {
		baseURL = "http://localhost:8080"
	}

	return Config{
		BaseURL: baseURL,
	}
}
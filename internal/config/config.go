package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
)

type Config struct {
	Port               string `json:"port" env:"PORT"`
	GoogleSearchAPIKey string `json:"google_search_api_key" env:"GOOGLE_SEARCH_API_KEY"`
	GoogleSearchCX     string `json:"google_search_cx" env:"GOOGLE_SEARCH_CX"`
	GeminiAPIKey       string `json:"gemini_api_key" env:"GEMINI_API_KEY"`
}

func New() (*Config, error) {
	_ = godotenv.Load()

	cfg := &Config{
		Port:               getEnv("PORT", "8080"),
		GoogleSearchAPIKey: getEnv("GOOGLE_SEARCH_API_KEY", ""),
		GoogleSearchCX:     getEnv("GOOGLE_SEARCH_CX", ""),
		GeminiAPIKey:       getEnv("GEMINI_API_KEY", ""),
	}

	if cfg.GoogleSearchAPIKey == "" || cfg.GoogleSearchCX == "" || cfg.GeminiAPIKey == "" {
		return nil, fmt.Errorf("some required environment variables are missing")
	}

	return cfg, nil
}
func getEnv(key string, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}

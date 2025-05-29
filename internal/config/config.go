package config

import (
	"os"

	"github.com/joho/godotenv"
)

// Config holds all configuration for the application
type Config struct {
	Port   string
	BapID  string
	BapURI string
	BppURI string
}

// Load loads the configuration from environment variables
func Load() (*Config, error) {
	// Load .env file if it exists
	_ = godotenv.Load()

	return &Config{
		Port:   getEnv("PORT", "3000"),
		BapID:  getEnv("BAP_ID", "marblex-bap-preprod.marblex.ai"),
		BapURI: getEnv("BAP_URI", "https://evolving-fawn-cleanly.ngrok-free.app/v1"),
		BppURI: getEnv("BPP_URI", "https://dev-automation.ondc.org/api-service/ONDC:TRV11/2.0.0/seller"),
	}, nil
}

// getEnv gets an environment variable or returns a default value
func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

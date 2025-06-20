package config

import (
	"os"

	"github.com/joho/godotenv"
)

// Config holds all configuration for the application
type Config struct {
	// Server configuration
	Port   string
	BapID  string
	BapURI string
	BppURI string

	// Authentication configuration
	PrivateKey   string
	PublicKey    string
	SubscriberID string
	UniqueKeyID  string
}

// Load loads the configuration from environment variables
func Load() (*Config, error) {
	// Load .env file if it exists
	_ = godotenv.Load()

	return &Config{
		// Server configuration
		Port:   getEnv("PORT", "4000"),
		BapID:  getEnv("BAP_ID", ""),
		BapURI: getEnv("BAP_URI", ""),
		BppURI: getEnv("BPP_URI", "https://dev-automation.ondc.org/api-service/ONDC:TRV10/2.1.0/seller"),

		// Authentication configuration
		PrivateKey:   getEnv("PRIVATE_KEY", ""),
		PublicKey:    getEnv("PUBLIC_KEY", ""),
		SubscriberID: getEnv("SUBSCRIBER_ID", "buyer-app.ondc.org"),
		UniqueKeyID:  getEnv("UNIQUE_KEY_ID", "207"),
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

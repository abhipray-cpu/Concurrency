package main

import (
	"os"
)

// Config struct to hold all application configurations
type Config struct {
	ElasticsearchURL string
	// Add other configuration fields as needed
}

// InitializeConfig initializes application configuration from environment variables
// InitializeConfig initializes application configuration from environment variables
func InitializeConfig() *Config {
	appEnv := getEnv("APP_ENV", "dev") // Default to 'dev' if not specified

	var elasticsearchURL string

	switch appEnv {
	case "dev":
		elasticsearchURL = getEnv("ELASTICSEARCH_URL_DEV", "http://localhost:9200")
	case "stage":
		elasticsearchURL = getEnv("ELASTICSEARCH_URL_STAGE", "http://stage-elastic:9200")
	case "prod":
		elasticsearchURL = getEnv("ELASTICSEARCH_URL_PROD", "http://prod-elastic:9200")
	default:
		elasticsearchURL = getEnv("ELASTICSEARCH_URL", "http://localhost:9200")
	}

	return &Config{
		ElasticsearchURL: elasticsearchURL,
		// Initialize other configurations here based on appEnv if needed
	}
}

// getEnv is a helper function to read an environment variable or return a default value
func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

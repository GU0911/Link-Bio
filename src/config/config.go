package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Config holds all configuration for the application
type Config struct {
	ServerPort string
	DB         DatabaseConfig
}

// DatabaseConfig holds all configuration for the database connection
type DatabaseConfig struct {
	Host     string
	User     string
	Password string
	DBName   string
	Port     string
}

// LoadConfig loads application configuration from environment variables
func LoadConfig() Config {
	// godotenv.Load() will ignore the error if .env file is not found
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: .env file not found, using system environment variables")
	}

	return Config{
		ServerPort: getEnv("PORT", "8080"),
		DB: DatabaseConfig{
			Host:     getEnv("POSTGRES_HOST", "db"),
			User:     getEnv("POSTGRES_USER", "admin"),
			Password: getEnv("POSTGRES_PASSWORD", "secret"),
			DBName:   getEnv("POSTGRES_DB", "linkbio_go_db"),
			Port:     getEnv("POSTGRES_PORT", "5432"),
		},
	}
}

// Helper function to get an environment variable or return a default value
func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
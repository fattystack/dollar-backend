package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	SupabaseUrl        string
	SupabaseServiceKey string
	Port               string
	Environment        string
}

func LoadConfig() *Config {
	if os.Getenv("ENVIRONMENT") != "production" {
		if err := godotenv.Load(); err != nil {
			log.Print("No .env file found")
		}
	}

	return &Config{
		SupabaseUrl:        getEnv("SupabaseUrl", ""),
		SupabaseServiceKey: getEnv("SupabaseServiceKey", ""),
		Port:               getEnv("PORT", "8080"),
		Environment:        getEnv("ENVIRONMENT", "development"),
	}
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

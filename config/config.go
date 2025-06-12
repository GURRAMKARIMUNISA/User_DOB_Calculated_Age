package config

import (
	"log"
	"os"
)

type Config struct {
	DatabaseURL string
	Port        string
	Environment string
}

func LoadConfig() *Config {
	// Directly get environment variables.
	// When running with 'docker run --env-file .env', Docker injects these.
	// When running locally outside Docker, they can be set manually or via system environment vars.

	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		// Fallback for local testing without docker or explicit env vars set
		dbURL = "postgres://user_api_user:user_api_password@localhost:5432/user_api_db?sslmode=disable"
		log.Println("DATABASE_URL not found in environment, using default for local development.")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
		log.Println("PORT not found in environment, using default 3000.")
	}

	env := os.Getenv("ENVIRONMENT")
	if env == "" {
		env = "development"
		log.Println("ENVIRONMENT not found in environment, using default development.")
	}

	return &Config{
		DatabaseURL: dbURL,
		Port:        port,
		Environment: env,
	}
}

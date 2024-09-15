package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	RESTPort string
	GRPCPort string
	DbURI    string
}

func LoadConfig() Config {
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: Error loading .env file. Using environment variables.")
	}

	return Config{
		RESTPort: getEnv("REST_PORT", "8080"),
		GRPCPort: getEnv("GRPC_PORT", "50051"),
		DbURI:    getEnv("DB_URI", "postgresql://postgres:14022014@localhost/expenses?sslmode=disable"),
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
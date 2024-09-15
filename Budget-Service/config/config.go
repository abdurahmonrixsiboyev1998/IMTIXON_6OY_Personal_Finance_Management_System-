package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	MongoURI  string
	RedisAddr string
}

func LoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return &Config{
		MongoURI:  os.Getenv("MONGO_URI"),
		RedisAddr: os.Getenv("REDIS_ADDR"),
	}
}
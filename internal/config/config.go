package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	ApiKey    string
	SecretKey string
}

func LoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return &Config{
		ApiKey:    os.Getenv("DADATA_API_KEY"),
		SecretKey: os.Getenv("DADATA_SECRET_KEY"),
	}
}

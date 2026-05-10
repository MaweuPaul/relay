package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct{ Port string }

func Load() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("No .env file found ,using default values")
	}
	return &Config{
		Port: getEnv("PORT", "9090"),
	}
}

func getEnv(key, defaultValue string) string {

	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

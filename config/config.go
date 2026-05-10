package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port string
}

func Load() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file . Make sure it exists and is properly formatted. Using default values")
	}
	return &Config{
		Port: getEnv("PORT", "8080"),
	}
}

func getEnv(key, defaultString string) string {
	value := os.Getenv(key)
	if value == "" {
		value = defaultString
	}
	return value
}

package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port string
	Host string
}

func Load() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file . Make sure it exists and is properly formatted. Using default values")
	}
	return &Config{
		Port: getEnv("PORT", "8080"),
		Host: getEnv("HOST", "localhost"),
	}
}

func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		log.Printf("%s not set in .env, using default: %s\n", key, defaultValue)
		return defaultValue
	}
	return value
}

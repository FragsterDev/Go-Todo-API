package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type EnvConfig struct {
	MONGO_URI string
	PORT      string
	DBNAME    string
}

func LoadConfig() *EnvConfig {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading environment variables: %w", err)
	}

	envconfig := &EnvConfig{
		MONGO_URI: os.Getenv("MONGO"),
		PORT:      os.Getenv("PORT"),
		DBNAME:    os.Getenv("DATABASE"),
	}

	return envconfig
}

package config

import (
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
)

type Config struct {
	Port         string
	Dsn          string
	CacheHost    string
	CachePort    string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
	IdleTimeout  time.Duration
	Environment  Environment
}

type Environment string

const (
	EnvDevelopment Environment = "development"
	EnvProduction  Environment = "production"
	EnvStaging     Environment = "staging"
)

func Load() *Config {
	godotenv.Load()

	return &Config{
		Port: getEnvVar("PORT", ":3000"),
		Dsn:  getEnvVar("POSTGRES_DSN", ""),
		// CacheHost:    getEnvVar("CACHE_HOST", ""),
		// CachePort:    getEnvVar("CACHE_PORT", ""),
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
		Environment:  getEnv(),
	}
}

func getEnvVar(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}

	return defaultValue
}

func getEnv() Environment {
	env := os.Getenv("ENVIRONMENT")

	switch Environment(env) {
	case EnvDevelopment, EnvProduction, EnvStaging:
		return Environment(env)
	default:
		log.Fatalf("Invalid Environment value: %s", env)
		return EnvDevelopment
	}
}

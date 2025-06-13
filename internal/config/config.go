package config

import (
	"log"
	"os"
)

type Config struct {
	AppPort      string
	PostgresHost string
	PostgresUser string
	PostgresPass string
	PostgresDB   string
	RedisAddr    string
}

func LoadConfig() *Config {
	cfg := &Config{
		AppPort:      getEnv("APP_PORT", "8080"),
		PostgresHost: getEnv("POSTGRES_HOST", "localhost"),
		PostgresUser: getEnv("POSTGRES_USER", "postgres"),
		PostgresPass: getEnv("POSTGRES_PASSWORD", "postgres"),
		PostgresDB:   getEnv("POSTGRES_DB", "inote"),
		RedisAddr:    getEnv("REDIS_ADDR", "localhost:6379"),
	}
	return cfg
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	log.Printf("Warning: %s not set, using default: %s", key, fallback)
	return fallback
}

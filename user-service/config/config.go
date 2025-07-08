package config

import (
	"os"
	"strconv"
	"time"
)

type Config struct {
	AppName string
	Port    int

	DatabaseURL string

	ReadTimeout     time.Duration
	WriteTimeout    time.Duration
	IdleTimeout     time.Duration
	ShutdownTimeout time.Duration
}

func Load() *Config {
	cfg := &Config{
		AppName:         getEnv("APP_NAME", "user-service"),
		Port:            getEnvAsInt("PORT", 8080),
		DatabaseURL:     getEnv("DATABASE_URL", "postgres://postgres:1234@localhost:5432/multi-tenent-e-com?sslmode=disable"),
		ReadTimeout:     getEnvAsDuration("READ_TIMEOUT", 10*time.Second),
		WriteTimeout:    getEnvAsDuration("WRITE_TIMEOUT", 10*time.Second),
		IdleTimeout:     getEnvAsDuration("IDLE_TIMEOUT", 60*time.Second),
		ShutdownTimeout: getEnvAsDuration("SHUTDOWN_TIMEOUT", 15*time.Second),
	}
	return cfg
}

// Helper functions
func getEnv(key, fallback string) string {
	if v, ok := os.LookupEnv(key); ok {
		return v
	}
	return fallback
}

func getEnvAsInt(name string, defaultVal int) int {
	if valStr := os.Getenv(name); valStr != "" {
		if val, err := strconv.Atoi(valStr); err == nil {
			return val
		}
	}
	return defaultVal
}

func getEnvAsDuration(name string, defaultVal time.Duration) time.Duration {
	if valStr := os.Getenv(name); valStr != "" {
		if val, err := time.ParseDuration(valStr); err == nil {
			return val
		}
	}
	return defaultVal
}

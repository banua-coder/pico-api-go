package config

import (
	"log"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

type Config struct {
	Database  DatabaseConfig
	Server    ServerConfig
	RateLimit RateLimitConfig
}

type DatabaseConfig struct {
	Host            string
	Port            int
	Username        string
	Password        string
	DBName          string
	MaxOpenConns    int
	MaxIdleConns    int
	ConnMaxLifetime time.Duration
	ConnMaxIdleTime time.Duration
}

type ServerConfig struct {
	Port int
	Host string
}

type RateLimitConfig struct {
	Enabled           bool
	RequestsPerMinute int
	BurstSize         int
	WindowSize        time.Duration
}

func Load() *Config {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using environment variables or defaults")
	}

	return &Config{
		Database: DatabaseConfig{
			Host:            getEnv("DB_HOST", "127.0.0.1"), // Changed default to 127.0.0.1
			Port:            getEnvAsInt("DB_PORT", 3306),
			Username:        getEnv("DB_USERNAME", ""),
			Password:        getEnv("DB_PASSWORD", ""),
			DBName:          getEnv("DB_NAME", ""),
			MaxOpenConns:    getEnvAsInt("MYSQL_MAX_OPEN_CONNS", 5),
			MaxIdleConns:    getEnvAsInt("MYSQL_MAX_IDLE_CONNS", 2),
			ConnMaxLifetime: getEnvAsDuration("MYSQL_CONN_MAX_LIFETIME", 30*time.Second),
			ConnMaxIdleTime: getEnvAsDuration("MYSQL_CONN_MAX_IDLE_TIME", 15*time.Second),
		},
		Server: ServerConfig{
			Port: getEnvAsInt("SERVER_PORT", 8080),
			Host: getEnv("SERVER_HOST", "localhost"),
		},
		RateLimit: RateLimitConfig{
			Enabled:           getEnvAsBool("RATE_LIMIT_ENABLED", true),
			RequestsPerMinute: getEnvAsInt("RATE_LIMIT_REQUESTS_PER_MINUTE", 100),
			BurstSize:         getEnvAsInt("RATE_LIMIT_BURST_SIZE", 20),
			WindowSize:        getEnvAsDuration("RATE_LIMIT_WINDOW_SIZE", 1*time.Minute),
		},
	}
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func getEnvAsInt(key string, defaultValue int) int {
	if value := os.Getenv(key); value != "" {
		if intValue, err := strconv.Atoi(value); err == nil {
			return intValue
		}
	}
	return defaultValue
}

func getEnvAsDuration(key string, defaultValue time.Duration) time.Duration {
	if value := os.Getenv(key); value != "" {
		if duration, err := time.ParseDuration(value); err == nil {
			return duration
		}
	}
	return defaultValue
}

func getEnvAsBool(key string, defaultValue bool) bool {
	if value := os.Getenv(key); value != "" {
		if boolValue, err := strconv.ParseBool(value); err == nil {
			return boolValue
		}
	}
	return defaultValue
}

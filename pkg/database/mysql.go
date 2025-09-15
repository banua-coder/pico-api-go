package database

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"math"
	"time"

	"github.com/banua-coder/pico-api-go/internal/config"
	_ "github.com/go-sql-driver/mysql"
)

type DB struct {
	*sql.DB
}

type ConnectionConfig struct {
	MaxOpenConns    int
	MaxIdleConns    int
	ConnMaxLifetime time.Duration
	ConnMaxIdleTime time.Duration
	RetryAttempts   int
	RetryDelay      time.Duration
}

func NewMySQLConnection(cfg *config.DatabaseConfig) (*DB, error) {
	connCfg := ConnectionConfig{
		MaxOpenConns:    cfg.MaxOpenConns,
		MaxIdleConns:    cfg.MaxIdleConns,
		ConnMaxLifetime: cfg.ConnMaxLifetime,
		ConnMaxIdleTime: cfg.ConnMaxIdleTime,
		RetryAttempts:   3,
		RetryDelay:      1 * time.Second,
	}

	return NewMySQLConnectionWithConfig(cfg, connCfg)
}

func NewMySQLConnectionWithConfig(cfg *config.DatabaseConfig, connCfg ConnectionConfig) (*DB, error) {
	// Enhanced DSN with better timeout and connection parameters for shared hosting
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local&timeout=10s&readTimeout=10s&writeTimeout=10s&maxAllowedPacket=0&tls=false&allowOldPasswords=1&clientFoundRows=false&columnsWithAlias=false&interpolateParams=true",
		cfg.Username,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.DBName,
	)

	var db *sql.DB
	var err error

	// Retry connection with exponential backoff
	for attempt := 1; attempt <= connCfg.RetryAttempts; attempt++ {
		log.Printf("Attempting to connect to database (attempt %d/%d)", attempt, connCfg.RetryAttempts)

		db, err = sql.Open("mysql", dsn)
		if err != nil {
			if attempt == connCfg.RetryAttempts {
				return nil, fmt.Errorf("failed to open database connection after %d attempts: %w", connCfg.RetryAttempts, err)
			}

			backoffDelay := time.Duration(math.Pow(2, float64(attempt-1))) * connCfg.RetryDelay
			log.Printf("Database connection failed (attempt %d), retrying in %v: %v", attempt, backoffDelay, err)
			time.Sleep(backoffDelay)
			continue
		}

		// Configure connection pool
		db.SetMaxOpenConns(connCfg.MaxOpenConns)
		db.SetMaxIdleConns(connCfg.MaxIdleConns)
		db.SetConnMaxLifetime(connCfg.ConnMaxLifetime)
		db.SetConnMaxIdleTime(connCfg.ConnMaxIdleTime)

		// Test the connection
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		if err = db.PingContext(ctx); err != nil {
			if closeErr := db.Close(); closeErr != nil {
				log.Printf("Error closing database connection: %v", closeErr)
			}
			if attempt == connCfg.RetryAttempts {
				return nil, fmt.Errorf("failed to ping database after %d attempts: %w", connCfg.RetryAttempts, err)
			}

			backoffDelay := time.Duration(math.Pow(2, float64(attempt-1))) * connCfg.RetryDelay
			log.Printf("Database ping failed (attempt %d), retrying in %v: %v", attempt, backoffDelay, err)
			time.Sleep(backoffDelay)
			continue
		}

		log.Printf("Database connection established successfully on attempt %d", attempt)
		break
	}

	return &DB{db}, nil
}

func DefaultConnectionConfig() ConnectionConfig {
	return ConnectionConfig{
		MaxOpenConns:    5,                // Very conservative for shared hosting
		MaxIdleConns:    2,                // Minimal idle connections to prevent timeouts
		ConnMaxLifetime: 30 * time.Second, // Very short-lived connections for shared hosting
		ConnMaxIdleTime: 15 * time.Second, // Close idle connections very quickly
		RetryAttempts:   3,
		RetryDelay:      1 * time.Second,
	}
}

// HealthCheck performs a health check on the database connection
func (db *DB) HealthCheck() error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := db.DB.PingContext(ctx); err != nil {
		return fmt.Errorf("database health check failed: %w", err)
	}

	// Perform a simple query to ensure the database is responsive
	var result int
	if err := db.DB.QueryRowContext(ctx, "SELECT 1").Scan(&result); err != nil {
		return fmt.Errorf("database query test failed: %w", err)
	}

	return nil
}

// GetConnectionStats returns database connection statistics
func (db *DB) GetConnectionStats() sql.DBStats {
	return db.DB.Stats()
}

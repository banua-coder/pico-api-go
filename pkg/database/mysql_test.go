package database

import (
	"testing"
	"time"

	"github.com/banua-coder/pico-api-go/internal/config"
	"github.com/stretchr/testify/assert"
)

func TestDefaultConnectionConfig(t *testing.T) {
	cfg := DefaultConnectionConfig()

	assert.Equal(t, 5, cfg.MaxOpenConns)
	assert.Equal(t, 2, cfg.MaxIdleConns)
	assert.Equal(t, 30*time.Second, cfg.ConnMaxLifetime)
	assert.Equal(t, 15*time.Second, cfg.ConnMaxIdleTime)
	assert.Equal(t, 3, cfg.RetryAttempts)
	assert.Equal(t, 1*time.Second, cfg.RetryDelay)
}

func TestNewMySQLConnectionWithConfig_ConnectionFailed(t *testing.T) {
	cfg := &config.DatabaseConfig{
		Host:     "127.0.0.1",
		Port:     1, // invalid port
		Username: "user",
		Password: "pass",
		DBName:   "testdb",
	}
	connCfg := ConnectionConfig{
		MaxOpenConns:    1,
		MaxIdleConns:    1,
		ConnMaxLifetime: 1 * time.Second,
		ConnMaxIdleTime: 1 * time.Second,
		RetryAttempts:   1,
		RetryDelay:      1 * time.Millisecond,
	}

	_, err := NewMySQLConnectionWithConfig(cfg, connCfg)
	assert.Error(t, err)
}

func TestNewMySQLConnection_ConnectionFailed(t *testing.T) {
	cfg := &config.DatabaseConfig{
		Host:            "127.0.0.1",
		Port:            1,
		Username:        "user",
		Password:        "pass",
		DBName:          "testdb",
		MaxOpenConns:    1,
		MaxIdleConns:    1,
		ConnMaxLifetime: 1 * time.Second,
		ConnMaxIdleTime: 1 * time.Second,
	}

	_, err := NewMySQLConnection(cfg)
	assert.Error(t, err)
}

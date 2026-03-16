package config

import (
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func unsetEnvVars(keys ...string) {
	for _, k := range keys {
		_ = os.Unsetenv(k)
	}
}

func TestLoad_Defaults(t *testing.T) {
	unsetEnvVars("DB_HOST", "DB_PORT", "DB_USERNAME", "DB_PASSWORD", "DB_NAME",
		"SERVER_PORT", "SERVER_HOST", "RATE_LIMIT_ENABLED", "RATE_LIMIT_REQUESTS_PER_MINUTE",
		"RATE_LIMIT_BURST_SIZE", "RATE_LIMIT_WINDOW_SIZE",
		"MYSQL_MAX_OPEN_CONNS", "MYSQL_MAX_IDLE_CONNS", "MYSQL_CONN_MAX_LIFETIME", "MYSQL_CONN_MAX_IDLE_TIME")

	cfg := Load()

	assert.Equal(t, "127.0.0.1", cfg.Database.Host)
	assert.Equal(t, 3306, cfg.Database.Port)
	assert.Equal(t, 5, cfg.Database.MaxOpenConns)
	assert.Equal(t, 2, cfg.Database.MaxIdleConns)
	assert.Equal(t, 30*time.Second, cfg.Database.ConnMaxLifetime)
	assert.Equal(t, 15*time.Second, cfg.Database.ConnMaxIdleTime)
	assert.Equal(t, 8080, cfg.Server.Port)
	assert.Equal(t, "localhost", cfg.Server.Host)
	assert.True(t, cfg.RateLimit.Enabled)
	assert.Equal(t, 100, cfg.RateLimit.RequestsPerMinute)
	assert.Equal(t, 20, cfg.RateLimit.BurstSize)
	assert.Equal(t, 1*time.Minute, cfg.RateLimit.WindowSize)
}

func TestLoad_FromEnv(t *testing.T) {
	require.NoError(t, os.Setenv("DB_HOST", "db.example.com"))
	require.NoError(t, os.Setenv("DB_PORT", "5432"))
	require.NoError(t, os.Setenv("DB_USERNAME", "admin"))
	require.NoError(t, os.Setenv("DB_PASSWORD", "secret"))
	require.NoError(t, os.Setenv("DB_NAME", "pico_db"))
	require.NoError(t, os.Setenv("SERVER_PORT", "9090"))
	require.NoError(t, os.Setenv("RATE_LIMIT_ENABLED", "false"))
	require.NoError(t, os.Setenv("RATE_LIMIT_REQUESTS_PER_MINUTE", "200"))
	t.Cleanup(func() {
		unsetEnvVars("DB_HOST", "DB_PORT", "DB_USERNAME", "DB_PASSWORD", "DB_NAME",
			"SERVER_PORT", "RATE_LIMIT_ENABLED", "RATE_LIMIT_REQUESTS_PER_MINUTE")
	})

	cfg := Load()

	assert.Equal(t, "db.example.com", cfg.Database.Host)
	assert.Equal(t, 5432, cfg.Database.Port)
	assert.Equal(t, "admin", cfg.Database.Username)
	assert.Equal(t, "secret", cfg.Database.Password)
	assert.Equal(t, "pico_db", cfg.Database.DBName)
	assert.Equal(t, 9090, cfg.Server.Port)
	assert.False(t, cfg.RateLimit.Enabled)
	assert.Equal(t, 200, cfg.RateLimit.RequestsPerMinute)
}

func TestGetEnv_Default(t *testing.T) {
	unsetEnvVars("TEST_KEY_FORGE")
	assert.Equal(t, "default_val", getEnv("TEST_KEY_FORGE", "default_val"))
}

func TestGetEnv_FromEnv(t *testing.T) {
	require.NoError(t, os.Setenv("TEST_KEY_FORGE", "actual_val"))
	t.Cleanup(func() { unsetEnvVars("TEST_KEY_FORGE") })
	assert.Equal(t, "actual_val", getEnv("TEST_KEY_FORGE", "default_val"))
}

func TestGetEnvAsInt_Default(t *testing.T) {
	unsetEnvVars("TEST_INT_FORGE")
	assert.Equal(t, 42, getEnvAsInt("TEST_INT_FORGE", 42))
}

func TestGetEnvAsInt_Valid(t *testing.T) {
	require.NoError(t, os.Setenv("TEST_INT_FORGE", "99"))
	t.Cleanup(func() { unsetEnvVars("TEST_INT_FORGE") })
	assert.Equal(t, 99, getEnvAsInt("TEST_INT_FORGE", 42))
}

func TestGetEnvAsInt_Invalid(t *testing.T) {
	require.NoError(t, os.Setenv("TEST_INT_FORGE", "not_a_number"))
	t.Cleanup(func() { unsetEnvVars("TEST_INT_FORGE") })
	assert.Equal(t, 42, getEnvAsInt("TEST_INT_FORGE", 42))
}

func TestGetEnvAsDuration_Default(t *testing.T) {
	unsetEnvVars("TEST_DUR_FORGE")
	assert.Equal(t, 5*time.Second, getEnvAsDuration("TEST_DUR_FORGE", 5*time.Second))
}

func TestGetEnvAsDuration_Valid(t *testing.T) {
	require.NoError(t, os.Setenv("TEST_DUR_FORGE", "2m"))
	t.Cleanup(func() { unsetEnvVars("TEST_DUR_FORGE") })
	assert.Equal(t, 2*time.Minute, getEnvAsDuration("TEST_DUR_FORGE", 5*time.Second))
}

func TestGetEnvAsDuration_Invalid(t *testing.T) {
	require.NoError(t, os.Setenv("TEST_DUR_FORGE", "invalid_duration"))
	t.Cleanup(func() { unsetEnvVars("TEST_DUR_FORGE") })
	assert.Equal(t, 5*time.Second, getEnvAsDuration("TEST_DUR_FORGE", 5*time.Second))
}

func TestGetEnvAsBool_Default(t *testing.T) {
	unsetEnvVars("TEST_BOOL_FORGE")
	assert.True(t, getEnvAsBool("TEST_BOOL_FORGE", true))
}

func TestGetEnvAsBool_Valid(t *testing.T) {
	require.NoError(t, os.Setenv("TEST_BOOL_FORGE", "false"))
	t.Cleanup(func() { unsetEnvVars("TEST_BOOL_FORGE") })
	assert.False(t, getEnvAsBool("TEST_BOOL_FORGE", true))
}

func TestGetEnvAsBool_Invalid(t *testing.T) {
	require.NoError(t, os.Setenv("TEST_BOOL_FORGE", "notabool"))
	t.Cleanup(func() { unsetEnvVars("TEST_BOOL_FORGE") })
	assert.True(t, getEnvAsBool("TEST_BOOL_FORGE", true))
}

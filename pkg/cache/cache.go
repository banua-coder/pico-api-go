// Package cache provides caching utilities for pico-api-go.
package cache

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

// Options holds Redis connection options.
type Options struct {
	Addr     string
	Password string
	DB       int
}

// Cache defines the caching interface.
type Cache interface {
	Get(ctx context.Context, key string) (string, error)
	Set(ctx context.Context, key string, value string, ttl time.Duration) error
	Delete(ctx context.Context, key string) error
}

// RedisCache is a Redis-backed cache implementation.
type RedisCache struct {
	client     *redis.Client
	defaultTTL time.Duration
}

// New creates a new RedisCache with the given options and default TTL.
func New(opts Options, defaultTTL time.Duration) *RedisCache {
	client := redis.NewClient(&redis.Options{
		Addr:     opts.Addr,
		Password: opts.Password,
		DB:       opts.DB,
	})
	return &RedisCache{client: client, defaultTTL: defaultTTL}
}

// Ping verifies connectivity to Redis.
func (c *RedisCache) Ping(ctx context.Context) error {
	return c.client.Ping(ctx).Err()
}

// Get retrieves a value by key. Returns redis.Nil if not found.
func (c *RedisCache) Get(ctx context.Context, key string) (string, error) {
	val, err := c.client.Get(ctx, key).Result()
	if err != nil {
		return "", err
	}
	return val, nil
}

// Set stores a key-value pair with the given TTL (uses defaultTTL if ttl is 0).
func (c *RedisCache) Set(ctx context.Context, key string, value string, ttl time.Duration) error {
	if ttl == 0 {
		ttl = c.defaultTTL
	}
	return c.client.Set(ctx, key, value, ttl).Err()
}

// Delete removes a key from the cache.
func (c *RedisCache) Delete(ctx context.Context, key string) error {
	return c.client.Del(ctx, key).Err()
}

// Key builds a namespaced cache key.
func Key(parts ...string) string {
	key := "pico:v1"
	for _, p := range parts {
		key = fmt.Sprintf("%s:%s", key, p)
	}
	return key
}

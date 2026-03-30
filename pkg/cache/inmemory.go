package cache

import (
	"context"
	"sync"
	"time"
)

type entry struct {
	value     string
	expiresAt time.Time
}

// InMemoryCache is a simple in-memory cache backed by sync.Map. Used as a
// fallback when Redis is unavailable, or in tests.
type InMemoryCache struct {
	mu         sync.Map
	defaultTTL time.Duration
}

// NewInMemory creates a new InMemoryCache.
func NewInMemory(defaultTTL time.Duration) *InMemoryCache {
	return &InMemoryCache{defaultTTL: defaultTTL}
}

func (c *InMemoryCache) Get(_ context.Context, key string) (string, error) {
	v, ok := c.mu.Load(key)
	if !ok {
		return "", ErrCacheMiss
	}
	e := v.(entry)
	if time.Now().After(e.expiresAt) {
		c.mu.Delete(key)
		return "", ErrCacheMiss
	}
	return e.value, nil
}

func (c *InMemoryCache) Set(_ context.Context, key string, value string, ttl time.Duration) error {
	if ttl == 0 {
		ttl = c.defaultTTL
	}
	c.mu.Store(key, entry{value: value, expiresAt: time.Now().Add(ttl)})
	return nil
}

func (c *InMemoryCache) Delete(_ context.Context, key string) error {
	c.mu.Delete(key)
	return nil
}

// ErrCacheMiss is returned when a key is not found or has expired.
var ErrCacheMiss = &cacheMissError{}

type cacheMissError struct{}

func (e *cacheMissError) Error() string { return "cache miss" }

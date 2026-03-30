package cache

import (
	"strings"
	"sync"
	"time"
)

type entry struct {
	value     interface{}
	expiresAt time.Time
}

// Cache is a thread-safe in-memory cache with TTL support.
type Cache struct {
	mu         sync.RWMutex
	items      map[string]entry
	defaultTTL time.Duration
}

// New creates a new Cache with the given default TTL.
func New(defaultTTL time.Duration) *Cache {
	return &Cache{
		items:      make(map[string]entry),
		defaultTTL: defaultTTL,
	}
}

// Set stores a value in the cache. An optional TTL overrides the default.
func (c *Cache) Set(key string, value interface{}, ttl ...time.Duration) {
	d := c.defaultTTL
	if len(ttl) > 0 && ttl[0] > 0 {
		d = ttl[0]
	}
	c.mu.Lock()
	c.items[key] = entry{value: value, expiresAt: time.Now().Add(d)}
	c.mu.Unlock()
}

// Get retrieves a value from the cache. Returns (value, true) on hit, (nil, false) on miss or expiry.
func (c *Cache) Get(key string) (interface{}, bool) {
	c.mu.RLock()
	e, ok := c.items[key]
	c.mu.RUnlock()
	if !ok || time.Now().After(e.expiresAt) {
		return nil, false
	}
	return e.value, true
}

// Delete removes a single key from the cache.
func (c *Cache) Delete(key string) {
	c.mu.Lock()
	delete(c.items, key)
	c.mu.Unlock()
}

// DeletePrefix removes all keys that start with the given prefix.
func (c *Cache) DeletePrefix(prefix string) {
	c.mu.Lock()
	for k := range c.items {
		if strings.HasPrefix(k, prefix) {
			delete(c.items, k)
		}
	}
	c.mu.Unlock()
}

// Clear removes all entries from the cache.
func (c *Cache) Clear() {
	c.mu.Lock()
	c.items = make(map[string]entry)
	c.mu.Unlock()
}

// StartCleanup launches a background goroutine that removes expired entries at the given interval.
func (c *Cache) StartCleanup(interval time.Duration) {
	go func() {
		ticker := time.NewTicker(interval)
		defer ticker.Stop()
		for range ticker.C {
			now := time.Now()
			c.mu.Lock()
			for k, e := range c.items {
				if now.After(e.expiresAt) {
					delete(c.items, k)
				}
			}
			c.mu.Unlock()
		}
	}()
}

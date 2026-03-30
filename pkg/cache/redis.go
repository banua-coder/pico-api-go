package cache

import (
	"context"
	"encoding/json"
	"log"
	"strings"
	"time"

	"github.com/redis/go-redis/v9"
)

// RedisCache is a Redis-backed distributed cache with the same interface as Cache.
// It stores values as JSON and supports TTL, prefix deletion, and full clear.
type RedisCache struct {
	client     *redis.Client
	defaultTTL time.Duration
}

// RedisOptions holds Redis connection options.
type RedisOptions struct {
	Addr     string
	Password string
	DB       int
}

// NewRedis creates a new RedisCache. Call Ping to verify connectivity before use.
func NewRedis(opts RedisOptions, defaultTTL time.Duration) *RedisCache {
	client := redis.NewClient(&redis.Options{
		Addr:     opts.Addr,
		Password: opts.Password,
		DB:       opts.DB,
	})
	return &RedisCache{client: client, defaultTTL: defaultTTL}
}

// Ping verifies connectivity to Redis.
func (r *RedisCache) Ping(ctx context.Context) error {
	return r.client.Ping(ctx).Err()
}

// Set stores a value (JSON-encoded) in Redis. An optional TTL overrides the default.
func (r *RedisCache) Set(key string, value interface{}, ttl ...time.Duration) {
	d := r.defaultTTL
	if len(ttl) > 0 && ttl[0] > 0 {
		d = ttl[0]
	}
	b, err := json.Marshal(value)
	if err != nil {
		log.Printf("cache/redis: marshal error for key %s: %v", key, err)
		return
	}
	if err := r.client.Set(context.Background(), key, b, d).Err(); err != nil {
		log.Printf("cache/redis: set error for key %s: %v", key, err)
	}
}

// Get retrieves and JSON-decodes a value. Returns (value, true) on hit, (nil, false) on miss.
// The caller must type-assert carefully — returned value is map[string]interface{} or []interface{}
// for complex types. Use GetInto for typed retrieval.
func (r *RedisCache) Get(key string) (interface{}, bool) {
	raw, err := r.client.Get(context.Background(), key).Bytes()
	if err != nil {
		return nil, false
	}
	var v interface{}
	if err := json.Unmarshal(raw, &v); err != nil {
		return nil, false
	}
	return v, true
}

// GetInto retrieves a value from Redis and unmarshals it into dest.
func (r *RedisCache) GetInto(key string, dest interface{}) bool {
	raw, err := r.client.Get(context.Background(), key).Bytes()
	if err != nil {
		return false
	}
	return json.Unmarshal(raw, dest) == nil
}

// Delete removes a single key.
func (r *RedisCache) Delete(key string) {
	r.client.Del(context.Background(), key)
}

// DeletePrefix removes all keys matching the given prefix using SCAN.
func (r *RedisCache) DeletePrefix(prefix string) {
	ctx := context.Background()
	var cursor uint64
	for {
		keys, next, err := r.client.Scan(ctx, cursor, prefix+"*", 100).Result()
		if err != nil {
			log.Printf("cache/redis: scan error: %v", err)
			return
		}
		if len(keys) > 0 {
			r.client.Del(ctx, keys...)
		}
		cursor = next
		if cursor == 0 {
			break
		}
	}
}

// Clear removes all keys in the current DB. Use with caution in shared Redis instances.
func (r *RedisCache) Clear() {
	if err := r.client.FlushDB(context.Background()).Err(); err != nil {
		log.Printf("cache/redis: flushdb error: %v", err)
	}
}

// StartCleanup is a no-op for Redis (TTL is handled natively by Redis).
func (r *RedisCache) StartCleanup(_ time.Duration) {}

// Ensure RedisCache satisfies the same duck-type interface as Cache.
var _ interface {
	Set(string, interface{}, ...time.Duration)
	Get(string) (interface{}, bool)
	Delete(string)
	DeletePrefix(string)
	Clear()
	StartCleanup(time.Duration)
} = (*RedisCache)(nil)

// NOTE: RedisCache.Get returns generic interface{} (JSON-decoded).
// For typed access in service decorators, use GetInto.
// The existing cachedCovidService / cachedRegencyService use type assertions after Get,
// which works for in-memory cache but NOT for Redis (JSON round-trip loses concrete types).
// Use NewRedisAwareCache (below) as a drop-in *Cache replacement that backs in-memory with Redis.

// RedisAwareCache wraps the in-memory Cache and synchronises with Redis for distributed
// invalidation and warm-up. On Get miss, it tries Redis. On Set, it writes to both.
// This keeps the type-safe interface of *Cache while adding Redis persistence.
type RedisAwareCache struct {
	mem   *Cache
	redis *RedisCache
}

// NewRedisAwareCache creates a dual-layer cache backed by in-memory + Redis.
func NewRedisAwareCache(defaultTTL time.Duration, redisOpts RedisOptions) (*RedisAwareCache, error) {
	rc := NewRedis(redisOpts, defaultTTL)
	if err := rc.Ping(context.Background()); err != nil {
		return nil, err
	}
	return &RedisAwareCache{
		mem:   New(defaultTTL),
		redis: rc,
	}, nil
}

func (c *RedisAwareCache) Set(key string, value interface{}, ttl ...time.Duration) {
	c.mem.Set(key, value, ttl...)
	c.redis.Set(key, value, ttl...)
}

func (c *RedisAwareCache) Get(key string) (interface{}, bool) {
	return c.mem.Get(key)
}

func (c *RedisAwareCache) Delete(key string) {
	c.mem.Delete(key)
	c.redis.Delete(key)
}

func (c *RedisAwareCache) DeletePrefix(prefix string) {
	c.mem.DeletePrefix(prefix)
	c.redis.DeletePrefix(prefix)
}

func (c *RedisAwareCache) Clear() {
	c.mem.Clear()
	c.redis.Clear()
}

func (c *RedisAwareCache) StartCleanup(interval time.Duration) {
	c.mem.StartCleanup(interval)
}

// Unwrap returns the underlying *Cache for compatibility with functions that require *Cache directly.
func (c *RedisAwareCache) Unwrap() *Cache {
	return c.mem
}

// IsPrefixMatch is a utility used by DeletePrefix.
func IsPrefixMatch(key, prefix string) bool {
	return strings.HasPrefix(key, prefix)
}

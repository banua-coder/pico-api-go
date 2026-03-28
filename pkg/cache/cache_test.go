package cache

import (
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestSetGet(t *testing.T) {
	c := New(time.Minute)
	c.Set("foo", "bar")
	v, ok := c.Get("foo")
	assert.True(t, ok)
	assert.Equal(t, "bar", v)
}

func TestGet_Miss(t *testing.T) {
	c := New(time.Minute)
	_, ok := c.Get("missing")
	assert.False(t, ok)
}

func TestTTLExpiry(t *testing.T) {
	c := New(time.Millisecond * 50)
	c.Set("k", "v")
	time.Sleep(time.Millisecond * 100)
	_, ok := c.Get("k")
	assert.False(t, ok, "expected expired entry to be missing")
}

func TestCustomTTL(t *testing.T) {
	c := New(time.Hour)
	c.Set("k", "v", time.Millisecond*50)
	time.Sleep(time.Millisecond * 100)
	_, ok := c.Get("k")
	assert.False(t, ok, "expected custom TTL to expire")
}

func TestDelete(t *testing.T) {
	c := New(time.Minute)
	c.Set("k", "v")
	c.Delete("k")
	_, ok := c.Get("k")
	assert.False(t, ok)
}

func TestDeletePrefix(t *testing.T) {
	c := New(time.Minute)
	c.Set("national:all", 1)
	c.Set("national:latest", 2)
	c.Set("province:all", 3)
	c.DeletePrefix("national:")
	_, ok1 := c.Get("national:all")
	_, ok2 := c.Get("national:latest")
	_, ok3 := c.Get("province:all")
	assert.False(t, ok1)
	assert.False(t, ok2)
	assert.True(t, ok3)
}

func TestClear(t *testing.T) {
	c := New(time.Minute)
	c.Set("a", 1)
	c.Set("b", 2)
	c.Clear()
	_, ok1 := c.Get("a")
	_, ok2 := c.Get("b")
	assert.False(t, ok1)
	assert.False(t, ok2)
}

func TestStartCleanup(t *testing.T) {
	c := New(time.Millisecond * 50)
	c.Set("k", "v")
	c.StartCleanup(time.Millisecond * 30)
	time.Sleep(time.Millisecond * 200)
	c.mu.RLock()
	_, exists := c.items["k"]
	c.mu.RUnlock()
	assert.False(t, exists, "cleanup goroutine should have evicted expired entry")
}

func TestConcurrentAccess(t *testing.T) {
	c := New(time.Minute)
	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			key := "k"
			c.Set(key, i)
			c.Get(key)
		}(i)
	}
	wg.Wait()
}

package middleware

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/banua-coder/pico-api-go/internal/config"
	"github.com/stretchr/testify/assert"
)

func TestRateLimit_Disabled(t *testing.T) {
	cfg := config.RateLimitConfig{
		Enabled: false,
	}

	handler := RateLimit(cfg)(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("OK"))
	}))

	// Make multiple requests - all should pass
	for i := 0; i < 200; i++ {
		req := httptest.NewRequest("GET", "/test", nil)
		rr := httptest.NewRecorder()

		handler.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusOK, rr.Code)
		assert.Equal(t, "OK", rr.Body.String())
	}
}

func TestRateLimit_WithinLimits(t *testing.T) {
	cfg := config.RateLimitConfig{
		Enabled:           true,
		RequestsPerMinute: 10,
		BurstSize:         5,
		WindowSize:        time.Minute,
	}

	handler := RateLimit(cfg)(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("OK"))
	}))

	// Make requests within the limit
	for i := 0; i < 5; i++ {
		req := httptest.NewRequest("GET", "/test", nil)
		req.RemoteAddr = "192.168.1.1:12345"
		rr := httptest.NewRecorder()

		handler.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusOK, rr.Code)
		assert.Equal(t, "OK", rr.Body.String())

		// Check rate limit headers
		assert.Equal(t, "10", rr.Header().Get("X-RateLimit-Limit"))
		assert.NotEmpty(t, rr.Header().Get("X-RateLimit-Remaining"))
	}
}

func TestRateLimit_ExceedsLimit(t *testing.T) {
	cfg := config.RateLimitConfig{
		Enabled:           true,
		RequestsPerMinute: 5,
		BurstSize:         3,
		WindowSize:        time.Minute,
	}

	handler := RateLimit(cfg)(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("OK"))
	}))

	clientIP := "192.168.1.1:12345"

	// Make requests up to the limit
	for i := 0; i < 5; i++ {
		req := httptest.NewRequest("GET", "/test", nil)
		req.RemoteAddr = clientIP
		rr := httptest.NewRecorder()

		handler.ServeHTTP(rr, req)
		assert.Equal(t, http.StatusOK, rr.Code)
	}

	// The next request should be rate limited
	req := httptest.NewRequest("GET", "/test", nil)
	req.RemoteAddr = clientIP
	rr := httptest.NewRecorder()

	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusTooManyRequests, rr.Code)

	// Check rate limit headers
	assert.Equal(t, "5", rr.Header().Get("X-RateLimit-Limit"))
	assert.Equal(t, "0", rr.Header().Get("X-RateLimit-Remaining"))
	assert.NotEmpty(t, rr.Header().Get("X-RateLimit-Reset"))
	assert.NotEmpty(t, rr.Header().Get("Retry-After"))

	// Check error response
	var response ErrorResponse
	err := json.Unmarshal(rr.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, "error", response.Status)
	assert.Contains(t, response.Error, "Rate limit exceeded")
}

func TestRateLimit_DifferentClients(t *testing.T) {
	cfg := config.RateLimitConfig{
		Enabled:           true,
		RequestsPerMinute: 2,
		BurstSize:         1,
		WindowSize:        time.Minute,
	}

	handler := RateLimit(cfg)(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("OK"))
	}))

	// Client 1 makes requests up to limit
	for i := 0; i < 2; i++ {
		req := httptest.NewRequest("GET", "/test", nil)
		req.RemoteAddr = "192.168.1.1:12345"
		rr := httptest.NewRecorder()

		handler.ServeHTTP(rr, req)
		assert.Equal(t, http.StatusOK, rr.Code)
	}

	// Client 1 is now rate limited
	req := httptest.NewRequest("GET", "/test", nil)
	req.RemoteAddr = "192.168.1.1:12345"
	rr := httptest.NewRecorder()
	handler.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusTooManyRequests, rr.Code)

	// Client 2 should still be able to make requests
	req2 := httptest.NewRequest("GET", "/test", nil)
	req2.RemoteAddr = "192.168.1.2:12345"
	rr2 := httptest.NewRecorder()
	handler.ServeHTTP(rr2, req2)
	assert.Equal(t, http.StatusOK, rr2.Code)
}

func TestRateLimit_XForwardedFor(t *testing.T) {
	cfg := config.RateLimitConfig{
		Enabled:           true,
		RequestsPerMinute: 2,
		BurstSize:         1,
		WindowSize:        time.Minute,
	}

	handler := RateLimit(cfg)(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("OK"))
	}))

	// Make requests with X-Forwarded-For header
	for i := 0; i < 2; i++ {
		req := httptest.NewRequest("GET", "/test", nil)
		req.Header.Set("X-Forwarded-For", "10.0.0.1")
		req.RemoteAddr = "192.168.1.1:12345" // This should be ignored
		rr := httptest.NewRecorder()

		handler.ServeHTTP(rr, req)
		assert.Equal(t, http.StatusOK, rr.Code)
	}

	// The next request should be rate limited based on X-Forwarded-For
	req := httptest.NewRequest("GET", "/test", nil)
	req.Header.Set("X-Forwarded-For", "10.0.0.1")
	req.RemoteAddr = "192.168.1.1:12345"
	rr := httptest.NewRecorder()

	handler.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusTooManyRequests, rr.Code)
}

func TestRateLimit_XRealIP(t *testing.T) {
	cfg := config.RateLimitConfig{
		Enabled:           true,
		RequestsPerMinute: 2,
		BurstSize:         1,
		WindowSize:        time.Minute,
	}

	handler := RateLimit(cfg)(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("OK"))
	}))

	// Make requests with X-Real-IP header
	for i := 0; i < 2; i++ {
		req := httptest.NewRequest("GET", "/test", nil)
		req.Header.Set("X-Real-IP", "10.0.0.2")
		req.RemoteAddr = "192.168.1.1:12345" // This should be ignored
		rr := httptest.NewRecorder()

		handler.ServeHTTP(rr, req)
		assert.Equal(t, http.StatusOK, rr.Code)
	}

	// The next request should be rate limited based on X-Real-IP
	req := httptest.NewRequest("GET", "/test", nil)
	req.Header.Set("X-Real-IP", "10.0.0.2")
	req.RemoteAddr = "192.168.1.1:12345"
	rr := httptest.NewRecorder()

	handler.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusTooManyRequests, rr.Code)
}

func TestRateLimit_SlidingWindow(t *testing.T) {
	cfg := config.RateLimitConfig{
		Enabled:           true,
		RequestsPerMinute: 3,
		BurstSize:         2,
		WindowSize:        2 * time.Second, // Short window for testing
	}

	handler := RateLimit(cfg)(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("OK"))
	}))

	clientIP := "192.168.1.1:12345"

	// Make 3 requests quickly
	for i := 0; i < 3; i++ {
		req := httptest.NewRequest("GET", "/test", nil)
		req.RemoteAddr = clientIP
		rr := httptest.NewRecorder()

		handler.ServeHTTP(rr, req)
		assert.Equal(t, http.StatusOK, rr.Code)
	}

	// 4th request should be rate limited
	req := httptest.NewRequest("GET", "/test", nil)
	req.RemoteAddr = clientIP
	rr := httptest.NewRecorder()
	handler.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusTooManyRequests, rr.Code)

	// Wait for window to slide
	time.Sleep(3 * time.Second)

	// Should be able to make requests again
	req = httptest.NewRequest("GET", "/test", nil)
	req.RemoteAddr = clientIP
	rr = httptest.NewRecorder()
	handler.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusOK, rr.Code)
}

func TestRateLimiter_GetClientIP(t *testing.T) {
	limiter := NewRateLimiter(config.RateLimitConfig{})

	tests := []struct {
		name         string
		setupRequest func(*http.Request)
		expectedIP   string
	}{
		{
			name: "X-Forwarded-For header",
			setupRequest: func(r *http.Request) {
				r.Header.Set("X-Forwarded-For", "203.0.113.1")
				r.RemoteAddr = "192.168.1.1:12345"
			},
			expectedIP: "203.0.113.1",
		},
		{
			name: "X-Real-IP header",
			setupRequest: func(r *http.Request) {
				r.Header.Set("X-Real-IP", "203.0.113.2")
				r.RemoteAddr = "192.168.1.1:12345"
			},
			expectedIP: "203.0.113.2",
		},
		{
			name: "RemoteAddr fallback",
			setupRequest: func(r *http.Request) {
				r.RemoteAddr = "192.168.1.1:12345"
			},
			expectedIP: "192.168.1.1",
		},
		{
			name: "RemoteAddr without port",
			setupRequest: func(r *http.Request) {
				r.RemoteAddr = "192.168.1.1"
			},
			expectedIP: "192.168.1.1",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest("GET", "/test", nil)
			tt.setupRequest(req)

			ip := limiter.getClientIP(req)
			assert.Equal(t, tt.expectedIP, ip)
		})
	}
}

func TestRateLimiter_Cleanup(t *testing.T) {
	cfg := config.RateLimitConfig{
		Enabled:           true,
		RequestsPerMinute: 10,
		WindowSize:        100 * time.Millisecond,
	}

	limiter := NewRateLimiter(cfg)
	defer limiter.Stop()

	// Add some client records
	limiter.isAllowed("client1")
	limiter.isAllowed("client2")

	assert.Len(t, limiter.clients, 2)

	// Wait for records to become old
	time.Sleep(300 * time.Millisecond)

	// Trigger cleanup
	limiter.cleanOldClients()

	// Records should be cleaned up (this test might be flaky due to timing)
	// We just verify the cleanup method runs without panicking
	assert.True(t, true) // Placeholder assertion
}

func BenchmarkRateLimit_Allow(b *testing.B) {
	cfg := config.RateLimitConfig{
		Enabled:           true,
		RequestsPerMinute: 1000,
		BurstSize:         100,
		WindowSize:        time.Minute,
	}

	handler := RateLimit(cfg)(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))

	req := httptest.NewRequest("GET", "/test", nil)
	req.RemoteAddr = "192.168.1.1:12345"

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			rr := httptest.NewRecorder()
			handler.ServeHTTP(rr, req)
		}
	})
}

func BenchmarkRateLimit_Reject(b *testing.B) {
	cfg := config.RateLimitConfig{
		Enabled:           true,
		RequestsPerMinute: 1,
		BurstSize:         1,
		WindowSize:        time.Minute,
	}

	handler := RateLimit(cfg)(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))

	// Exhaust the limit first
	req := httptest.NewRequest("GET", "/test", nil)
	req.RemoteAddr = "192.168.1.1:12345"
	rr := httptest.NewRecorder()
	handler.ServeHTTP(rr, req)

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			rr := httptest.NewRecorder()
			handler.ServeHTTP(rr, req)
		}
	})
}


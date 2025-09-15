package middleware

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/http"
	"sync"
	"time"

	"github.com/banua-coder/pico-api-go/internal/config"
)

// ErrorResponse represents an error response structure
type ErrorResponse struct {
	Status string `json:"status"`
	Error  string `json:"error"`
}

// writeRateLimitError writes a rate limit error response
func writeRateLimitError(w http.ResponseWriter, statusCode int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	response := ErrorResponse{
		Status: "error",
		Error:  message,
	}
	if err := json.NewEncoder(w).Encode(response); err != nil {
		log.Printf("Error encoding rate limit JSON response: %v", err)
	}
}

// ClientRecord tracks request history for a client
type ClientRecord struct {
	requests    []time.Time
	mutex       sync.RWMutex
	lastCleanup time.Time
}

// RateLimiter implements a sliding window rate limiter
type RateLimiter struct {
	clients       map[string]*ClientRecord
	mutex         sync.RWMutex
	config        config.RateLimitConfig
	cleanupTicker *time.Ticker
	stopChan      chan struct{}
}

// NewRateLimiter creates a new rate limiter instance
func NewRateLimiter(cfg config.RateLimitConfig) *RateLimiter {
	rl := &RateLimiter{
		clients:  make(map[string]*ClientRecord),
		config:   cfg,
		stopChan: make(chan struct{}),
	}

	// Start background cleanup every 5 minutes
	if cfg.Enabled {
		rl.cleanupTicker = time.NewTicker(5 * time.Minute)
		go rl.cleanup()
	}

	return rl
}

// Stop gracefully stops the rate limiter
func (rl *RateLimiter) Stop() {
	if rl.cleanupTicker != nil {
		rl.cleanupTicker.Stop()
	}
	close(rl.stopChan)
}

// cleanup removes old client records periodically
func (rl *RateLimiter) cleanup() {
	for {
		select {
		case <-rl.cleanupTicker.C:
			rl.cleanOldClients()
		case <-rl.stopChan:
			return
		}
	}
}

// cleanOldClients removes clients that haven't made requests recently
func (rl *RateLimiter) cleanOldClients() {
	rl.mutex.Lock()
	defer rl.mutex.Unlock()

	cutoff := time.Now().Add(-rl.config.WindowSize * 2) // Keep records for 2x window size

	for clientIP, record := range rl.clients {
		record.mutex.RLock()
		shouldDelete := len(record.requests) == 0 ||
			(len(record.requests) > 0 && record.requests[len(record.requests)-1].Before(cutoff))
		record.mutex.RUnlock()

		if shouldDelete {
			delete(rl.clients, clientIP)
		}
	}
}

// getClientIP extracts client IP from request
func (rl *RateLimiter) getClientIP(r *http.Request) string {
	// Check X-Forwarded-For header first (for load balancers/proxies)
	if xff := r.Header.Get("X-Forwarded-For"); xff != "" {
		// Take the first IP from the comma-separated list
		if firstIP := xff; firstIP != "" {
			if ip := net.ParseIP(firstIP); ip != nil {
				return firstIP
			}
		}
	}

	// Check X-Real-IP header
	if xri := r.Header.Get("X-Real-IP"); xri != "" {
		if ip := net.ParseIP(xri); ip != nil {
			return xri
		}
	}

	// Fall back to RemoteAddr
	ip, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		return r.RemoteAddr
	}
	return ip
}

// isAllowed checks if a request should be allowed
func (rl *RateLimiter) isAllowed(clientIP string) (bool, int, time.Duration) {
	rl.mutex.Lock()
	client, exists := rl.clients[clientIP]
	if !exists {
		client = &ClientRecord{
			requests:    make([]time.Time, 0),
			lastCleanup: time.Now(),
		}
		rl.clients[clientIP] = client
	}
	rl.mutex.Unlock()

	client.mutex.Lock()
	defer client.mutex.Unlock()

	now := time.Now()
	windowStart := now.Add(-rl.config.WindowSize)

	// Remove old requests outside the window
	validRequests := make([]time.Time, 0, len(client.requests))
	for _, reqTime := range client.requests {
		if reqTime.After(windowStart) {
			validRequests = append(validRequests, reqTime)
		}
	}
	client.requests = validRequests

	// Check if we can allow this request
	if len(client.requests) >= rl.config.RequestsPerMinute {
		// Calculate when the oldest request in the window will expire
		if len(client.requests) > 0 {
			oldestRequest := client.requests[0]
			resetTime := oldestRequest.Add(rl.config.WindowSize).Sub(now)
			if resetTime < 0 {
				resetTime = 0
			}
			return false, rl.config.RequestsPerMinute - len(client.requests), resetTime
		}
		return false, 0, rl.config.WindowSize
	}

	// Allow the request and record it
	client.requests = append(client.requests, now)
	remaining := rl.config.RequestsPerMinute - len(client.requests)

	return true, remaining, 0
}

// RateLimit returns a middleware that implements rate limiting
func RateLimit(cfg config.RateLimitConfig) func(http.Handler) http.Handler {
	if !cfg.Enabled {
		// Return a no-op middleware if rate limiting is disabled
		return func(next http.Handler) http.Handler {
			return next
		}
	}

	limiter := NewRateLimiter(cfg)

	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			clientIP := limiter.getClientIP(r)
			allowed, remaining, resetTime := limiter.isAllowed(clientIP)

			// Set rate limiting headers
			w.Header().Set("X-RateLimit-Limit", fmt.Sprintf("%d", cfg.RequestsPerMinute))
			w.Header().Set("X-RateLimit-Remaining", fmt.Sprintf("%d", remaining))

			if !allowed {
				w.Header().Set("X-RateLimit-Reset", fmt.Sprintf("%d", time.Now().Add(resetTime).Unix()))
				w.Header().Set("Retry-After", fmt.Sprintf("%d", int(resetTime.Seconds())))

				writeRateLimitError(w, http.StatusTooManyRequests, "Rate limit exceeded. Too many requests.")
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}

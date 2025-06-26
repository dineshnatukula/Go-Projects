package ratelimiter

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

// RateLimiter implements the token bucket algorithm
type RateLimiter struct {
	rate      int       // tokens per second
	burst     int       // maximum bucket size
	tokens    float64   // current number of tokens
	lastCheck time.Time // last token refill time
	mu        sync.Mutex
}

// NewRateLimiter creates a new RateLimiter
func NewRateLimiter(rate, burst int) *RateLimiter {
	return &RateLimiter{
		rate:      rate,
		burst:     burst,
		tokens:    float64(burst),
		lastCheck: time.Now(),
	}
}

// Allow checks if a request can proceed
func (rl *RateLimiter) Allow() bool {
	rl.mu.Lock()
	defer rl.mu.Unlock()

	now := time.Now()
	elapsed := now.Sub(rl.lastCheck).Seconds()
	fmt.Printf("\nElapsed time is: %f tokens are %f", elapsed, rl.tokens)
	rl.lastCheck = now

	// Refill tokens based on elapsed time
	rl.tokens += elapsed * float64(rl.rate)
	if rl.tokens > float64(rl.burst) {
		rl.tokens = float64(rl.burst)
	}

	if rl.tokens >= 1 {
		rl.tokens--
		return true
	}

	return false
}

// LimiterStore manages rate limiters per key (e.g., IP or API key)
type LimiterStore struct {
	mu       sync.Mutex
	limiters map[string]*RateLimiter
	rate     int
	burst    int
}

// NewLimiterStore initializes a LimiterStore
func NewLimiterStore(rate, burst int) *LimiterStore {
	return &LimiterStore{
		limiters: make(map[string]*RateLimiter),
		rate:     rate,
		burst:    burst,
	}
}

// GetLimiter retrieves or creates a limiter for the given key
func (ls *LimiterStore) GetLimiter(key string) *RateLimiter {
	ls.mu.Lock()
	defer ls.mu.Unlock()

	limiter, exists := ls.limiters[key]
	if !exists {
		limiter = NewRateLimiter(ls.rate, ls.burst)
		ls.limiters[key] = limiter
	}
	return limiter
}

// RateLimitMiddleware returns a middleware that applies rate limiting
func RateLimitMiddleware(store *LimiterStore) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			key := r.RemoteAddr // or extract API key / user ID from headers

			limiter := store.GetLimiter(key)
			if !limiter.Allow() {
				http.Error(w, "Rate limit exceeded", http.StatusTooManyRequests)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}

package ratelimiter

import (
	"testing"
	"time"
)

func TestRateLimiter_Allow(t *testing.T) {
	rl := NewRateLimiter(5, 10) // 5 tokens/sec, burst 10

	// Should allow up to 10 initial tokens (burst)
	for i := 0; i < 10; i++ {
		if !rl.Allow() {
			t.Errorf("expected Allow to return true, got false at i=%d", i)
		}
	}

	// Now it should block (tokens depleted)
	if rl.Allow() {
		t.Errorf("expected Allow to return false after burst limit reached")
	}

	// Wait for refill
	time.Sleep(250 * time.Millisecond) // ~1.25 tokens refill
	if !rl.Allow() {
		t.Errorf("expected Allow to return true after refill")
	}
}

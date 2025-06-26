package ratelimiter_test

import (
	"net/http"
	"net/http/httptest"
	"ratelimiter/internal/ratelimiter"
	"testing"
)

func TestRateLimitMiddleware(t *testing.T) {
	store := ratelimiter.NewLimiterStore(2, 2) // 2 RPS, burst 2

	handler := ratelimiter.RateLimitMiddleware(store)(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	}))

	server := httptest.NewServer(handler)
	defer server.Close()

	client := &http.Client{}

	// 2 requests should succeed
	for i := 0; i < 2; i++ {
		resp, err := client.Get(server.URL)
		if err != nil || resp.StatusCode != http.StatusOK {
			t.Fatalf("expected 200 OK, got %v", resp.StatusCode)
		}
	}

	// 3rd should be rate limited
	resp, err := client.Get(server.URL)
	if err != nil {
		t.Fatalf("request failed: %v", err)
	}

	if resp.StatusCode != http.StatusTooManyRequests {
		t.Errorf("expected 429 Too Many Requests, got %v", resp.StatusCode)
	}
}

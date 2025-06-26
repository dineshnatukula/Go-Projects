package ratelimiter

import (
	"testing"
)

func TestLimiterStore_GetLimiter(t *testing.T) {
	store := NewLimiterStore(1, 2)

	l1 := store.GetLimiter("client1")
	l2 := store.GetLimiter("client2")
	l1Again := store.GetLimiter("client1")

	if l1 == nil || l2 == nil || l1Again == nil {
		t.Fatal("expected non-nil limiters")
	}

	if l1 != l1Again {
		t.Error("expected same limiter for same key")
	}

	if l1 == l2 {
		t.Error("expected different limiters for different keys")
	}
}

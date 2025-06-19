package backend

import (
	"math/rand"
	"sync"
	"testing"
	"time"
)

// Number of operations per goroutine
// const ops = 1000
// const goroutines = 100

// Benchmark using map + RWMutex
func BenchmarkMapRWMutex(b *testing.B) {
	sm := struct {
		mu sync.RWMutex
		m  map[int]int
	}{
		m: make(map[int]int),
	}

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		for pb.Next() {
			key := r.Intn(100)
			if r.Intn(2) == 0 {
				sm.mu.RLock()
				_ = sm.m[key]
				sm.mu.RUnlock()
			} else {
				sm.mu.Lock()
				sm.m[key] = key
				sm.mu.Unlock()
			}
		}
	})
}

// Benchmark using sync.Map
func BenchmarkSyncMap(b *testing.B) {
	var sm sync.Map

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		for pb.Next() {
			key := r.Intn(100)
			if r.Intn(2) == 0 {
				sm.Load(key)
			} else {
				sm.Store(key, key)
			}
		}
	})
}

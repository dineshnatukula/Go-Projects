package backend

import (
	"fmt"
	"sync"
)

// SafeCounter is safe to use concurrently.
type SafeCounter struct {
	mu sync.Mutex
	v  map[int]int
}

// Inc increments the counter for the given key.
func (c *SafeCounter) Inc(key int) {
	c.mu.Lock()
	// Lock so only one goroutine at a time can access the map c.v.
	c.v[key]++
	c.mu.Unlock()
}

// Value returns the current value of the counter for the given key.
func (c *SafeCounter) Value(key int) int {
	c.mu.Lock()
	// Lock so only one goroutine at a time can access the map c.v.
	defer c.mu.Unlock()
	return c.v[key]
}

func SyncMutex() {
	c := SafeCounter{v: make(map[int]int)}
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		for i := 0; i < 1000; i++ {
			go c.Inc(i)
		}
		wg.Done()
	}()
	wg.Add(1)
	go func() {
		for i := 0; i < 1000; i++ {
			fmt.Println(i, c.Value(i))
		}
		wg.Done()
	}()
	wg.Wait()
}

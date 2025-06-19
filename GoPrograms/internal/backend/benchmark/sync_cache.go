package backend

import (
	"fmt"
	"sync"
)

type Cache struct {
	data map[string]string
	mu   sync.RWMutex
}

func NewCache() *Cache {
	return &Cache{
		data: make(map[string]string),
	}
}

// Get attempts to retrieve a value from the cache (thread-safe)
func (c *Cache) Get(key string) (string, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	val, found := c.data[key]
	return val, found
}

// Set adds a value to the cache (thread-safe)
func (c *Cache) Set(key, value string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.data[key] = value
}

type MarketSelection struct {
	EventID     string
	MarketID    string
	SelectionID string
}

func SyncMapWithIndvKeys() {
	cache := NewCache()

	keys := []string{"U-123", "S-121", "S-120", "S-124", "U-125", "S-126", "S-127", "U-128", "U-129"}

	// Add initial values to the cache
	cache.Set("U-123", "Ind Vs Pak")
	cache.Set("S-121", "Ind Wins")
	cache.Set("S-120", "Pak Wins")
	cache.Set("S-124", "May be Draw")
	cache.Set("U-125", "Pak Vs Aus")
	cache.Set("S-126", "Pak Wins")
	cache.Set("S-127", "Aus Wins")
	cache.Set("U-128", "Draw for Pak Vs Aus")
	cache.Set("U-129", "Aus Vs SL")

	for _, key := range keys {
		if val, found := cache.Get(key); found {
			fmt.Printf("Cache HIT for key '%s': %s\n", key, val)
		} else {
			fmt.Printf("Cache MISS for key '%s'. Fetching and adding to cache...\n", key)
			fetched := "Fetched description for " + key
			cache.Set(key, fetched)
		}
	}
}

func SyncMapWithCombKeys() {
	cache := NewCache()

	keys := []string{"U-123:S-121", "U-123:S-120", "U-123:S-124", "U-125:S-126", "U-125:S-127", "U-125:U-128", "U-125:U-129"}
	cache.Set("U-123:S-121", "Ind Vs Pak:Ind Wins")
	cache.Set("U-123:S-120", "Ind Vs Pak:Pak Wins")
	cache.Set("U-123:S-124", "Ind Vs Pak:May be Draw")
	cache.Set("U-125:S-126", "Pak Vs Aus:Pak Wins")
	cache.Set("U-125:S-127", "Pak Vs Aus:Aus Wins")
	cache.Set("U-125:U-128", "Pak Vs Aus:May be Draw")

	// Add initial values to the cache

	for _, key := range keys {
		if val, found := cache.Get(key); found {
			fmt.Printf("Cache HIT for key '%s': %s\n", key, val)
		} else {
			fmt.Printf("Cache MISS for key '%s'. Fetching and adding to cache...\n", key)
			fetched := "Fetched description for " + key
			cache.Set(key, fetched)
		}
	}
}

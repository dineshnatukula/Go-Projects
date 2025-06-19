package backend

import (
	"fmt"
	"sync"
)

func SyncMap() {
	var m sync.Map

	// Store
	m.Store("name", "Alice")
	m.Store("age", 30)

	// Load
	if val, ok := m.Load("name"); ok {
		fmt.Println("Name:", val)
	}

	// LoadOrStore
	actual, loaded := m.LoadOrStore("city", "Paris")
	fmt.Println("City:", actual, "Already present?", loaded)

	// Delete
	m.Delete("age")

	// Range (iterating over map)
	m.Range(func(key, value any) bool {
		fmt.Printf("Key: %v, Value: %v\n", key, value)
		return true // continue iteration
	})
}

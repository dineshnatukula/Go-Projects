package backend

import (
	"fmt"
	"sync"
)

func SyncMapDemo() {
	var m sync.Map

	m.Store("Name", "Diensh")
	m.Store("Age", 30)
	m.Store(10, 20)

	if val, ok := m.Load("Name"); ok {
		fmt.Println("Name is :", val)
	}

	if val, ok := m.LoadOrStore("Salary", "40 LPA"); !ok {
		fmt.Println("Salary not found and stored in cache", val)
	}

	m.Range(func(key, value any) bool {
		fmt.Println("Key is", key, "Value is", value)
		return true
	})

}

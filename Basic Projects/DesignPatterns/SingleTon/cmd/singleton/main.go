package main

import (
	"fmt"
	"sync"
)

var (
	once     sync.Once
	instance *Config // Move this to package level
)

type Config struct {
	Tolerance float64 `json:"tolerance,omitempty"`
}

func GetConfig() *Config {
	once.Do(func() {
		fmt.Println("First Time Calling...")
		instance = &Config{Tolerance: 5.01}
	})
	return instance
}

func main() {
	instance1 := GetConfig()
	fmt.Println(instance1.Tolerance)

	instance2 := GetConfig()
	fmt.Println(instance2.Tolerance)
}

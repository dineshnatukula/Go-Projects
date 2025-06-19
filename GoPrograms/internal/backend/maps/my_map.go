package backend

import (
	"fmt"
	"runtime"
	"time"
)

func MyMap() {
	type SafeCounter struct {
		v map[int]int
	}

	sc := SafeCounter{
		v: make(map[int]int),
	}
	fmt.Println("After init")
	printMemStats()
	fmt.Println("-----------------------------")

	for i := range 1000 {
		// fmt.Printf("Size of Map %d and Len of Map %d \n", unsafe.Size(sc.v), len(sc.v))
		sc.v[i] = i
	}
	fmt.Println("After insertions")
	printMemStats()
	fmt.Println("-----------------------------")

	for i := range 800 {
		delete(sc.v, i)
	}

	fmt.Println("After Deletions")
	printMemStats()
	fmt.Println("-----------------------------")

	go func() {
		for {
			fmt.Println("After GC")
			runtime.GC()
			printMemStats()
			fmt.Println("-----------------------------")
			time.Sleep(5 * time.Second)
		}
	}()

	select {}
}

func printMemStats() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)

	fmt.Printf("Alloc = %v MiB\n", m.Alloc)           // Memory allocated by the program
	fmt.Printf("TotalAlloc = %v MiB\n", m.TotalAlloc) // Total memory allocated (including freed memory)
	fmt.Printf("Sys = %v MiB\n", m.Sys)               // Total memory obtained from the OS
	fmt.Printf("NumGC = %v\n", m.NumGC)               // Number of GC cycles
}

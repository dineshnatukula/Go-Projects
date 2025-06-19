package backend

import (
	"fmt"
	"sync"
)

func GoRoutine() {
	ch := make(chan int, 2)
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 1; i <= 100; i++ {
			ch <- i
		}
		close(ch)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for value := range ch {
			fmt.Println(value)
		}
	}()
	wg.Wait()
}

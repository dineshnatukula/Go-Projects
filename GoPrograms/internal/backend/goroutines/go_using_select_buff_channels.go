package backend

import (
	"fmt"
)

func put(ch chan int, done chan int) {
	var i int
	for {
		select {
		case ch <- i:
			fmt.Println("Channel executed")
			i++
		case <-done:
			// close(ch)
			fmt.Println("Done executed...................")
			return
		}
	}
}

func GoRoutineWithSelect() {
	ch := make(chan int)
	done := make(chan int)
	// done1 := make(chan int)
	go func() {
		put(ch, done)
	}()
	go func() {
		// for val := range ch {
		for i := 1; i <= 10; i++ {
			<-ch
			// fmt.Printf("Value of i is : %d\n", val)
		}
		done <- 0
		// done1 <- 0
	}()

	fmt.Println("Executing done from main..")
	<-done
	fmt.Println("Executed done from main..")
}

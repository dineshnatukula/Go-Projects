package context

import (
	"context"
	"fmt"
	"time"
)

func ContextDoneMultiple() {

	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()

	go func() {
		<-ctx.Done()
		fmt.Println("Goroutine 1: context done")
	}()

	go func() {
		<-ctx.Done()
		fmt.Println("Goroutine 2: context done")
	}()

	<-ctx.Done()
	fmt.Println("Main: context done")

}

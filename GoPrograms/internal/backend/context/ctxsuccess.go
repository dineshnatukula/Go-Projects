package context

import (
	"context"
	"fmt"
	"time"
)

func slowTaskSuccess(ctx context.Context, done chan int) {
	for i := 0; i < 5; i++ {
		select {
		case <-ctx.Done(): // If the context is canceled, stop the task
			fmt.Println("Task canceled")
			return
		default:
			// Simulate a task taking time (like a network request or heavy computation)
			time.Sleep(1000 * time.Millisecond)
			fmt.Printf("Doing task %d...\n", i)
		}
	}
	close(done) // Signal that the task completed successfully
}

func ContextSuccess() {
	// Create a context with a timeout
	ctx, cancel := context.WithTimeout(context.Background(), 4*time.Second)
	defer cancel() // Ensure the cancel function is called to release resources

	// Run slowTask in a goroutine
	done := make(chan int)
	go slowTaskSuccess(ctx, done)

	select {
	case <-ctx.Done():
		fmt.Println("Operation timed out!")
	case <-done:
		fmt.Println("Task completed successfully.")
	}
}

package context

import (
	"context"
	"fmt"
	"time"
)

func slowTask(ctx context.Context) {
	for i := 0; i < 5; i++ {
		select {
		case <-ctx.Done(): // If the context is canceled, stop the task
			fmt.Println("Task canceled")
			return
		default:
			// Simulate a task taking time (like a network request or heavy computation)
			time.Sleep(300 * time.Millisecond)
			fmt.Printf("Doing task %d...\n", i)
		}
	}
}

func Context1() {
	// Create a context with a timeout
	ctx, cancel := context.WithTimeout(context.Background(), 4*time.Second)
	defer cancel() // Ensure the cancel function is called to release resources

	// Run slowTask in a goroutine
	go slowTask(ctx)

	// Wait for the context to be done (i.e., the timeout)
	<-ctx.Done()

	// Check if the context was canceled due to timeout
	if ctx.Err() == context.DeadlineExceeded {
		fmt.Println("Operation timed out!")
	} else {
		fmt.Println("Task completed successfully.")
	}
}

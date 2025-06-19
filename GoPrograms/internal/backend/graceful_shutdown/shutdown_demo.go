package gracefulshutdown

import (
	"context"
	"fmt"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func GracefulShutdownDemo() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	var wg sync.WaitGroup

	// Start a long-running worker
	wg.Add(1)
	go func() {
		defer wg.Done()
		worker(ctx)
	}()

	// Wait for interrupt signal
	<-ctx.Done()
	fmt.Println("\nMain: received shutdown signal")

	// Optional timeout to force shutdown
	shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Wait for all goroutines to finish or timeout
	done := make(chan struct{})
	go func() {
		wg.Wait()
		close(done)
	}()

	select {
	case <-done:
		fmt.Println("Main: all workers shut down gracefully")
	case <-shutdownCtx.Done():
		fmt.Println("Main: shutdown timeout reached, forcing exit")
	}
}

func worker(ctx context.Context) {
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			fmt.Println("Worker: context cancelled, cleaning up...")
			// Do any necessary cleanup here
			time.Sleep(1 * time.Second) // simulate cleanup
			fmt.Println("Worker: shutdown complete")
			return
		case t := <-ticker.C:
			fmt.Println("Worker: doing work at", t)
		}
	}
}

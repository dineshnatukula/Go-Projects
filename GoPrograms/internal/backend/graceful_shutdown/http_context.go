package gracefulshutdown

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

type Job struct {
	ID int
}

var (
	jobQueue    = make(chan Job, 100)
	jobCounter  = 0
	jobCounterM sync.Mutex
)

func httpWorker(ctx context.Context, id int, jobs <-chan Job, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			log.Printf("Worker %d: shutting down", id)
			return
		case job, ok := <-jobs:
			if !ok {
				log.Printf("Worker %d: job queue closed", id)
				return
			}
			log.Printf("Worker %d: processing job %d", id, job.ID)
			time.Sleep(2 * time.Second) // simulate work
			log.Printf("Worker %d: finished job %d", id, job.ID)
		}
	}
}

func enqueueHandler(w http.ResponseWriter, r *http.Request) {
	jobCounterM.Lock()
	jobCounter++
	id := jobCounter
	jobCounterM.Unlock()

	select {
	case jobQueue <- Job{ID: id}:
		fmt.Fprintf(w, "Enqueued job %d\n", id)
	default:
		http.Error(w, "Job queue full", http.StatusServiceUnavailable)
	}
}

func HttpWorkerDemo() {
	const numWorkers = 4

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	// Start workers
	var wg sync.WaitGroup
	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go httpWorker(ctx, i, jobQueue, &wg)
	}

	// HTTP server setup
	mux := http.NewServeMux()
	mux.HandleFunc("/enqueue", enqueueHandler)

	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	// Start server in background
	go func() {
		log.Println("HTTP server listening on :8080")
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server error: %v", err)
		}
	}()

	// Wait for shutdown signal
	<-ctx.Done()
	log.Println("Shutdown signal received")

	// Start graceful shutdown
	shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Stop accepting new connections
	if err := server.Shutdown(shutdownCtx); err != nil {
		log.Printf("HTTP server shutdown error: %v", err)
	}

	// Close job queue so workers can exit
	close(jobQueue)

	// Wait for workers
	done := make(chan struct{})
	go func() {
		wg.Wait()
		close(done)
	}()

	select {
	case <-done:
		log.Println("All workers shut down cleanly")
	case <-shutdownCtx.Done():
		log.Println("Forcing shutdown: timeout reached")
	}

	log.Println("Server exited")
}

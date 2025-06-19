package wp

import (
	"context"
	"fmt"
	"sync"
	"time"
)

type Job struct {
	ID     int
	Number int
}

type Result struct {
	JobID  int
	Output int
}

func startWorker(wpContext context.Context, id int, jobs <-chan Job, results chan<- Result, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-wpContext.Done():
			// fmt.Println("Worker Pool Context timeout...")
			fmt.Printf("Worker %d: context cancelled, stopping\n", id)
			return
		case job, ok := <-jobs:
			if !ok {
				return
			}
			// Simulate work
			time.Sleep(500 * time.Millisecond)
			fmt.Printf("Worker %d processing job %d\n", id, job.ID)
			select {
			case <-wpContext.Done():
				return
			case results <- Result{JobID: job.ID, Output: job.Number * job.Number}:
			}
		}
	}
}

func WPDemo() {
	const noOfJosb = 10
	const noOfWorkers = 3

	wpContext, stop := context.WithTimeout(context.Background(), 1*time.Second)
	defer stop()

	jobs := make(chan Job, noOfJosb)
	results := make(chan Result, noOfWorkers)

	var wg sync.WaitGroup

	// Start Workers
	for w := 1; w <= noOfWorkers; w++ {
		wg.Add(1)
		go startWorker(wpContext, w, jobs, results, &wg)
	}

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	for i := range noOfJosb {
		select {
		case <-wpContext.Done():
			fmt.Println("Worker Pool Context timeout...")
			break
		case jobs <- Job{ID: i, Number: i}:
		}
	}
	close(jobs)

	go func() {
		wg.Wait()
		close(results)
	}()

	for res := range results {
		fmt.Println("Job with ID", res.JobID, "and its result:", res.Output)
	}

	<-shutdownCtx.Done()
}

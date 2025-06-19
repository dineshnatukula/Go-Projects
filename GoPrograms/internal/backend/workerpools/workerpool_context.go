package workerpools

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// type Job struct {
// 	ID     int
// 	Number int
// }

// type Result struct {
// 	JobID  int
// 	Output int
// }

func workerCTXFN(ctx context.Context, id int, jobs <-chan Job, results chan<- Result, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("Worker %d: context cancelled, stopping\n", id)
			return
		case job, ok := <-jobs:
			if !ok {
				// channel closed, no more jobs
				return
			}
			// simulate work
			time.Sleep(100 * time.Millisecond)
			fmt.Printf("Worker %d processing job %d\n", id, job.ID)
			select {
			case <-ctx.Done():
				return
			case results <- Result{JobID: job.ID, Output: job.Number * job.Number}:
			}
		}
	}
}

func WorkerPoolWithContext() {
	const numJobs = 10
	const numWorkers = 3

	jobs := make(chan Job, numJobs)
	results := make(chan Result, numJobs)

	workerCTX, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()

	var wgCTX sync.WaitGroup

	for w := 1; w <= numWorkers; w++ {
		wgCTX.Add(1)
		go workerCTXFN(workerCTX, w, jobs, results, &wgCTX)
	}

	// Send jobs
	for j := 1; j <= numJobs; j++ {
		select {
		case <-workerCTX.Done():
			fmt.Println("Main: context done before sending all jobs")
			break
		case jobs <- Job{ID: j, Number: j}:
		}
	}
	close(jobs)

	// Wait for workers in a separate goroutine
	go func() {
		wgCTX.Wait()
		close(results)
	}()

	// Collect results
	for result := range results {
		fmt.Printf("Result for job %d = %d\n", result.JobID, result.Output)
	}

	fmt.Println("Main: finished")
}

package workerpools

import (
	"fmt"
	"sync"
	"time"
)

// Job represents a unit of work
type Job struct {
	ID     int
	Number int
}

// Result represents the result of a job
type Result struct {
	JobID  int
	Output int
}

// Worker function
func worker(id int, jobs <-chan Job, results chan<- Result, wg *sync.WaitGroup) {
	defer wg.Done()
	for job := range jobs {
		// Simulate work
		time.Sleep(100 * time.Millisecond)
		fmt.Printf("Worker %d processing job %d\n", id, job.ID)
		results <- Result{JobID: job.ID, Output: job.Number * job.Number}
	}
}

func WorkerPoolDemo() {
	const numJobs = 10
	const numWorkers = 3

	jobs := make(chan Job, numJobs)
	results := make(chan Result, numJobs)

	var wg sync.WaitGroup

	// Start workers
	for w := 1; w <= numWorkers; w++ {
		wg.Add(1)
		go worker(w, jobs, results, &wg)
	}

	// Send jobs
	for j := 1; j <= numJobs; j++ {
		jobs <- Job{ID: j, Number: j}
	}
	close(jobs)

	// Wait for all workers to finish
	go func() {
		wg.Wait()
		close(results)
	}()

	// Read results
	for result := range results {
		fmt.Printf("Result for job %d = %d\n", result.JobID, result.Output)
	}

	select {}
}

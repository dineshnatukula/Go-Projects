package workerpools

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"
)

type DeepPool struct {
	NumWorkers int
	Tasks      chan *int

	// ensure the pool can only be started once
	Start sync.Once
	// ensure the pool can only be stopped once
	Stop sync.Once

	// close to signal the workers to stop working
	quit chan struct{}

	// This is for waiting till all the goroutine is done with their work
	stopWait sync.WaitGroup
	Closed   bool
	mutex    sync.Mutex
}

func (dp *DeepPool) StartWorkers() {
	for i := range dp.NumWorkers {
		// Adding one go routine to tbe WorkerPool waitgroup
		dp.stopWait.Add(1)

		// Creating a go routine
		go func(workerNum int) {
			fmt.Println("Starting worker with workerNum", workerNum)
			defer func() {
				dp.stopWait.Done()
			}()

			for {
				select {
				case <-dp.quit:
					fmt.Println("Stopping Worker with workerNum", workerNum)
					return
				case t, ok := <-dp.Tasks:
					if !ok {
						fmt.Println("Stopping worker with workerNum", workerNum)
						return
					}

					if *t >= 0 {
						// time.Sleep(500 * time.Millisecond)
						fmt.Println("Task Performed...", (*t)*(*t))
					}
				}
			}
		}(i)
	}
}

func (dp *DeepPool) StartDeepPool() {
	dp.Start.Do(func() {
		dp.StartWorkers()
	})
}

func (dp *DeepPool) StopDeepPool() {
	dp.Stop.Do(func() {
		dp.mutex.Lock()
		defer dp.mutex.Unlock()
		dp.Closed = true
		close(dp.quit)
		dp.stopWait.Wait()
	})
}

func (dp *DeepPool) PerformTasks(n int) {
	for i := range n {
		dp.AddWorkNonBlocking(i)
	}
}

func (dp *DeepPool) AddTask(n int) {
	dp.mutex.Lock()
	defer dp.mutex.Unlock()

	if dp.Closed {
		return
	}

	dp.Tasks <- &n
	fmt.Println("Task added to the Worker pool....", n)
}

func (dp *DeepPool) AddWorkNonBlocking(n int) {
	go dp.AddTask(n)
}

// Initiliasing the DeepPool after validations.
func NewDeepPool(numWorkers int, channelSize int) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel() // Ensure the cancel function is called to release resources

	if numWorkers <= 0 {
		return errors.New("numberOfWorkers should be > 0")
	}
	if channelSize < 10 {
		return errors.New("channelSize should be > 0")
	}

	dp := &DeepPool{
		Tasks:      make(chan *int, channelSize),
		NumWorkers: channelSize,
		quit:       make(chan struct{}),
		Closed:     false,
		Start:      sync.Once{},
		Stop:       sync.Once{},
	}

	dp.StartDeepPool()
	dp.PerformTasks(channelSize)

	select {
	case <-ctx.Done():
		fmt.Println("Tasks incomplete, context timeout...")
	default:
		time.Sleep(5 * time.Second)
		dp.StopDeepPool()
		// fmt.Println("Tasks are running...")
	}

	return nil
}

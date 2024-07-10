package pkg

import (
	"context"
	"sync"
)

type Worker func(job interface{})

type Scheduler struct {
	jobs    chan interface{}
	workers int
	wg      sync.WaitGroup
}

func NewScheduler(workers int) *Scheduler {
	return &Scheduler{
		jobs:    make(chan interface{}),
		workers: workers,
	}
}

// Start initializes the worker pool and starts processing jobs.
func (s *Scheduler) Start(ctx context.Context, worker Worker) {
	for i := 0; i < s.workers; i++ {
		s.wg.Add(1)
		go func() {
			defer s.wg.Done()
			for {
				select {
				case job, ok := <-s.jobs:
					if !ok {
						return // Channel closed, stop the worker.
					}
					worker(job)
				case <-ctx.Done():
					return // Context cancelled, stop the worker.
				}
			}
		}()
	}
}

// Submit adds a job to the queue to be processed by the workers.
func (s *Scheduler) Submit(job interface{}) {
	s.jobs <- job
}

// Stop waits for all workers to finish processing and closes the job channel.
func (s *Scheduler) Stop() {
	close(s.jobs) // Close the jobs channel to signal workers to stop.
	s.wg.Wait()   // Wait for all workers to finish.
}

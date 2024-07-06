package main

import "log"

// Worker executes jobs from the job queue.
type Worker struct {
	ID         int
	JobQueue   *JobQueue
	WorkerPool chan chan *Job
	JobChannel chan *Job
	quit       chan bool
}

// NewWorker creates a new Worker instance.
func NewWorker(id int, jobQueue *JobQueue, workerPool chan chan *Job) *Worker {
	return &Worker{
		ID:         id,
		JobQueue:   jobQueue,
		WorkerPool: workerPool,
		JobChannel: make(chan *Job),
		quit:       make(chan bool),
	}
}

// Start begins the worker's job processing loop.
func (w *Worker) Start() {
	go func() {
		for {
			// Register the current worker into the worker queue.
			w.WorkerPool <- w.JobChannel

			select {
			case job := <-w.JobChannel:
				// Received a job to process.
				if err := job.Function(); err != nil {
					log.Printf("Error executing job %s: %s", job.ID, err)
				}
			case <-w.quit:
				// Stop the worker.
				return
			}
		}
	}()
}

// Stop signals the worker to stop listening for job requests.
func (w *Worker) Stop() {
	go func() {
		w.quit <- true
	}()
}

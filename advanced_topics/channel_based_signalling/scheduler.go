package main

// Scheduler coordinates job execution.
type Scheduler struct {
	jobQueue   *JobQueue
	workerPool chan chan *Job
	maxWorkers int
}

// NewScheduler creates a new Scheduler instance.
func NewScheduler(maxWorkers int) *Scheduler {
	jobQueue := NewJobQueue()
	workerPool := make(chan chan *Job, maxWorkers)
	return &Scheduler{
		jobQueue:   jobQueue,
		workerPool: workerPool,
		maxWorkers: maxWorkers,
	}
}

// Start initializes the scheduler and its workers.
func (s *Scheduler) Start() {
	// Start workers.
	for i := 0; i < s.maxWorkers; i++ {
		worker := NewWorker(i+1, s.jobQueue, s.workerPool)
		worker.Start()
	}

	// Start the job dispatcher.
	go s.dispatch()
}

// dispatch sends jobs to available workers.
func (s *Scheduler) dispatch() {
	for {
		job := s.jobQueue.Dequeue()
		if job == nil {
			continue // Skip if no job is dequeued.
		}

		go func(job *Job) {
			jobChannel := <-s.workerPool
			jobChannel <- job
		}(job)
	}
}

// ScheduleJob adds a job to the scheduler's job queue.
func (s *Scheduler) ScheduleJob(job *Job) {
	s.jobQueue.Enqueue(job)
}

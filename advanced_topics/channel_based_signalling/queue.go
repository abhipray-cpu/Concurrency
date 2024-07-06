package main

import "sync"

// JobQueue manages a queue of jobs.
type JobQueue struct {
	queue []*Job
	lock  sync.Mutex
}

// NewJobQueue creates a new JobQueue instance.
func NewJobQueue() *JobQueue {
	return &JobQueue{}
}

// Enqueue adds a job to the queue.
func (jq *JobQueue) Enqueue(job *Job) {
	jq.lock.Lock()
	defer jq.lock.Unlock()
	jq.queue = append(jq.queue, job)
}

// Dequeue removes and returns the first job in the queue.
func (jq *JobQueue) Dequeue() *Job {
	jq.lock.Lock()
	defer jq.lock.Unlock()
	if len(jq.queue) == 0 {
		return nil
	}
	job := jq.queue[0]
	jq.queue = jq.queue[1:]
	return job
}

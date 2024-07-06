package main

/*
Condition variables are used with mutexes to allow threads to wait for certain conditions to become true.
They are useful for scenarios where you need more fine-grained control over thread execution.
*/
import (
	"fmt"
	"sync"
	"time"
)

// A struct to hold the condition and the associated lock
type JobQueue struct {
	sync.Mutex
	jobsAvailable bool
	cond          *sync.Cond
}

// Worker function that waits for jobs to become available
func (jq *JobQueue) worker(id int) {
	jq.Lock()
	for !jq.jobsAvailable {
		fmt.Printf("Worker %d: Waiting for jobs...\n", id)
		jq.cond.Wait() // Wait for the condition (jobsAvailable == true)
	}
	fmt.Printf("Worker %d: Starting job\n", id)
	jq.Unlock()
}

// Function to add jobs to the queue and signal workers
func (jq *JobQueue) addJobs() {
	jq.Lock()
	fmt.Println("Adding jobs to the queue...")
	jq.jobsAvailable = true
	jq.cond.Broadcast() // Signal all waiting goroutines that the condition is met
	jq.Unlock()
}

func CondVariable() {
	jobQueue := JobQueue{}
	jobQueue.cond = sync.NewCond(&jobQueue.Mutex)

	// Start workers
	for i := 1; i <= 3; i++ {
		go jobQueue.worker(i)
	}

	// Simulate adding jobs after some time
	time.Sleep(2 * time.Second)
	jobQueue.addJobs()

	// Wait to observe output
	time.Sleep(1 * time.Second)
}

package main

import (
	"log"
	"strconv"
	"time"
)

func main() {
	scheduler := NewScheduler(5) // Initialize scheduler with 5 workers.
	scheduler.Start()

	// Example job function.
	jobFunc := func() error {
		log.Println("Executing job...")
		time.Sleep(2 * time.Second) // Simulate work.
		log.Println("Job completed.")
		return nil
	}

	// Schedule a few jobs.
	for i := 0; i < 10; i++ {
		job := NewJob(strconv.Itoa(i), jobFunc)
		scheduler.ScheduleJob(job)
	}

	// Prevent the main goroutine from exiting immediately.
	time.Sleep(10 * time.Second)
}

package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Task struct {
	ID int
}

// ExecuteWithRetry simulates task execution with retries on failure.
func (t *Task) ExecuteWithRetry(maxRetries int) error {
	for i := 0; i <= maxRetries; i++ {
		if err := t.execute(); err != nil {
			fmt.Printf("Task %d failed: %s. Retrying...\n", t.ID, err)
			time.Sleep(time.Second) // Simulate retry delay
			continue
		}
		return nil // Success
	}
	return fmt.Errorf("task %d exceeded max retries", t.ID)
}

// execute simulates a single attempt of task execution, randomly failing.
func (t *Task) execute() error {
	if rand.Intn(10) < 3 { // 30% chance of failure
		return fmt.Errorf("task %d execution failed", t.ID)
	}
	fmt.Printf("Task %d executed successfully\n", t.ID)
	return nil
}

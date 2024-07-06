package main

/*
Enforce a deterministic execution order of concurrent operations during testing to make the tests
reproducible.
This can involve using a deterministic scheduler or controlling the concurrency primitives
(like locks and semaphores) to enforce a specific execution order.
*/

import (
	"sync"
	"testing"
)

func TestDeterministicExecution(t *testing.T) {
	var wg sync.WaitGroup
	start := make(chan struct{})
	mutex1Unlocked := make(chan struct{})
	mutex2Unlocked := make(chan struct{})

	wg.Add(2)

	// Goroutine to lock mutexes in the defined order
	go func() {
		defer wg.Done()
		<-start // Wait for the signal to start
		lockMutexesInOrder()
	}()

	// Goroutine to lock mutexes in reverse order
	go func() {
		defer wg.Done()
		<-start // Wait for the signal to start
		lockMutexesInReverseOrder()
	}()

	close(start) // Signal both goroutines to start
	wg.Wait()    // Wait for both goroutines to finish

	// If the test reaches this point without deadlock, it passes
}

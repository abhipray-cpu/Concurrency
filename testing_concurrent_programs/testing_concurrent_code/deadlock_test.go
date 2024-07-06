/*
Use tools and techniques to detect deadlocks, where two or more operations wait on each other to release resources, causing a standstill.
Static analysis tools and runtime deadlock detectors can be employed.
*/
package main

import (
	"sync"
	"testing"
	"time"
)

func TestDeadlock(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		lockMutexes()
	}()

	go func() {
		defer wg.Done()
		reverseLockMutexes()
	}()

	done := make(chan struct{})
	go func() {
		wg.Wait()
		close(done)
	}()

	select {
	case <-done:
		// Success, no deadlock
	case <-time.After(1 * time.Second):
		t.Fatal("Test failed due to deadlock")
	}
}

package main

/*
Employ dynamic analysis tools to detect race conditions, where the outcome depends on the non-deterministic ordering of operations.
Tools like ThreadSanitizer (for C/C++ and Go) and Java's Race Catcher can help identify these issues.
*/

import (
	"sync"
	"testing"
)

func TestRaceCondition(t *testing.T) {
	counter = 0 // Reset counter
	var wg sync.WaitGroup
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			IncrementCounter()
			wg.Done()
		}()
	}
	wg.Wait()

	if counter != 1000 {
		t.Errorf("Expected counter to be 1000, got %d", counter)
	}
}

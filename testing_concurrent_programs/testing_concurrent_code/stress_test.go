package main

/*
Increase the load on the system to stress the concurrent aspects, such as creating a high number of threads or processes, to identify potential bottlenecks or race conditions that only occur under load.
Useful for identifying performance issues and ensuring that the system can handle peak loads.
*/

import (
	"fmt"
	"sync"
	"testing"
)

func TestDBStress(t *testing.T) {
	db := NewInMemoryDB()
	var wg sync.WaitGroup
	numOperations := 10000 // Number of operations to simulate

	for i := 0; i < numOperations; i++ {
		wg.Add(2) // One for read, one for write

		// Simulate write
		go func(i int) {
			key := fmt.Sprintf("key%d", i)
			value := fmt.Sprintf("value%d", i)
			db.Set(key, value)
			wg.Done()
		}(i)

		// Simulate read
		go func(i int) {
			key := fmt.Sprintf("key%d", i)
			if val, exists := db.Get(key); exists {
				// Optionally, verify the value
				expectedValue := fmt.Sprintf("value%d", i)
				if val != expectedValue {
					t.Errorf("Mismatched value for %s: got %s, want %s", key, val, expectedValue)
				}
			}
			wg.Done()
		}(i)
	}

	wg.Wait()

	// Verify database size
	if len(db.data) != numOperations {
		t.Errorf("Expected database size to be %d, got %d", numOperations, len(db.data))
	}
}

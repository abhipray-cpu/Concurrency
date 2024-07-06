package main

/*
Read-Write Locks (RWLocks) allow multiple readers to access a resource simultaneously but require exclusive access for writers.
This improves performance when reads are more frequent than writes.
*/

import (
	"fmt"
	"sync"
	"time"
)

// A shared resource
var counter int

// A Read-Write lock
var rwLock sync.RWMutex

// Simulates reading the counter value
func read(id int) {
	rwLock.RLock() // Acquire read lock
	fmt.Printf("Reader %d: %d\n", id, counter)
	rwLock.RUnlock() // Release read lock
}

// Simulates writing to the counter value
func write(id, value int) {
	rwLock.Lock() // Acquire write lock
	fmt.Printf("Writer %d: %d\n", id, value)
	counter = value
	rwLock.Unlock() // Release write lock
}

func ReadWriteLock() {
	var wg sync.WaitGroup

	// Start 5 readers
	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			read(id)
		}(i)
	}

	// Start a writer
	wg.Add(1)
	go func() {
		defer wg.Done()
		time.Sleep(1 * time.Second) // Delay to show readers reading the initial value
		write(1, 10)
	}()

	wg.Wait()
}

/*
A semaphore controls access to a common resource by multiple threads and is initialized with a maximum count.
It's a more flexible tool than mutexes for some scenarios.
*/

package main

import (
	"fmt"
	"sync"
	"time"
)

// SharedData structure with Read-Write lock and Condition variable
type SharedData struct {
	items []int
	lock  sync.RWMutex
	cond  *sync.Cond
	limit int
}

// NewSharedData initializes the SharedData structure
func NewSharedData(limit int) *SharedData {
	sd := &SharedData{limit: limit}
	sd.cond = sync.NewCond(&sd.lock)
	return sd
}

// AddItem tries to add an item if under limit, using a semaphore-like mechanism
func (sd *SharedData) AddItem(item int) {
	sd.lock.Lock()
	defer sd.lock.Unlock()

	for len(sd.items) >= sd.limit {
		fmt.Println("AddItem: Reached limit, waiting...")
		sd.cond.Wait() // Wait for signal that an item was removed
	}

	sd.items = append(sd.items, item)
	fmt.Printf("Added item: %d, total: %d\n", item, len(sd.items))
	sd.cond.Broadcast() // Signal any waiting goroutines
}

// RemoveItem removes an item if available
func (sd *SharedData) RemoveItem() int {
	sd.lock.Lock()
	defer sd.lock.Unlock()

	for len(sd.items) == 0 {
		fmt.Println("RemoveItem: No items, waiting...")
		sd.cond.Wait() // Wait for signal that an item was added
	}

	item := sd.items[0]
	sd.items = sd.items[1:]
	fmt.Printf("Removed item: %d, total: %d\n", item, len(sd.items))
	sd.cond.Broadcast() // Signal any waiting goroutines
	return item
}

// Worker to add items
func workerAdd(sd *SharedData, item int) {
	time.Sleep(time.Millisecond * 500) // Simulate work
	sd.AddItem(item)
}

// Worker to remove items
func workerRemove(sd *SharedData) {
	time.Sleep(time.Millisecond * 500) // Simulate work
	sd.RemoveItem()
}

func Semaphore() {
	sd := NewSharedData(5) // Limit of items

	// Start workers
	for i := 0; i < 10; i++ {
		go workerAdd(sd, i)
		go workerRemove(sd)
	}

	// Wait to observe the output
	time.Sleep(10 * time.Second)
}

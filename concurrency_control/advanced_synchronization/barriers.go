/*
Barriers are synchronization primitives that allow threads to wait until a predefined number of threads have reached
a certain point in their execution. They are useful for coordinating phases in parallel algorithms.
*/

package main

import (
	"fmt"
	"sync"
	"time"
)

// Barrier struct to synchronize goroutines
type Barrier struct {
	total    int
	waiting  int
	mutex    sync.Mutex
	waitCond *sync.Cond
}

// NewBarrier creates a new Barrier for n goroutines
func NewBarrier(n int) *Barrier {
	b := &Barrier{total: n}
	b.waitCond = sync.NewCond(&b.mutex)
	return b
}

// Wait waits at the barrier until all goroutines have reached this point
func (b *Barrier) Wait() {
	b.mutex.Lock()
	b.waiting++
	if b.waiting == b.total {
		b.waiting = 0          // Reset for potential reuse
		b.waitCond.Broadcast() // Wake up all waiting goroutines
	} else {
		b.waitCond.Wait() // Wait until the barrier is fulfilled
	}
	b.mutex.Unlock()
}

func worker(id int, barrier *Barrier) {
	// Phase 1
	fmt.Printf("Worker %d is starting phase 1\n", id)
	time.Sleep(time.Second)
	fmt.Printf("Worker %d finished phase 1\n", id)
	barrier.Wait() // Wait for all goroutines to finish phase 1

	// Phase 2
	fmt.Printf("Worker %d is starting phase 2\n", id)
	time.Sleep(time.Second)
	fmt.Printf("Worker %d finished phase 2\n", id)
	barrier.Wait() // Optional: Wait for all goroutines to finish phase 2
}

func BarrierExample() {
	numWorkers := 5
	barrier := NewBarrier(numWorkers)
	var wg sync.WaitGroup

	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			worker(id, barrier)
		}(i)
	}

	wg.Wait()
	fmt.Println("All workers have finished their tasks.")
}

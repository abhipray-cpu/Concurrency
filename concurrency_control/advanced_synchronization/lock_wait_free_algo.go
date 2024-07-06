package main

/*
Lock-free and wait-free algorithms avoid locks altogether, ensuring that at least one thread makes progress
in a finite number of steps. These algorithms use atomic operations to manage shared data.
*/

import (
	"fmt"
	"math"
	"sync"
	"sync/atomic"
)

// AtomicInt64 provides atomic operations for int64 values
type AtomicInt64 struct {
	value int64
}

// Increment atomically increments the AtomicInt64's value by 1
func (ai *AtomicInt64) Increment() {
	atomic.AddInt64(&ai.value, 1)
}

// Decrement atomically decrements the AtomicInt64's value by 1
func (ai *AtomicInt64) Decrement() {
	atomic.AddInt64(&ai.value, -1)
}

// Load atomically loads and returns the AtomicInt64's value
func (ai *AtomicInt64) Load() int64 {
	return atomic.LoadInt64(&ai.value)
}

// LockFreeCounterWithMin tracks the count and the minimum value atomically
type LockFreeCounterWithMin struct {
	count AtomicInt64
	min   AtomicInt64
}

// NewLockFreeCounterWithMin initializes a new LockFreeCounterWithMin with min set to max int64 value
func NewLockFreeCounterWithMin() *LockFreeCounterWithMin {
	return &LockFreeCounterWithMin{
		min: AtomicInt64{value: math.MaxInt64},
	}
}

// Increment increases the counter and updates the min if necessary
func (c *LockFreeCounterWithMin) Increment() {
	c.count.Increment()
}

// Decrement decreases the counter and updates the min if necessary
func (c *LockFreeCounterWithMin) Decrement() {
	c.count.Decrement()
	currentValue := c.count.Load()
	for {
		minValue := c.min.Load()
		if currentValue >= minValue || atomic.CompareAndSwapInt64(&c.min.value, minValue, currentValue) {
			break
		}
	}
}

// Value returns the current value of the counter
func (c *LockFreeCounterWithMin) Value() int64 {
	return c.count.Load()
}

// Min returns the minimum value seen by the counter
func (c *LockFreeCounterWithMin) Min() int64 {
	return c.min.Load()
}

func LockFree() {
	var wg sync.WaitGroup
	counter := NewLockFreeCounterWithMin()

	// Simulate concurrent increments and decrements
	for i := 0; i < 100; i++ {
		wg.Add(2)
		go func() {
			defer wg.Done()
			counter.Increment()
		}()
		go func() {
			defer wg.Done()
			counter.Decrement()
		}()
	}

	wg.Wait()
	fmt.Printf("Counter value: %d\n", counter.Value())
	fmt.Printf("Minimum value seen: %d\n", counter.Min())
}

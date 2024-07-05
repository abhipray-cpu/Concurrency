package main

import (
	"fmt"
	"sync"
	"time"
)

/*
The `sync.RWMutex` type in Go is part of the `sync` package and provides a reader/writer mutual exclusion lock.
The lock can be held by an arbitrary number of readers or a single writer. Here is a list of all the methods
associated with `sync.RWMutex`:

1. **`func (*RWMutex) Lock()`**
   - Locks the mutex for writing. If the lock is already locked for reading or writing,
    `Lock` blocks until the lock is available. Only one goroutine can hold the writer lock at a time.

2. **`func (*RWMutex) Unlock()`**
   - Unlocks the mutex for writing. It panics if the mutex is not locked for writing on entry to `Unlock`.
    After `Unlock`, the lock is available for other goroutines to lock for reading or writing.

3. **`func (*RWMutex) RLock()`**
   - Locks the mutex for reading. It should not be used for recursive read locking;
    a single goroutine calling `RLock` twice will deadlock.
	 Multiple goroutines can hold the reader lock at the same time.

4. **`func (*RWMutex) RUnlock()`**
   - Unlocks the mutex for reading. It panics if the mutex is not locked for reading on entry to `RUnlock`.
    If there are no more readers, the lock becomes available for a writer.

5. **`func (*RWMutex) RLocker() Locker`**
   - Returns a `Locker` interface that implements `Lock` and `Unlock` methods by calling `RLock` and `RUnlock`,
    respectively. This is useful when a piece of code that already uses `sync.Mutex` needs to be adapted
	to allow multiple readers in a straightforward way without changing the API.

These methods allow `sync.RWMutex` to be used in scenarios where read operations outnumber write operations,
and there is a need to allow multiple goroutines to execute read operations concurrently for efficiency,
while still ensuring that write operations have exclusive access to the resource.
*/

// building an in memmory cache system using RWMutex

// Cache struct holds the data and a mutex to lock the data
type Cache struct {
	data  map[string]string
	mutex sync.RWMutex
}

// NewCache creates a new Cache instance
func NewCache() *Cache {
	return &Cache{
		data: make(map[string]string),
	}
}

// Get retrieves the value from the cache.Multiple goroutines can call Get concurrently.

func (c *Cache) Get(key string) (string, bool) {
	c.mutex.RLock() // Lock for reading
	defer c.mutex.RUnlock()
	value, exists := c.data[key]
	return value, exists
}

// Set adds or updates a value in the cache. Only one goroutine can call Set at a time.
func (c *Cache) Set(key, value string) {
	c.mutex.Lock() // Lock for writing
	defer c.mutex.Unlock()
	c.data[key] = value
}

func InMemoryCache() {
	cache := NewCache()

	// simulate concurrent writes
	for i := 0; i < 5; i++ {
		go func(n int) {
			key := fmt.Sprintf("key%d", n)
			value := fmt.Sprintf("value%d", n)
			cache.Set(key, value)
			fmt.Printf("goroutine %d: key set\n", n)
		}(i)
	}

	// simulate concurrent reads
	for i := 0; i < 5; i++ {
		go func(n int) {
			key := fmt.Sprintf("key%d", n)
			if value, exists := cache.Get(key); exists {
				fmt.Printf("goroutine %d: %s\n", n, value)
			} else {
				fmt.Printf("goroutine %d: key not found\n", n)
			}
		}(i)
	}

	// Wait a bit for goroutines to finish (for demonstration purposes)
	time.Sleep(1 * time.Second)
}

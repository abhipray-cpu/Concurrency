package main

import (
	"sync"
)

var (
	mutex1 sync.Mutex
	mutex2 sync.Mutex
)

func lockMutexes() {
	mutex1.Lock()
	mutex2.Lock() // This will never be reached in one of the goroutines if the other has already locked mutex2
	mutex2.Unlock()
	mutex1.Unlock()
}

func reverseLockMutexes() {
	mutex2.Lock()
	mutex1.Lock() // This will never be reached in one of the goroutines if the other has already locked mutex1
	mutex1.Unlock()
	mutex2.Unlock()
}

package main

import (
	"sync"
)

var (
	mutex1 sync.Mutex
	mutex2 sync.Mutex
)

func lockMutexesInOrder() {
	mutex1.Lock()
	mutex2.Lock()
	mutex2.Unlock()
	mutex1.Unlock()
}

func lockMutexesInReverseOrder() {
	mutex2.Lock()
	mutex1.Lock()
	mutex1.Unlock()
	mutex2.Unlock()
}

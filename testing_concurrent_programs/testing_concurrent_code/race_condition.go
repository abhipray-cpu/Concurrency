package main

import (
	"sync"
)

var counter int

func IncrementCounter() {
	counter++
}

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			IncrementCounter()
			wg.Done()
		}()
	}
	wg.Wait()
	println("Final counter value:", counter)
}

package main

import (
	"fmt"
	"sync"
)

/*
Collect and aggregate errors from multiple goroutines. This can be done using a slice of errors,
but make sure access to the slice is synchronized (e.g., using a mutex) to avoid data races.

*/

type ConcurrentError struct {
	sync.Mutex
	Errors []error
}

func (ce *ConcurrentError) Add(err error) {
	ce.Lock()
	defer ce.Unlock()
	ce.Errors = append(ce.Errors, err)
}

func worker3(ce *ConcurrentError) {
	if err := doWork(); err != nil {
		ce.Add(err)
	}
}

func ErrorAggregation() {
	var wg sync.WaitGroup
	ce := &ConcurrentError{}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			worker3(ce)
		}()
	}

	wg.Wait()
	if len(ce.Errors) > 0 {
		for _, err := range ce.Errors {
			fmt.Println(err)
		}
	}
}

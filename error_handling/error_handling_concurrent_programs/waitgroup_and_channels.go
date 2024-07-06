package main

/*Combine sync.WaitGroup with an error channel to wait for multiple goroutines and collect errors.
Ensure the error channel is buffered or read from concurrently to prevent blocking.*/

import (
	"log"
	"sync"
)

func worker2(wg *sync.WaitGroup, errChan chan<- error) {
	defer wg.Done()
	// Work
	if err := doWork(); err != nil {
		errChan <- err
	}
}

func WaitGroup() {
	var wg sync.WaitGroup
	errChan := make(chan error, 10) // numWorkers is the number of goroutines

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go worker2(&wg, errChan)
	}

	go func() {
		wg.Wait()
		close(errChan)
	}()

	for err := range errChan {
		if err != nil {
			log.Println(err)
		}
	}
}

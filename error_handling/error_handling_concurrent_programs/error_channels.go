package main

import "fmt"

/*Use a dedicated channel to communicate errors from goroutines back to the main goroutine or the goroutine that spawned them.
This allows centralized error handling.*/

func worker(errChan chan<- error) {
	// Simulate work
	if err := doWork(); err != nil {
		errChan <- err // Send error back to the main goroutine
	}
}

func doWork() error {
	return fmt.Errorf("error")
}

func ErrorChannel() {
	errChan := make(chan error, 1) // Buffered channel
	go worker(errChan)

	for err := range errChan {
		fmt.Println(err)
	}
}

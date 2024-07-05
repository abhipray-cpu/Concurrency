package main

import (
	"fmt"
	"sync"
	"time"
)

/*
This program does the following:

- Defines a `performTask` function that simulates doing work and then sends a result through a channel.
- In the `main` function, it starts 5 goroutines, each executing `performTask`.
- It uses a `sync.WaitGroup` to wait for all goroutines to finish their execution.
- Results from each goroutine are sent through a buffered channel `results`.
- After starting the goroutines, it waits for results using a `select` statement that also listens on a timeout channel. This demonstrates how to handle potential blocking indefinitely by using a timeout.
- Once all goroutines are done, or a timeout occurs, the program prints the results received or a timeout message and exits.

This example covers the creation and synchronization of goroutines, communication via channels, and the use of the `select` statement for non-blocking communication patterns.
*/

func performTask(id int, wg *sync.WaitGroup, results chan<- string) {
	defer wg.Done()
	time.Sleep(time.Duration(id) * 100 * time.Millisecond)
	// this will cause the 4 go routine to timeout
	if id == 4 {
		time.Sleep(10000 * time.Millisecond)

	}
	results <- fmt.Sprintf("Result from task %d", id)
}
func main() {
	var wg sync.WaitGroup
	results := make(chan string, 5) // Buffered channel

	// start multiple goroutines
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go performTask(i, &wg, results)
	}
	// Close the results channel once all goroutines are done
	go func() {
		wg.Wait()
		close(results)
	}()

	// use select to receive from that channel with a timeout
	timeout := time.After(500 * time.Millisecond)
	//This code snippet is a loop that continuously listens for two types of events using a select statement.
	//The select statement in Go allows a goroutine to wait on multiple communication operations.
	//Here's a breakdown of its components:
	for {
		select {
		case result, ok := <-results:
			if !ok {
				fmt.Println("All tasks completed.")
				return
			}
			fmt.Println(result)
		case <-timeout:
			fmt.Println("Timed out waiting for results.")
			return
		}
	}
}

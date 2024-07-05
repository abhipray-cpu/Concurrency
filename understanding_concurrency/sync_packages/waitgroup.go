package main

import (
	"fmt"
	"sync"
	"time"
)

/*
The `sync.WaitGroup` in Go is part of the `sync` package and is used to synchronize multiple goroutines. It waits for a collection of goroutines to finish executing. The main functionality of a `sync.WaitGroup` is to manage and wait for a set of goroutines to complete their tasks.

### Key Methods

1. **`Add(delta int)`**
   - This method sets the number of goroutines to wait for. It takes an integer argument `delta`, which specifies the number of goroutines to add to the waitgroup. If `delta` is negative, it reduces the count of goroutines to wait for. It's important to call `Add` before starting the goroutine.

2. **`Done()`**
   - This method decrements the WaitGroup counter by one, indicating that a goroutine has completed its work. It's equivalent to calling `Add(-1)`.

3. **`Wait()`**
   - This method blocks the calling goroutine until the WaitGroup counter returns to zero, meaning all goroutines have finished their execution. It's typically called after all `Add` calls and concurrently running goroutines.

### Best Practices

- Always pair calls to `Add` with `Done` within the goroutine to avoid panics or deadlocks.
- Avoid adding to the WaitGroup (`Add`) after calling `Wait` as this can cause a race condition.
- Use `defer wg.Done()` at the beginning of the goroutine to ensure it's called even if the goroutine exits early due to an error or return statement.

`sync.WaitGroup` is a powerful tool for managing concurrency in Go, ensuring that your program can wait for multiple tasks to complete before proceeding.
*/

// a real world scenario will be a web server that fetches data from multiple sources concurrently and waits for all the data to be fetched before responding to the client

func fetchHotelDetails(wg *sync.WaitGroup, resultChan chan<- string) {
	defer wg.Done()

	time.Sleep(2 * time.Second)
	resultChan <- "Hotel Details"
}

func fetchFlightDetails(wg *sync.WaitGroup, resultChan chan<- string) {
	defer wg.Done()
	time.Sleep(3 * time.Second)
	resultChan <- "Flight Details"
}

func fetchWeatherDetails(wg *sync.WaitGroup, resultChan chan<- string) {
	defer wg.Done()
	time.Sleep(4 * time.Second)
	resultChan <- "Weather Details"
}

func WaitGroup() {
	var wg sync.WaitGroup
	wg.Add(3)
	hotelDetailsChan := make(chan string, 1)
	flightDetailsChan := make(chan string, 1)
	weatherDetailsChan := make(chan string, 1)
	go fetchHotelDetails(&wg, hotelDetailsChan)
	go fetchFlightDetails(&wg, flightDetailsChan)
	go fetchWeatherDetails(&wg, weatherDetailsChan)

	wg.Wait()
	hotelDetails := <-hotelDetailsChan
	flightDetails := <-flightDetailsChan
	weatherDetails := <-weatherDetailsChan
	fmt.Println(hotelDetails, flightDetails, weatherDetails)
	fmt.Println("All data fetched successfully")
}

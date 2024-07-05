package main

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"
)

/*
A buffered channel in Go is a type of channel that has a limited capacity to hold values.
Unlike unbuffered channels, which require the sender and receiver to be ready to send and receive
simultaneously (thus facilitating synchronization between goroutines), buffered channels allow
senders to send values up to the channel's capacity without blocking, even if the receiver is not
ready to receive them immediately.
*/

// buffered Channel for processing incoming request in a web server

// RequestData represents the data needed to process a request.
type RequestData struct {
	RequestID string
	Data      string // Placeholder for request-specific data.
}

// ProcessedData represents the result of processing a request.
type ProcessedData struct {
	RequestID string
	Result    string // Placeholder for the processed result.
}

// Simulate database fetch operation.
func fetchFromDatabase(data RequestData) (string, error) {
	time.Sleep(100 * time.Millisecond) // Simulate delay.
	return "fetchedData:" + data.Data, nil
}

// Simulate CPU-intensive calculation.
func performCalculation(data string) (string, error) {
	time.Sleep(200 * time.Millisecond) // Simulate processing time.
	return "calculatedResult:" + data, nil
}

// databaseWorker simulates a worker that fetches data from a database.
func databaseWorker(ctx context.Context, wg *sync.WaitGroup, requests <-chan RequestData, toCalculation chan<- ProcessedData) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			return
		case requestData, ok := <-requests:
			if !ok {
				return // Channel closed.
			}
			fetchedData, err := fetchFromDatabase(requestData)
			if err != nil {
				log.Printf("Error fetching data for request %s: %v", requestData.RequestID, err)
				continue
			}
			toCalculation <- ProcessedData{RequestID: requestData.RequestID, Result: fetchedData}
		}
	}
}

func calculationWorker(ctx context.Context, wg *sync.WaitGroup, toCalculation <-chan ProcessedData, done chan<- bool) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			return
		case processedData, ok := <-toCalculation:
			if !ok {
				return // Channel closed.
			}
			calculatedResult, err := performCalculation(processedData.Result)
			if err != nil {
				log.Printf("Error calculating result for request %s: %v", processedData.RequestID, err)
				continue
			}
			fmt.Printf("Request %s processed: %s\n", processedData.RequestID, calculatedResult)
		}
	}
}

func BufferedChannel() {
	requests := make(chan RequestData, 10)        // Buffered channel for requests.
	toCalculation := make(chan ProcessedData, 10) // Buffered channel for data ready for calculation.

	ctx, cancel := context.WithCancel(context.Background())
	var wg sync.WaitGroup

	// Start database workers.
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go databaseWorker(ctx, &wg, requests, toCalculation)
	}

	// Start calculation workers.
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go calculationWorker(ctx, &wg, toCalculation, nil)
	}

	// Simulate incoming requests.
	for i := 0; i < 20; i++ {
		requests <- RequestData{RequestID: fmt.Sprintf("req%d", i), Data: fmt.Sprintf("data%d", i)}
	}

	time.Sleep(5 * time.Second) // Simulate running time.
	cancel()                    // Signal workers to stop.
	close(requests)             // Close the requests channel.
	wg.Wait()                   // Wait for all workers to finish.
	fmt.Println("All workers finished.")
}

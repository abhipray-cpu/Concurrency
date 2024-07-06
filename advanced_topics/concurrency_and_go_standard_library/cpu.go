package main

import (
	"fmt"
	"strings"
	"sync"
)

// Concurrently processes each string in a slice.
func processStringsConcurrently(stringsToProcess []string, workerCount int) []string {
	var wg sync.WaitGroup
	stringsChan := make(chan string, workerCount)
	resultsChan := make(chan string, len(stringsToProcess))

	// Start worker goroutines
	for i := 0; i < workerCount; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for str := range stringsChan {
				// Example processing: converting to uppercase
				resultsChan <- strings.ToUpper(str)
			}
		}()
	}

	// Distribute work
	for _, str := range stringsToProcess {
		stringsChan <- str
	}
	close(stringsChan)

	wg.Wait()
	close(resultsChan)

	var results []string
	for result := range resultsChan {
		results = append(results, result)
	}

	return results
}

func CPUBound() {
	inputs := []string{"hello", "world", "concurrent", "processing"}
	results := processStringsConcurrently(inputs, 4)
	fmt.Println(results)
}

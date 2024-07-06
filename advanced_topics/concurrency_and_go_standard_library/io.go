package main

import (
	"bufio"
	"fmt"
	"os"
	"sync"
)

// Concurrently reads lines from a file and processes them.
func processLinesConcurrently(filePath string, workerCount int) {
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	var wg sync.WaitGroup
	lineChan := make(chan string, workerCount)

	// Start worker goroutines
	for i := 0; i < workerCount; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for line := range lineChan {
				// Process line here
				fmt.Println(line) // Example processing
			}
		}()
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lineChan <- scanner.Text()
	}

	close(lineChan)
	wg.Wait()
}

func IOBound() {
	processLinesConcurrently("example.txt", 4)
}

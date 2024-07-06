package main

import (
	"crypto/sha256"
	"fmt"
	"sync"
)

// Concurrently computes SHA-256 hashes of a slice of strings.
func hashStringsConcurrently(stringsToHash []string, workerCount int) []string {
	var wg sync.WaitGroup
	stringsChan := make(chan string, workerCount)
	hashesChan := make(chan string, len(stringsToHash))

	// Start worker goroutines
	for i := 0; i < workerCount; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for str := range stringsChan {
				hash := sha256.Sum256([]byte(str))
				hashesChan <- fmt.Sprintf("%x", hash)
			}
		}()
	}

	// Distribute work
	for _, str := range stringsToHash {
		stringsChan <- str
	}
	close(stringsChan)

	wg.Wait()
	close(hashesChan)

	var hashes []string
	for hash := range hashesChan {
		hashes = append(hashes, hash)
	}

	return hashes
}

func Utilities() {
	inputs := []string{"hello", "world", "concurrent", "hashing"}
	hashes := hashStringsConcurrently(inputs, 4)
	fmt.Println(hashes)
}

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sync"
)

// this program demonstrates the fan-in, fan-out concurrency pattern by processing a list of files in parallel

func countWordsInFile(filepath string) (int, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return 0, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)

	wordCount := 0
	for scanner.Scan() {
		wordCount++
	}
	return wordCount, nil
}

// processfiles starts a goroutine for each file to count words.
func processFiles(files []string) <-chan int {
	results := make(chan int)
	var wg sync.WaitGroup
	for _, file := range files {
		wg.Add(1)
		go func(file string) {
			defer wg.Done()
			count, err := countWordsInFile(file)
			if err != nil {
				log.Printf("error counting words in file %s: %v", file, err)
				return
			}
			results <- count
		}(file)
	}

	// close the result channels once all the goroutines have finished
	go func() {
		wg.Wait()
		close(results)
	}()
	return results
}

// list files returns a slice of file paths in the given directory
func listFiles(directory string) ([]string, error) {
	var files []string
	err := filepath.Walk(directory, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			files = append(files, path)
		}
		return nil
	})
	return files, err
}

func ProcessFiles() {
	directory := "./texts" // Directory containing text files
	files, err := listFiles(directory)
	if err != nil {
		log.Fatalf("Failed to list files in directory %s: %v", directory, err)
	}

	results := processFiles(files)

	totalWords := 0
	for count := range results {
		totalWords += count
	}

	fmt.Printf("Total words counted: %d\n", totalWords)
}

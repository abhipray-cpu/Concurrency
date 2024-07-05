package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("This demo application shows you an example of concurrency and parallelism in Go.")

	fmt.Println("\nConcurrency:")
	start := time.Now() // Capture start time
	ConcurrencyDemo()
	duration := time.Since(start) // Calculate duration
	fmt.Printf("Execution Time: %s\n", duration)

	fmt.Println("\nParallelism:")
	start = time.Now() // Reset start time for the next measurement
	ParallelismDemo()
	duration = time.Since(start) // Calculate new duration
	fmt.Printf("Execution Time: %s\n", duration)
}

package main

import (
	"fmt"
	"runtime"
	"sync"
)

/*
Parallelism in computing refers to the technique of running multiple computations or processes simultaneously.
This approach leverages multi-core processors to execute different parts of a program, or multiple programs,
at the same time, thus potentially reducing the overall execution time. Parallelism is about doing many things
at once, as opposed to concurrency, which is about dealing with many things at once. In programming,
parallelism involves dividing a task into subtasks that can be processed in parallel, typically to achieve
faster computation by making use of multiple processing units (cores).
*/

// sumOfSquares calculates the sum of squares of the numbers in a given slice.
// It safely updates the shared result variable using a mutex to prevent race conditions.
// After updating the result, it signals the completion of its work by calling Done on the WaitGroup.
func sumOfSquares(slice []int, result *int, wg *sync.WaitGroup, mutex *sync.Mutex) {
	sum := 0 // Initialize sum to 0.
	// Iterate over each value in the slice.
	for _, value := range slice {
		sum += value * value // Add the square of the value to the sum.
	}
	mutex.Lock()   // Lock the mutex to ensure exclusive access to the result variable.
	*result += sum // Safely update the shared result variable.
	mutex.Unlock() // Unlock the mutex to allow other goroutines to update the result.
	wg.Done()      // Signal that this goroutine's work is done.
}

// parallelismDemo demonstrates how to use parallelism to calculate the sum of squares of a large set of numbers.
func ParallelismDemo() {
	numbers := make([]int, 1000) // Create a slice of 1000 integers.
	// Initialize the slice with values 1 through 1000.
	for i := range numbers {
		numbers[i] = i + 1
	}

	numCores := runtime.NumCPU()                       // Get the number of CPU cores available.
	fmt.Print("Number of CPU cores: ", numCores, "\n") // Print the number of CPU cores.

	var wg sync.WaitGroup // Create a WaitGroup to wait for all goroutines to complete.
	var mutex sync.Mutex  // Create a mutex for safely updating the shared result variable.
	result := 0           // Initialize the result variable to store the sum of squares.

	chunkSize := len(numbers) / numCores // Calculate the size of each chunk based on the number of cores.
	// Divide the work among the available CPU cores.
	for i := 0; i < numCores; i++ {
		start := i * chunkSize   // Calculate the start index for the chunk.
		end := start + chunkSize // Calculate the end index for the chunk.

		// Ensure the last chunk includes any remaining elements.
		if i == numCores-1 {
			end = len(numbers)
		}
		wg.Add(1) // Indicate that a new goroutine is starting.
		// Start a new goroutine to calculate the sum of squares for the chunk.
		go sumOfSquares(numbers[start:end], &result, &wg, &mutex)
	}
	wg.Wait()                                    // Wait for all goroutines to complete their work.
	fmt.Println("Total sum of squares:", result) // Print the total sum of squares.
}

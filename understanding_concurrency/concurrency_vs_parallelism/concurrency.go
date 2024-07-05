package main

import "fmt"

/*
Concurrency in programming refers to the ability of different parts or units of a program to execute out-of-order
or in partial order, without affecting the final outcome. This allows for parallel execution of processes or
threads to improve the efficiency of a program, especially on multi-core processors.
In the context of Go (Golang), concurrency is a fundamental concept, built into the language.
Go provides goroutines, which are functions that can run concurrently with other functions,
and channels, which are used for communication between goroutines. This model of concurrency enables developers
to write highly scalable and efficient applications.*/

// Define a function to sum the elements of a slice of integers.
// This function takes a slice of integers and a channel for sending the sum back to the caller.
func sumSlice(slice []int, resultChan chan int) {
	sum := 0 // Initialize sum to 0.
	// Iterate over each value in the slice.
	for _, value := range slice {
		sum += value // Add the value to the sum.
	}
	resultChan <- sum // Send the sum back through the channel.
}

// Define a function to demonstrate concurrency in Go.
func ConcurrencyDemo() {
	// Initialize a slice of integers.
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	// Create a channel for communicating the sum of slices.
	// Channels are used for synchronizing and passing data between goroutines.
	resultChan := make(chan int)

	// Find the midpoint of the slice to divide it into two halves.
	mid := len(numbers) / 2
	// Divide the slice into two halves.
	firstHalf := numbers[:mid]
	secondHalf := numbers[mid:]

	// Start two goroutines to sum each half of the slice concurrently.
	// Goroutines are lightweight threads managed by the Go runtime.
	go sumSlice(firstHalf, resultChan)
	go sumSlice(secondHalf, resultChan)

	// Receive the sums from the goroutines through the channel.
	// Reading from a channel is a blocking operation until data is available.
	firstSum := <-resultChan
	secondSum := <-resultChan

	// Combine the sums received from the goroutines.
	total := firstSum + secondSum

	// Print the total sum of all elements in the slice.
	fmt.Println("Total sum:", total)
}

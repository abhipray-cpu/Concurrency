package main

/*
The pipeline concurrency pattern in Go is a powerful way to organize concurrent computation,
where each stage of computation is separated into goroutines connected by channels.
Data flows through the pipeline, being processed by each stage in turn.
This pattern is particularly useful for stream processing or any scenario
where data can be processed in discrete steps.
*/

/*
1)Data processing pipeline(ETL)
2)Stream processing
3)Image and video streaming
4)Financial systems
5)Scientific computing
6)Machine learing workflows
*/

import (
	"fmt"
	"time"
)

/*
Generate Numbers (Stage 1): Generate a sequence of numbers to be processed.
Square Numbers (Stage 2): Receive numbers from the first stage and square them.
Print Numbers (Stage 3): Receive squared numbers from the second stage and print them.
*/

// gen generates integers in a separate goroutine and sends them to the returned channel.
func gen(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()
	return out
}

// sq receives integers from a channel, squares them, and sends them to the returned channel.
func sq(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
	}()
	return out
}

func main() {

	start := time.Now()
	// Stage 1: Generate numbers.
	genOut := gen(2, 3, 12, 232323, 121212, 2323232)

	// Stage 2: Square numbers.
	sqOut := sq(genOut)

	// Stage 3: Print numbers.
	for n := range sqOut {
		fmt.Println(n)
	}
	elapsed := time.Since(start) // Calculate elapsed time
	fmt.Printf("Execution took %s\n", elapsed)
}

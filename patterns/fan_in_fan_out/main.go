package main

import "fmt"

/*
The fan-in, fan-out concurrency pattern is a powerful technique in Go for managing parallel work
and consolidating results. It involves fanning out a series of goroutines to handle work in
parallel (fan-out) and then collecting all their results into a single channel (fan-in).
This pattern is particularly useful for tasks that can be broken down into independent,
concurrent operations.
*/

/*
1)Web crawling
2)Data processing pipeline
3)Image processing
4)Load testing tools
5)Paralle file processing
6)Distributed task execution
7)Real time data aggregation
*/

func main() {
	fmt.Println("Fan-in, Fan-out Concurrency Pattern")
	fmt.Println("1)Parallel file processing")
	ProcessFiles()
	fmt.Println("2)Real time data aggregation")
	AggregateSum()
}

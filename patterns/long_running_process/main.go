package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"sync"
	"time"
)

func readFile(filePath string, dataCh chan<- int) {
	file, err := os.Open(filePath)
	if err != nil {
		log.Println("Error opening file", err)
		close(dataCh)
		return
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		num, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Println("Error converting string to int", err)
			continue
		}
		dataCh <- num

	}
	close(dataCh)
}

func processData(dataCh <-chan int, resultCh chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for num := range dataCh {
		// Example processing: simply pass the number through
		resultCh <- num
	}
}

func aggregateResults(resultCh <-chan int, doneCh chan<- int) {
	sum := 0
	for num := range resultCh {
		sum += num
	}
	doneCh <- sum
}
func init() {
	DatGenerator()
}
func main() {
	startTime := time.Now()
	dataCh := make(chan int, 100)
	resultCh := make(chan int, 100)
	doneCh := make(chan int)

	var wg sync.WaitGroup

	go readFile("large_dataset.txt", dataCh)

	numWorkers := 5 // you can adjust this number based on your system
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go processData(dataCh, resultCh, &wg)
	}

	// ensure all workers are done before closing the resultCh
	go func() {
		wg.Wait()
		close(resultCh)
	}()

	go aggregateResults(resultCh, doneCh)

	// wait for aggregation to complete
	sum := <-doneCh
	fmt.Println("Sum of all numbers:", sum)
	fmt.Println(time.Since(startTime))
}

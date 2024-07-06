// main.go
package main

import (
	"fmt"
	"sync"
)

func main() {
	logger := NewLogger(100)
	var wg sync.WaitGroup
	go logger.StartLogWriter("log.txt")

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for j := 0; j < 10; j++ {
				logger.LogMessage(fmt.Sprintf("Goroutine %d: log message %d", id, j))
			}
		}(i)
	}

	wg.Wait()
	fmt.Println("Logging completed.")
}

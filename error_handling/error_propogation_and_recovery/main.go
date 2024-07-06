package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	var wg sync.WaitGroup
	ea := ErrorAggregator{}

	tasks := []Task{
		{ID: 1},
		{ID: 2},
		{ID: 3},
	}

	for _, task := range tasks {
		wg.Add(1)
		go func(t Task) {
			defer wg.Done()
			// Execute task with retry logic
			if err := t.ExecuteWithRetry(2); err != nil {
				fmt.Printf("Task %d ultimately failed after retries: %s\n", t.ID, err)
				ea.Add(err)
			}
		}(task)
	}

	wg.Wait()

	if err := ea.Aggregate(); err != nil {
		fmt.Println("Errors encountered:", err)
	} else {
		fmt.Println("All tasks completed successfully")
	}
}

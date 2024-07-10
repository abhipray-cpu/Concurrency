package main

import (
	"fmt"
	"math/rand"
	"net/http"
	_ "net/http/pprof"
	"sync"
	"time"
)

type Task struct {
	ID       int
	Duration time.Duration
}

func worker(tasks <-chan Task, wg *sync.WaitGroup) {
	for task := range tasks {
		fmt.Printf("Worker processing task: %d\n", task.ID)
		time.Sleep(task.Duration)
		wg.Done()
	}
}

func main() {
	// Start pprof server
	go func() {
		fmt.Println("Starting pprof server at :6060")
		if err := http.ListenAndServe("localhost:6060", nil); err != nil {
			fmt.Printf("pprof server error: %s\n", err)
		}
	}()

	numWorkers := 5
	numTasks := 20000000

	tasks := make(chan Task, numTasks)
	var wg sync.WaitGroup

	// Start workers
	for i := 0; i < numWorkers; i++ {
		go worker(tasks, &wg)
	}

	// Generate tasks
	for i := 0; i < numTasks; i++ {
		task := Task{
			ID:       i,
			Duration: time.Millisecond * time.Duration(rand.Intn(1000)),
		}
		wg.Add(1)
		tasks <- task
	}

	wg.Wait()
	close(tasks)
	fmt.Println("All tasks processed")
}

// command go tool pprof http://localhost:6060/debug/pprof/profile?seconds=30

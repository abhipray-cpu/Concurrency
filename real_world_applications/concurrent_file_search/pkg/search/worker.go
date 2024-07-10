package search

import (
	"fmt"
	"sync"
)

type Job struct {
	FilePath      string
	SearchPattern string
	SearchMode    string
}

type Result struct {
	Matches []string
	Found   bool
}

func Worker(workerId int, jobs <-chan Job, results chan<- Result, wg *sync.WaitGroup) {
	defer wg.Done()
	for job := range jobs {
		fmt.Printf("Worker %d started job for %s\n", workerId, job.FilePath)
		var matches []string
		var found bool

		if job.SearchMode == "regex" {
			matches, found = RegexpSearchFile(job.FilePath, job.SearchPattern)
		} else {
			matches, found = SearchFile(job.FilePath, job.SearchPattern)
		}

		result := Result{
			Matches: matches,
			Found:   found,
		}
		results <- result
	}
}

func StartDispatcher(workerCount int) (chan Job, chan Result) {
	jobs := make(chan Job, 100)
	results := make(chan Result, 100)
	var wg sync.WaitGroup
	for i := 0; i < workerCount; i++ {
		wg.Add(1)
		go Worker(i, jobs, results, &wg)
	}

	// close the results channel when all workers are done
	go func() {
		wg.Wait()
		close(results)
	}()
	return jobs, results
}

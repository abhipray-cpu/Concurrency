package main

import (
	"encoding/json"
	"fmt"
	"sync"
)

type Data struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

// Concurrently decodes JSON strings.
func decodeJSONConcurrently(jsonStrings []string, workerCount int) []Data {
	var wg sync.WaitGroup
	jsonChan := make(chan string, workerCount)
	resultsChan := make(chan Data, len(jsonStrings))

	// Start worker goroutines
	for i := 0; i < workerCount; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for jsonString := range jsonChan {
				var data Data
				if err := json.Unmarshal([]byte(jsonString), &data); err == nil {
					resultsChan <- data
				}
			}
		}()
	}

	// Distribute work
	for _, jsonString := range jsonStrings {
		jsonChan <- jsonString
	}
	close(jsonChan)

	wg.Wait()
	close(resultsChan)

	var results []Data
	for result := range resultsChan {
		results = append(results, result)
	}

	return results
}

func DataProcessing() {
	jsonInputs := []string{`{"name":"Alice","age":30}`, `{"name":"Bob","age":25}`}
	results := decodeJSONConcurrently(jsonInputs, 4)
	fmt.Println(results)
}

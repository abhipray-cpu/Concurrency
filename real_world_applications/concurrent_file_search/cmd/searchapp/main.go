/*
Parse command-line arguments for the root directory and search pattern.
Initialize and start the search process.
*/

package main

import (
	"concurrent_file_search/pkg/search"
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	// parse command line arguments
	rootDir := flag.String("root", ".", "Root directory to start the search from")
	pattern := flag.String("pattern", "", "Search pattern")
	searchMode := flag.String("mode", "plain", "Search mode: simple|regex")
	flag.Parse()

	if *pattern == "" {
		fmt.Println("Search pattern must not be empty")
		os.Exit(1)
	}

	// initialize dispatcher
	jobs, results := search.StartDispatcher(5) // starting with 5 workers

	// walk through directory structure and send jobs to workers
	go func() {
		filepath.Walk(*rootDir, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				fmt.Println("Error accessing path:", path, ":", err)
				return err
			}

			if !info.IsDir() {
				jobs <- search.Job{FilePath: path, SearchPattern: *pattern, SearchMode: *searchMode}
			}
			return nil
		})
		close(jobs)
	}()

	// Collect and display the results
	for result := range results {
		if result.Found {
			fmt.Printf("Match found: %v\n", result.Matches)
		}
	}
}

// go run main.go --pattern="searchPattern" --searchMode="mode"

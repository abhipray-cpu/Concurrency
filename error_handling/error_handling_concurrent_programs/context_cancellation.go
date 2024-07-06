package main

import (
	"context"
	"log"
	"time"
)

/*
Use context.Context to handle cancellation and timeouts across goroutines.
This doesn't directly handle errors but is crucial for controlling goroutine
lifecycles and avoiding leaks when errors occur.

*/

func worker4(ctx context.Context) error {
	select {
	// Handle cancellation
	case <-ctx.Done():
		return ctx.Err() // Handle cancellation or deadline exceed
	case <-time.After(time.Hour): // Simulate long-running work
		// Work completed
	}
	return nil
}

func Context() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := worker4(ctx)
	if err != nil {
		// Handle error
		log.Println("Error:", err)
		// Perform error handling logic here
	}
}

package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

/*
A "thread-safe map" refers to a map data structure that is safe to be accessed or modified concurrently by
multiple threads (or goroutines, in the context of Go) without causing data races or inconsistencies.
Thread safety is crucial in concurrent programming when shared data structures can be accessed by multiple
threads simultaneously. Without proper synchronization, concurrent accesses can lead to unpredictable behavior,
crashes, or corrupted data.In Go, the built-in map type is not thread-safe, meaning that concurrent reads and
writes to a map without proper synchronization can cause a runtime panic. To create a thread-safe map in Go,
you typically use synchronization primitives from the `sync` package, such as `sync.Mutex` or `sync.RWMutex`,
to control access to the map.Additionally, Go provides a specialized map type for concurrent use without
the need for explicit locks: `sync.Map`. The `sync.Map` type is part of the `sync` package and is designed
to be safe for concurrent access. It has methods for storing, retrieving, and deleting items that internally
manage synchronization, making it a convenient choice for certain use cases where maps are accessed by
multiple goroutines.
*/

// building a web application cache
// List of all the methods:
/*
1)Load
2)Store
3)LoadOrStore
4)Delete
5)Range
*/

var cache sync.Map

func expensiveOperation(key string) string {
	// simulating an expensive operation
	time.Sleep(2 * time.Second)
	return "data for " + key
}

func handler(w http.ResponseWriter, r *http.Request) {
	key := r.URL.Query().Get("key")
	if key == "" {
		http.Error(w, "Key is required", http.StatusBadRequest)
		return
	}

	action := r.URL.Query().Get("action")
	switch action {
	case "get":
		if value, ok := cache.Load(key); ok {
			fmt.Fprintf(w, "Loaded: %v", value)
		} else {
			fmt.Fprintln(w, "Key not found")
		}
	case "store":
		value := expensiveOperation(key)
		cache.Store(key, value)
		fmt.Fprintf(w, "Stored: %v", value)
	case "loadOrStore":
		value, _ := cache.LoadOrStore(key, expensiveOperation(key))
		fmt.Fprintf(w, "LoadOrStore: %v", value)
	case "delete":
		cache.Delete(key)
		fmt.Fprintln(w, "Deleted")
	case "range":
		cache.Range(func(k, v interface{}) bool {
			fmt.Fprintf(w, "%v: %v\n", k, v)
			return true
		})
	default:
		http.Error(w, "Invalid action", http.StatusBadRequest)
	}
}

func SyncMap() {
	http.HandleFunc("/", handler)
	fmt.Println("Server is running on port 8080")
	http.ListenAndServe(":8081", nil)
	// example url http://localhost:8081/?key=1&action=store
}

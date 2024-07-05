package main

import (
	"fmt"
	"io"
	"net/http"
	"sync"
)

/*
In Go, a pool is a set of temporary objects that may be individually saved and retrieved.
The `sync.Pool` type is designed to provide a pool of temporary objects that can be reused
to avoid unnecessary allocations. This can be particularly useful in applications that perform
a large number of allocations of short-lived objects, as it can help to reduce garbage collection overhead.

*/

/*
Sync.Pool type in go provides two primary methods:

1. **Get**: This method retrieves an item from the pool. If the pool is empty,
 `Get` will call the `New` function defined in the pool (if any) to create a new item. The method signature is:
   ```go
   func (p *Pool) Get() interface{}
   ```
   It returns an `interface{}` type, which means you will typically need to perform a type assertion to convert
it back to the expected type of the pooled objects.

2. **Put**: This method adds an item back to the pool, making it available for subsequent `Get` calls.
 The method signature is:
   ```go
   func (p *Pool) Put(x interface{})
   ```
   The `Put` method takes a parameter of type `interface{}`, allowing you to return any type of object to the pool.
    However, you should only put objects of the same type into a specific pool to avoid runtime panics due
	to invalid type assertions.

These methods are designed to be simple yet flexible, allowing `sync.Pool` to be used in a variety of scenarios where object reuse can significantly reduce the overhead of garbage collection by minimizing the number of allocations.

*/

var bufferPool = sync.Pool{
	New: func() interface{} {
		// Allocatte a new byte slice with a capacity of 1024
		return make([]byte, 1024)
	},
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	// Get a buffer from the pool
	buf := bufferPool.Get().([]byte)
	defer bufferPool.Put(buf) // Make sure to return buffer to the pool

	// Initialize an empty slice to hold the data read from the request body
	var data []byte

	// will read the entire request body in chunks till the end of file is reached the chunk size is 1024 byte
	for {
		// Read the request body using the buffer
		n, err := r.Body.Read(buf) // n is the number fo bytes read
		if n > 0 {
			// Append the data read to the data slice
			fmt.Println("Chunk Size: ", n)
			data = append(data, buf[:n]...)
		}
		if err == io.EOF {
			break // End of file (request body), break out of the loop
		} else if err != nil {
			fmt.Println(err)
			http.Error(w, "Failed to read request body", http.StatusInternalServerError)
			return
		}
	}

	// Process the request as per requirement
	// For example, write the data read from the request body back to the response
	w.Write(data)
}

func PoolHandler() {
	http.HandleFunc("/", handleRequest)
	http.ListenAndServe(":8081", nil)
}

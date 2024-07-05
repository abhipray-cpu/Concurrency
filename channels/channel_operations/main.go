// Package main defines the entry point for a network application that can operate in either server or client mode.
package main

import (
	"flag" // Import flag for parsing command-line options.
	"fmt"  // Import fmt for formatted I/O operations.
)

func main() {
	// Define a string flag "mode" with default value "server". This flag specifies the application's mode of operation.
	mode := flag.String("mode", "server", "start in client or server mode")
	// Define a string flag "address" with default value "localhost:8080". This flag specifies the network address to connect to or listen on.
	address := flag.String("address", "localhost:8080", "address to connect to or listen on")
	// Parse the command-line flags.
	flag.Parse()

	// Switch on the mode specified by the user.
	switch *mode {
	case "server":
		// If mode is "server", create a new server instance.
		server := NewServer()
		// Start the server on the specified address.
		server.Start(*address)
	case "client":
		// If mode is "client", start the client and connect to the specified address.
		StartClient(*address)
	default:
		// If an unknown mode is specified, print an error message.
		fmt.Println("Unknown mode:", *mode)
	}
}

/*
test
1)server:go run . -mode=server -address=localhost:8080
2)client:go run . -mode=client -address=localhost:8080
*/

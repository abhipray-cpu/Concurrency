// Package main is the entry point for the client application.
package main

import (
	"bufio"   // Import bufio for buffered I/O operations.
	"fmt"     // Import fmt for formatted I/O operations.
	"net"     // Import net for network I/O operations.
	"os"      // Import os for operating system functionality, like reading from stdin.
	"strings" // Import strings for string manipulation functions.
)

// StartClient initiates a connection to a server at the specified address and starts listening for user input to send as messages.
func StartClient(serverAddress string) {
	// Dial connects to the server at serverAddress using TCP.
	conn, err := net.Dial("tcp", serverAddress)
	if err != nil {
		// If there's an error connecting to the server, print the error and return.
		fmt.Println("Error connecting to server:", err)
		return
	}
	// Ensure the connection is closed when the function returns.
	defer conn.Close()

	// Start a goroutine to read messages from the server.
	go readMessages(conn)

	// Create a new scanner to read from standard input (stdin).
	scanner := bufio.NewScanner(os.Stdin)
	// Continuously read from stdin until there's an error or it's closed.
	for scanner.Scan() {
		// Read the next line from stdin.
		msg := scanner.Text()
		// If the message is only whitespace, skip it.
		if strings.TrimSpace(msg) == "" {
			continue
		}
		// Send the message to the server, appending a newline character.
		conn.Write([]byte(msg + "\n"))
	}
}

// readMessages continuously reads messages from the server and prints them.
func readMessages(conn net.Conn) {
	// Create a new scanner to read from the connection.
	scanner := bufio.NewScanner(conn)
	// Continuously read from the connection until there's an error or it's closed.
	for scanner.Scan() {
		// Print the message received from the server.
		fmt.Println(scanner.Text())
	}
}

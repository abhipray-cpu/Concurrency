// Package main defines the server part of a simple chat application.
package main

import (
	"bufio"   // Import bufio for buffered I/O operations.
	"fmt"     // Import fmt for formatted I/O operations.
	"net"     // Import net for network I/O operations.
	"strings" // Import strings for string manipulation functions.
	"sync"    // Import sync for synchronization primitives like Mutex.
)

// Client represents a single chat client.
type Client struct {
	conn     net.Conn    // Network connection to the client.
	name     string      // Name of the client.
	messages chan string // Channel for sending messages to the client.
}

// Server represents the chat server.
type Server struct {
	clients map[*Client]bool // Map of all connected clients.
	lock    sync.Mutex       // Mutex to protect access to the clients map.
}

// NewServer creates and returns a new Server instance.
func NewServer() *Server {
	return &Server{
		clients: make(map[*Client]bool), // Initialize the clients map.
	}
}

// Start begins the server's operation on the specified address.
func (s *Server) Start(address string) {
	// Listen on the specified network address.
	listener, err := net.Listen("tcp", address)
	if err != nil {
		// If there's an error starting the server, print the error and return.
		fmt.Println("Error starting server:", err)
		return
	}
	defer listener.Close() // Ensure the listener is closed when the function returns.
	fmt.Println("Server started on", address)

	for {
		// Accept a new connection.
		conn, err := listener.Accept()
		if err != nil {
			// If there's an error accepting a connection, print the error and continue to the next iteration.
			fmt.Println("Error accepting connection:", err)
			continue
		}
		// Create a new client instance.
		client := &Client{conn: conn, messages: make(chan string)}
		// Add the client to the server's clients map.
		s.clients[client] = true
		// Start handling the client in a new goroutine.
		go s.handleClient(client)
		// Start sending messages to the client in a new goroutine.
		go s.sendMessages(client)
	}
}

// handleClient manages communication with a single client.
func (s *Server) handleClient(client *Client) {
	defer func() {
		// Ensure the client's connection is closed and the client is removed from the server's clients map when the function returns.
		client.conn.Close()
		s.lock.Lock()
		delete(s.clients, client)
		s.lock.Unlock()
		fmt.Println(client.name, "disconnected")
	}()

	// Prompt the client to enter their name.
	client.conn.Write([]byte("Enter your name: "))
	scanner := bufio.NewScanner(client.conn)
	if scanner.Scan() {
		// Read the client's name.
		client.name = scanner.Text()
	}
	fmt.Println(client.name, "connected")

	// Continuously read messages from the client.
	for scanner.Scan() {
		msg := scanner.Text()
		// Ignore messages that are only whitespace.
		if strings.TrimSpace(msg) == "" {
			continue
		}
		// Broadcast the message to all clients.
		s.broadcast(client.name + ": " + msg)
	}
}

// sendMessages sends messages to a client from their messages channel.
func (s *Server) sendMessages(client *Client) {
	for msg := range client.messages {
		// Send each message to the client, appending a newline character.
		client.conn.Write([]byte(msg + "\n"))
	}
}

// broadcast sends a message to all connected clients.
func (s *Server) broadcast(message string) {
	s.lock.Lock()         // Lock the server's clients map.
	defer s.lock.Unlock() // Ensure the lock is released when the function returns.

	fmt.Println("Broadcast:", message)
	for client := range s.clients {
		select {
		case client.messages <- message:
			// Attempt to send the message to the client.
		default:
			// If the client's messages channel is blocked, continue to the next client.
			continue
		}
	}
}

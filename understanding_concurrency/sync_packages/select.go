package main

import (
	"fmt"
	"net"
	"time"
)

/*
The `select` statement in Go is a powerful feature used for choosing from multiple send/receive channel operations.
It blocks until one of its cases can run, then it executes that case. If multiple cases can proceed,
`select` picks one at random, ensuring fairness.

### Key Features of `select`:

- **Concurrency Control**: Primarily used with goroutines and channels to manage multiple channel operations.

- **Non-blocking Operations**: Can be used to perform non-blocking sends, receives, and even non-blocking default
 cases.

 - **Random Selection**: If multiple cases are ready, `select` chooses one at random, ensuring no case is starved.

- **Synchronization**: Helps in synchronizing multiple goroutines without explicit locks or condition variables.

*/

// building a chat server that manages multiple clients using `select` statement.

type Client struct {
	send chan string
}

var (
	entering = make(chan Client) // new clients
	leaving  = make(chan Client) // disconnected clients
	messages = make(chan string) // All incoming client messages
)

func broadcaster() {
	clients := make(map[Client]bool) // all connected clients
	for {
		select {
		case msg := <-messages:
			// Broadcast incoming message to all clients' outgoing message channels.
			for cli := range clients {
				cli.send <- msg
			}
		case cli := <-entering:
			clients[cli] = true

		case cli := <-leaving:
			delete(clients, cli)
			close(cli.send)
		}
	}
}

func handleConn(conn net.Conn) {
	cli := Client{send: make(chan string, 10)}
	go clientWriter(conn, cli.send)

	// simulate client name and message reception
	who := conn.RemoteAddr().String()
	cli.send <- "You are " + who
	messages <- who + " has arrived"
	entering <- cli
	time.Sleep(10 * time.Second) // Simulate chat activity
	leaving <- cli
	messages <- who + " has left"
	conn.Close()
}
func clientWriter(conn net.Conn, ch <-chan string) {
	for msg := range ch {
		fmt.Fprintln(conn, msg) // NOTE: Ignoring network errors
	}
}

func Chat() {
	listener, _ := net.Listen("tcp", "localhost:8000")
	go broadcaster()
	for {
		conn, _ := listener.Accept()
		go handleConn(conn)
	}
}

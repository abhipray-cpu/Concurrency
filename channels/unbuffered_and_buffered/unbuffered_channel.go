package main

import (
	"fmt"
	"time"
)

/*
An unbuffered channel in Go is a channel with no capacity to hold messages before they are received.
It is created using the make function without specifying a capacity.
Syncronization between the sender and receiver is required for the communication to take place.
Direct Communication: The sender will block until the receiver is ready to receive the message.
Guranteed Freshness: The receiver will always receive the latest message sent by the sender.
Dedalock Potential: Care must be taken to avoid deadlocks. If a goroutine is waiting to send on an
unbuffered channel and no other goroutine is ready to receive, the program will deadlock.
Similarly, if a goroutine is waiting to receive and no one sends a value, it will also deadloc
*/

// ordersing system in a restaurant

type Order struct {
	orderId int
	details string
}

func UnbufferedChannel() {
	ordersChannel := make(chan Order)
	readyChannel := make(chan Order)

	// Simulate wait staff taking orders
	go waitStaff(ordersChannel, readyChannel)

	// Simulate kitchen preparing orders
	go kitchen(ordersChannel, readyChannel)

	// Wait for input to exit (simulate real-time operation)
	fmt.Scanln()
}

func waitStaff(ordersChan, readyChan chan Order) {
	for i := 1; i <= 5; i++ {
		order := Order{orderId: i, details: fmt.Sprintf("Order #%d", i)}
		fmt.Printf("Waiter: Order for %s taken\n", order.details)
		ordersChan <- order       // Send order to kitchen.
		readyOrder := <-readyChan // Wait for order to be ready.
		fmt.Printf("Waiter: Order #%d ready\n", readyOrder.orderId)
	}
}

func kitchen(ordersChan, readyChan chan Order) {
	for order := range ordersChan {
		fmt.Printf("Kitchen: Preparing %s received\n", order.details)
		time.Sleep(2 * time.Second) // Simulate cooking time.
		fmt.Printf("Kitchen: %s is ready\n", order.details)
		readyChan <- order // Send ready order to waiter.
	}
}

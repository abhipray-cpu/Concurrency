package main

import (
	"fmt"
	"sync"
	"time"
)

/*
Dynamic channel creation in Go involves programmatically creating channels based on runtime
conditions or inputs, rather than declaring them statically at compile time.
This concept is particularly useful in scenarios where the number of goroutines or
tasks to be managed is not known beforehand and can vary during the execution of a program.
It allows for more flexible and scalable concurrent designs by adapting to the workload dynamically.
*/

/*
Some real world applications
1)Task distribution: A common use case for dynamic channel creation is task distribution among a pool of workers.
2)Event handling: In event-driven systems, dynamic channels can be used to handle incoming events from multiple sources.
3)Resource management: Dynamic channels can be used to manage resources such as connections, files, or memory buffers.
4)Pub/Sub systems: In publish-subscribe systems, dynamic channels can be used to create topic-specific channels for subscribers.
5)Real time data processing: In real-time data processing applications, dynamic channels can be used to process data streams concurrently.
*/

// let's implement a pub sub model using dynamic channel creation and teardown

type PubSub struct {
	mu     sync.RWMutex
	topics map[string]chan string // map of topic names to channels
}

func NewPubSub() *PubSub {
	return &PubSub{
		topics: make(map[string]chan string),
	}
}

func (ps *PubSub) CreateTopic(topic string) {
	ps.mu.Lock()
	defer ps.mu.Unlock()
	if _, ok := ps.topics[topic]; !ok {
		ps.topics[topic] = make(chan string, 10) // Buffered channel
	}
}
func (ps *PubSub) Publish(topic, message string) {
	ps.mu.RLock()
	defer ps.mu.RUnlock()

	if ch, ok := ps.topics[topic]; ok {
		ch <- message
	}
}
func (ps *PubSub) Subscribe(topic string) <-chan string {
	ps.mu.RLock()
	defer ps.mu.RUnlock()

	if ch, ok := ps.topics[topic]; ok {
		return ch
	}
	return nil
}

func main() {
	ps := NewPubSub()

	// Create topics
	ps.CreateTopic("news")
	ps.CreateTopic("weather")

	// Subscribe to topics
	newsChannel := ps.Subscribe("news")
	weatherChannel := ps.Subscribe("weather")

	// Start a goroutine to listen to subscription channels
	go func() {
		for {
			select {
			case news := <-newsChannel:
				fmt.Println("News:", news)
			case weather := <-weatherChannel:
				fmt.Println("Weather:", weather)
			}
		}
	}()

	// Publish messages
	ps.Publish("news", "New Go version released")
	ps.Publish("weather", "Sunny today with a chance of rain")
	time.Sleep(1 * time.Second)
	ps.Publish("news", "Go wins programming language of the year")
	// Wait to receive messages
	// In a real application, use proper synchronization or wait mechanism
	// Here, we use a simple sleep for demonstration
	time.Sleep(10 * time.Second)
}

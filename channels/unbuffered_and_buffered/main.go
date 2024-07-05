package main

import "fmt"

func main() {
	fmt.Println("Buffered Channel")
	BufferedChannel()
	fmt.Println("Unbuffered Channel")
	UnbufferedChannel()
}

/*
When to Use Buffered Channels:

1)Asynchronous Communication: When you want goroutines to send data without waiting for a receiver to be ready.
Buffered channels allow senders to continue without immediate blocking, up to the channel's capacity.

2)Rate Limiting: To limit the rate of processing, you can use buffered channels as a semaphore, controlling
the number of concurrent operations.

3)Worker Pools: In scenarios where you have multiple workers processing tasks concurrently, a buffered channel
can serve as a task queue, with workers pulling tasks as they are able to process them.

4)Gathering Results from Multiple Goroutines: When collecting results from several goroutines, a buffered
channel can store results as they come in until you're ready to process them.

5)Control Overload: To prevent a goroutine from being overwhelmed by too many incoming requests or data,
a buffered channel can act as a buffer, smoothing out spikes in demand.
*/

/*
When to Use Unbuffered Channels:

1)Synchronization: When you need to synchronize two goroutines, ensuring that a send operation directly
corresponds to a receive operation, providing a guarantee that a message is received as soon as it's sent.

2)Guarantee of Fresh Data: To ensure that data is always fresh and processed immediately, unbuffered channels
enforce direct handoff of data between goroutines.

3)Simple Coordination: For simple scenarios where you want to coordinate the start or end of goroutines,
unbuffered channels provide a straightforward mechanism.

4)Ensuring Mutual Exclusion: In cases where you need to ensure that only one goroutine accesses a resource
at a time, unbuffered channels can serve as a mechanism for mutual exclusion.

5)Order Preservation: When the order of operations or data processing is critical, unbuffered channels
ensure that messages are processed in the exact order they are sent, due to the direct synchronization
between sender and receiver.
*/

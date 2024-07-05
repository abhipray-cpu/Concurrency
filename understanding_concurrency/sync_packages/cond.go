package main

import (
	"fmt"
	"sync"
	"time"
)

/*
The sync.Cond in Go is a synchronization primitive that can block a goroutine until a certain condition is met.
It is part of the sync package, which provides basic synchronization primitives such as mutual exclusion locks
(sync.Mutex). A sync.Cond is used in conjunction with a sync.Locker (usually a sync.Mutex or sync.RWMutex)
to allow goroutines to wait for or announce the occurrence of certain conditions within the program,
enabling more sophisticated synchronization scenarios.
*/

/*
Key Concepts of sync.Cond:
1)Condition Variable: A sync.Cond acts as a condition variable, a synchronization primitive that does not
have a value but instead enables goroutines to wait for or signal the occurrence of certain conditions.

2)Waiting: A goroutine can wait on a sync.Cond by calling its Wait method. This puts the goroutine
into a waiting state, where it will remain until it is explicitly woken up by another goroutine.

3)Signaling: A goroutine can wake up another goroutine that is waiting on a sync.Cond by calling
the Signal method. This wakes up one of the waiting goroutines. To wake up all waiting goroutines,
the Broadcast method can be used.

4)Mutex Lock: A sync.Cond is associated with a mutex lock (sync.Locker), which must be locked by
the goroutine before calling Wait, Signal, or Broadcast. The Wait method automatically unlocks the mutex
while waiting and re-locks it once it is awakened.
*/

/*
Exposed methods
1. `Wait()`
   - Waits for a signal, automatically unlocking the mutex and suspending execution of the calling goroutine. It re-locks the mutex once the goroutine is awakened.

2. `Signal()`
   - Wakes up one goroutine waiting on the condition variable, if there is any.

3. `Broadcast()`
   - Wakes up all goroutines waiting on the condition variable.
*/

// shared queue and the its synchronization primitives

var queue []int
var mutex sync.Mutex
var cond = sync.NewCond(&mutex)

func producer(id int) {
	for i := 0; i < 5; i++ {
		mutex.Lock()
		item := id*10 + i // Produce an item
		queue = append(queue, item)
		fmt.Printf("Producer %d produced %d\n", id, item)
		if i == 0 {
			cond.Signal() //wake up one consumer to start consuming
		} else if i == 4 {
			cond.Broadcast() //wake up all consumers to start consuming
		}
		mutex.Unlock()
		time.Sleep(time.Millisecond * 500)
	}
}

func consumer(id int) {
	for {
		mutex.Lock()
		for len(queue) == 0 {
			cond.Wait() // the consumer will only check the len(queue) when sigalled from the producer
		}
		item := queue[0]
		queue = queue[1:]
		fmt.Printf("Consumer %d consumed %d\n", id, item)
		mutex.Unlock()
		time.Sleep(time.Millisecond * 100)
	}
}

func CondDemo() {
	for i := 0; i < 3; i++ {
		go producer(i)
		go consumer(i)
	}

	// prevent the main goroutine from exiting
	time.Sleep(time.Second * 10)
}

/*
This Go code demonstrates a producer-consumer problem using sync.Cond for synchronization.
It involves multiple producers and consumers operating on a shared queue.
The producers generate items and place them in the queue, while the consumers remove items
from the queue and process them. The sync.Cond is used to synchronize access to the queue
and efficiently manage the waiting state of consumers when the queue is empty.

Global Variables
queue []int: A slice that acts as a shared queue between producers and consumers.
mutex sync.Mutex: A mutex lock to protect access to the shared queue.
cond: A condition variable associated with the mutex lock. It is used to signal and broadcast to goroutines waiting on a specific condition (in this case, the availability of items in the queue).

Producer Function
producer(id int): A function that simulates a producer. Each producer is identified by an id.
It produces items in a loop, each item being an integer. The item's value is calculated using the producer's id
and the loop index i to ensure uniqueness.
After producing an item, it appends the item to the shared queue.
It uses cond.Signal() to wake up one waiting consumer when the first item is produced and cond.Broadcast()
to wake up all waiting consumers when the last item is produced in the loop.
The producer acquires the mutex lock before modifying the queue and releases it afterward to ensure
safe access to the shared resource.
It pauses for a short duration between producing items to simulate work.

Consumer Function
consumer(id int): A function that simulates a consumer. Each consumer is identified by an id.
It continuously tries to consume items from the shared queue.
If the queue is empty, the consumer waits on the condition variable cond until it is signaled or broadcasted by a producer.
Once signaled and the queue is not empty, it consumes (removes) the first item from the queue and processes it.
Similar to the producer, the consumer acquires the mutex lock before accessing the queue and releases it afterward.
It pauses for a short duration after consuming an item to simulate processing work.

CondDemo Function
CondDemo(): The main function that starts the producer-consumer simulation.
It launches multiple producers and consumers as goroutines by calling producer(i) and consumer(i) in a loop, where i is the goroutine's id.
To prevent the main goroutine from exiting immediately and allow the producers and consumers to run, it pauses the main goroutine for a specified duration at the end of the function.
This code showcases the use of sync.Cond for managing synchronization between goroutines in a classic producer-consumer scenario, demonstrating how to efficiently handle the waiting and signaling of goroutines based on the state of a shared resource.


*/

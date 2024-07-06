package main

import "fmt"

func main() {
	fmt.Println("This code will demonstrate different advanced synchronization techniques in Go. Please check the individual files for more details.")
	fmt.Println("Read-Write Locks:")
	ReadWriteLock()
	fmt.Println("Semaphore:")
	Semaphore()
	fmt.Println("Barrier:")
	BarrierExample()
	fmt.Println("Conditional Variables:")
	CondVariable()
	fmt.Println("Lock and waith free algo:")
	LockFree()
	fmt.Println("Done!")
}

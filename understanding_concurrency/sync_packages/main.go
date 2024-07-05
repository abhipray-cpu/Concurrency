package main

import "fmt"

func main() {
	fmt.Println("Mutex demo")
	MutexDemo()

	fmt.Println("\nRWMutex demo")
	InMemoryCache()

	fmt.Println("\nWaitGroup demo")
	WaitGroup()

	fmt.Println("\nOnce demo")
	ConfigIntializer()

	//fmt.Println("\nSelect demo")
	//Chat()

	//fmt.Println("\n Pool demo")
	//PoolHandler()

	//fmt.Println("\nMap demo")
	//SyncMap()

	fmt.Println("\nCond demo")
	CondDemo()
}

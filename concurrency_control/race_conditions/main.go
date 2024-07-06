/*
A race condition occurs in software when two or more threads or processes can access shared data and they try to change it at
the same time. Because the outcome depends on the non-deterministic ordering of their execution (the "race"),
the result of the operation is unpredictable and varies, leading to incorrect program behavior or software bugs.

Race conditions are common in concurrent programming when proper synchronization mechanisms (like locks, semaphores, barriers,
 etc.) are not used or are used incorrectly to coordinate access to shared resources.

Key characteristics of race conditions include:

1)Concurrency: They occur in a concurrent environment where multiple execution sequences interact.
2)Shared Resources: They involve access to shared resources (e.g., variables, files).
3)Modification: At least one of the threads/processes attempts to modify the shared resource.
4)Non-determinism: The outcome depends on the relative timing of execution between threads/processes.
5)Race conditions can lead to various issues, including data corruption, unpredictable behavior, crashes,
and security vulnerabilities. Identifying and fixing race conditions often requires careful analysis and testing,
typically involving tools like race detectors or employing synchronization techniques to ensure that only one thread
can access the shared resource at a time or that operations are performed in a controlled sequence.

start
*Product

*/

package main

import (
	"fmt"
	"sync"
)

// BankAccount represents a bank account with a balance
type BankAccount struct {
	balance int
}

// Deposit increases the account balance by the given amount
func (a *BankAccount) Deposit(amount int) {
	a.balance += amount
}

// Withdraw decreases the account balance by the given amount if sufficient funds are available
func (a *BankAccount) Withdraw(amount int) {
	if a.balance >= amount {
		a.balance -= amount
	}
}

func main() {
	account := BankAccount{balance: 100}
	var wg sync.WaitGroup

	// Simulate concurrent deposits and withdrawals
	for i := 0; i < 100; i++ {
		wg.Add(2)
		go func() {
			defer wg.Done()
			account.Deposit(10)
		}()
		go func() {
			defer wg.Done()
			account.Withdraw(5)
		}()
	}

	wg.Wait()
	fmt.Printf("Final account balance: $%d\n", account.balance)
}

// use this command to test racr condition go run -race main.go

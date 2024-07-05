package main

import (
	"fmt"
	"sync"
)

/*
It is used to ensure that only one goroutine can access a particular section of code or data at a time,
which is crucial for preventing race conditions in concurrent programming.
*/

// lets create a bank account simulation using mutex

// BankAccount struct holds the balance and a mutex to lock the balance
type BankAccount struct {
	balance int
	mutex   sync.Mutex
}

// Deposit method to safely deposit money into the account
func (a *BankAccount) Deposit(amount int) {
	a.mutex.Lock() // Lock the mutex to prevent any other operation from modifying the balance
	fmt.Println("Depositing", amount, "to account")
	a.balance += amount // Deposit the amount
	a.mutex.Unlock()    // Unlock the mutex
}

// Withdraw money safely from the account
func (a *BankAccount) Withdraw(amount int) {
	a.mutex.Lock() // Lock the mutex to prevent any other operation from modifying the balance
	fmt.Println("Withdrawing", amount, "from account")
	a.balance -= amount // Withdraw the amount
	a.mutex.Unlock()    // Unlock the mutex
}

func MutexDemo() {
	account := BankAccount{balance: 1000} // Create a new bank account with a balance of 1000
	var wg sync.WaitGroup

	// simulate concurrent deposit
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(amount int) {
			defer wg.Done()
			account.Deposit(amount)
		}(100)
	}

	// simulate concurrent withdraw
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(amount int) {
			defer wg.Done()
			account.Withdraw(amount)
		}(69)
	}
	wg.Wait() // Wait for all goroutines to finish
	fmt.Println("Final balance:", account.balance)
}

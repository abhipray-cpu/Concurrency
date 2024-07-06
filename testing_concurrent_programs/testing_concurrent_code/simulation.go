package main

import (
	"sync"
)

// Account represents a bank account with a balance.
type Account struct {
	balance int
	lock    sync.Mutex
}

// NewAccount creates a new Account with an initial balance.
func NewAccount(initialBalance int) *Account {
	return &Account{balance: initialBalance}
}

// Deposit increases the account balance by the given amount.
func (a *Account) Deposit(amount int) {
	a.lock.Lock()
	defer a.lock.Unlock()
	a.balance += amount
}

// Withdraw decreases the account balance by the given amount if sufficient funds exist.
func (a *Account) Withdraw(amount int) bool {
	a.lock.Lock()
	defer a.lock.Unlock()
	if a.balance < amount {
		return false // Insufficient funds
	}
	a.balance -= amount
	return true
}

// Balance returns the current balance of the account.
func (a *Account) Balance() int {
	a.lock.Lock()
	defer a.lock.Unlock()
	return a.balance
}

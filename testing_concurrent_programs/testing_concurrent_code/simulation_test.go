package main

/*
Simulate different scheduling orders or timing scenarios to see how they affect the program's correctness.
This can be done manually by inserting delays or using tools that control thread scheduling.
*/

import (
	"sync"
	"testing"
)

func TestAccountSimulation(t *testing.T) {
	account := NewAccount(1000) // Start with a balance of 1000

	var wg sync.WaitGroup
	// Simulate 100 deposits of 10
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			account.Deposit(10)
			wg.Done()
		}()
	}

	// Simulate 100 withdrawals of 5
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			account.Withdraw(5)
			wg.Done()
		}()
	}

	wg.Wait()

	finalBalance := account.Balance()
	expectedBalance := 1000 + (100 * 10) - (100 * 5)
	if finalBalance != expectedBalance {
		t.Errorf("Final balance got: %d, want: %d.", finalBalance, expectedBalance)
	}
}

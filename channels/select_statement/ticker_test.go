package main

import (
	"testing"
)

func TestNewTicker(t *testing.T) {
	symbol := "AAPL"
	ticker := NewTicker(symbol)
	if ticker.Symbol != symbol {
		t.Errorf("Expected symbol %s, got %s", symbol, ticker.Symbol)
	}
}

func TestTickerPriceUpdate(t *testing.T) {
	ticker := NewTicker("AAPL")
	initialPrice := ticker.Price
	// Simulate a price update
	priceUpdates := make(chan Ticker)
	go ticker.Start(priceUpdates)

	// Wait for the first price update
	updatedTicker := <-priceUpdates
	if updatedTicker.Price == initialPrice {
		t.Errorf("Expected the price to update from %f, but it stayed the same", initialPrice)
	}
}

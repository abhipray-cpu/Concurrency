package main

import (
	"math/rand"
	"time"
)

type Ticker struct {
	Symbol string
	Price  float64
}

func NewTicker(symbol string) *Ticker {
	return &Ticker{Symbol: symbol}
}

// Start starts the ticker and sends price updates to the specified channel.
// It simulates price changes by updating the ticker's price and sending it to the channel at random intervals.
func (t *Ticker) Start(priceUpdates chan<- Ticker) {
	for {
		// Simulate price change
		t.Price = 100 + rand.Float64()*10
		priceUpdates <- *t
		time.Sleep(time.Duration(rand.Intn(5)+1) * time.Second)
	}
}

package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Product represents the product entity with relevant fields
type Product struct {
	ProductID     int     `json:"productID"`
	ProductName   string  `json:"productName"`
	Category      string  `json:"category"`
	Price         float64 `json:"price"`
	StockQuantity int     `json:"stockQuantity"`
}

// generateMockProduct generates a mock product with random data
func generateMockProduct() Product {
	rand.Seed(time.Now().UnixNano())
	return Product{
		ProductID:     rand.Intn(1000),
		ProductName:   fmt.Sprintf("Product %d", rand.Intn(100)),
		Category:      []string{"Electronics", "Clothing", "Books"}[rand.Intn(3)],
		Price:         float64(rand.Intn(100)) + rand.Float64(),
		StockQuantity: rand.Intn(100),
	}
}

func ProductGenerator(productChannel chan Product) {
	ticker := time.NewTicker(2 * time.Microsecond)
	defer ticker.Stop()

	count := 0 // Counter for the number of ticks

	go func() {
		for range ticker.C {
			if count >= 1000 { // Check if the ticker has ticked 1000 times
				close(productChannel) // Close the channel to signal the receiver to stop
				return                // Stop the goroutine
			}
			product := generateMockProduct()
			productChannel <- product
			count++ // Increment the counter
		}
	}()

	for product := range productChannel {
		fmt.Printf("New product generated: %+v\n", product)
	}
}

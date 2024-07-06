package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

// Assuming User, Order, and Product structs are defined elsewhere

func loadToCSV(dataChannel <-chan DataWrapper) {
	fmt.Println("Loading data to CSV...")
	file, err := os.Create("final.csv")
	if err != nil {
		log.Fatalf("Failed to create file: %v", err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Define a more detailed header based on the combined data structure
	header := []string{"User ID", "User Name", "User Email", "User Age", "User Location", "Order ID", "Order User ID", "Order Product ID", "Order Quantity", "Order Date", "Product ID", "Product Name", "Product Category", "Product Price", "Product Stock Quantity"}
	if err := writer.Write(header); err != nil {
		log.Fatalf("Failed to write header to CSV: %v", err)
	}

	for data := range dataChannel {
		fmt.Println(data)
		record := []string{
			data.User.CustomerId,
			data.User.Name,
			data.User.Email,
			fmt.Sprintf("%d", data.User.Age),
			data.User.Location,
			fmt.Sprintf("%d", data.Order.OrderID),
			data.Order.UserID,
			fmt.Sprintf("%d", data.Order.ProductID),
			fmt.Sprintf("%d", data.Order.Quantity),
			data.Order.OrderDate,
			fmt.Sprintf("%d", data.Product.ProductID),
			data.Product.ProductName,
			data.Product.Category,
			fmt.Sprintf("%.2f", data.Product.Price),
			fmt.Sprintf("%d", data.Product.StockQuantity),
		}
		if err := writer.Write(record); err != nil {
			log.Fatalf("Failed to write data to CSV: %v", err)
		}
	}
}

package data

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"time"
)

// Order represents an order with OrderID, UserID (in a specific format), ProductID, Quantity, and OrderDate
type Order struct {
	OrderID   int
	UserID    string // Adjusted to string to accommodate the format "C001", "C002", etc.
	ProductID int
	Quantity  int
	OrderDate string
}

// generateOrders creates a slice of approximately 3000 Order instances with semi-realistic data
func generateOrders() []Order {
	totalUsers := 1000
	ordersPerUser := 3
	orders := make([]Order, totalUsers*ordersPerUser)
	orderID := 1

	for i := 0; i < totalUsers; i++ {
		userID := fmt.Sprintf("C%03d", i+1) // Format the UserID as specified
		for j := 0; j < ordersPerUser; j++ {
			orders[(i*ordersPerUser)+j] = Order{
				OrderID:   orderID,
				UserID:    userID,
				ProductID: (orderID % 50) + 1,                                      // Simulate 50 products
				Quantity:  (orderID % 5) + 1,                                       // Quantity between 1 and 5
				OrderDate: time.Now().AddDate(0, 0, -orderID).Format("2006-01-02"), // Different order dates
			}
			orderID++
		}
	}
	return orders
}

// writeOrdersToCSV writes the slice of Orders to a CSV file
func writeOrdersToCSV(orders []Order) error {
	// Check if the file exists and delete it if it does
	if _, err := os.Stat("./orders.csv"); err == nil {
		if err := os.Remove("./orders.csv"); err != nil {
			return fmt.Errorf("failed to remove existing file: %w", err)
		}
	}

	// Now, create a new file
	file, err := os.Create("./orders.csv")
	if err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Write the header
	if err := writer.Write([]string{"OrderID", "UserID", "ProductID", "Quantity", "OrderDate"}); err != nil {
		return fmt.Errorf("failed to write header: %w", err)
	}

	// Write the data
	for _, order := range orders {
		record := []string{
			strconv.Itoa(order.OrderID),
			order.UserID,
			strconv.Itoa(order.ProductID),
			strconv.Itoa(order.Quantity),
			order.OrderDate,
		}
		if err := writer.Write(record); err != nil {
			return fmt.Errorf("failed to write record: %w", err)
		}
	}

	return nil
}

func GenerateCSVData() {
	orders := generateOrders()
	if err := writeOrdersToCSV(orders); err != nil {
		fmt.Printf("Error writing orders to CSV: %s\n", err)
	}
	fmt.Println("Order Mock data generated successfully.")
}

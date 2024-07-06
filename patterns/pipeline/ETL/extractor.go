package main

import (
	"ETL/data"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"sync"
)

type DataWrapper struct {
	User    *data.User
	Order   *data.Order
	Product *Product
}

func extractUserFromJSON(out chan<- DataWrapper) {
	userData, err := ioutil.ReadFile("user.json")
	if err != nil {
		log.Fatalf("Error reading user.json file: %v", err)
	}
	var users []data.User
	if err := json.Unmarshal(userData, &users); err != nil {
		log.Fatalf("Error unmarshalling user data: %v", err)
	}
	for _, user := range users {
		out <- DataWrapper{User: &user, Order: &data.Order{
			OrderID:   89,
			UserID:    "",
			ProductID: 12,
			Quantity:  69,
			OrderDate: "2021-06-09",
		},
			Product: &Product{
				ProductID:     12,
				ProductName:   "Hash",
				Category:      "Food",
				Price:         69.69,
				StockQuantity: 12,
			}}
	}
}

func extractOrdersFromCSV(out chan<- DataWrapper) {
	file, err := os.Open("orders.csv")
	if err != nil {
		log.Fatalf("Error reading orders.csv file: %v", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	isFirstLine := true
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break // End of file
		}
		if err != nil {
			log.Fatalf("Error reading CSV record: %v", err)
		}

		// Skip the header line
		if isFirstLine {
			isFirstLine = false
			continue
		}

		// Assuming the CSV format matches: OrderID, UserID, ProductID, Quantity, OrderDate
		// Convert record to DataWrapper instance (example conversion, adjust according to actual structure)
		orderID, _ := strconv.Atoi(record[0])
		userID := record[1]
		productID, _ := strconv.Atoi(record[2])
		quantity, _ := strconv.Atoi(record[3])
		orderDate := record[4]

		dataWrapper := DataWrapper{
			User: &data.User{
				CustomerId: "123",
				Name:       "John Doe",
				Email:      "test@gmail.com",
				Age:        25,
				Location:   "USA",
			},
			Order: &data.Order{
				OrderID:   orderID,
				UserID:    userID,
				ProductID: productID,
				Quantity:  quantity,
				OrderDate: orderDate,
			},
			Product: &Product{
				ProductID:     12,
				ProductName:   "Hash",
				Category:      "Food",
				Price:         69.69,
				StockQuantity: 12,
			},
		}

		// Send the constructed DataWrapper to the channel
		out <- dataWrapper
	}
	close(out) // Close the channel when done
}

func listForProducts(productChannel <-chan Product, out chan<- DataWrapper) {
	for products := range productChannel {
		out <- DataWrapper{User: &data.User{
			CustomerId: "123",
			Name:       "John Doe",
			Email:      "test@gmail.com",
			Age:        25,
			Location:   "USA",
		}, Order: &data.Order{
			OrderID:   89,
			UserID:    "",
			ProductID: 12,
			Quantity:  69,
			OrderDate: "2021-06-09",
		}, Product: &products}
	}
}

func Extract(productChannel <-chan Product) <-chan DataWrapper {
	fmt.Println("Extracting data...")
	var pipeline = make(chan DataWrapper)
	var wg sync.WaitGroup // Used to wait for all goroutines to finish

	wg.Add(1)
	go func() {
		defer wg.Done() // Decrement the counter when the goroutine completes
		extractUserFromJSON(pipeline)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		extractOrdersFromCSV(pipeline)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		listForProducts(productChannel, pipeline)
	}()

	go func() {
		wg.Wait()       // Wait for all goroutines to finish
		close(pipeline) // Close the channel once all data has been extracted
	}()

	return pipeline
}

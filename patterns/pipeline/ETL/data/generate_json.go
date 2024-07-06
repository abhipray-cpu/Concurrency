package data

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"time"
)

type User struct {
	CustomerId string `json:"customerId"`
	Name       string `json:"name"`
	Email      string `json:"email"`
	Age        int    `json:"age"`
	Location   string `json:"location"`
}

// GenerateData generates a JSON file with random user data
func RemoveExistingData() error {
	_, err := os.Stat("./user.json")
	if os.IsNotExist(err) {
		return nil // File does not exist
	}
	if err != nil {
		return fmt.Errorf("error checking if file exists: %v", err)
	}

	err = os.Remove("./user.json")
	if err != nil {
		return fmt.Errorf("error removing file: %v", err)
	}
	fmt.Println("Existing data removed")
	return nil
}

// GenerateMockJsonData generates a JSON file with random user data
func GenerateMockJsonData() {
	err := RemoveExistingData()
	if err != nil {
		fmt.Println(err)
		return
	}
	rand.Seed(time.Now().UnixNano())
	users := make([]User, 1000)
	for i := 0; i < 1000; i++ {
		users[i] = User{
			CustomerId: fmt.Sprintf("C%03d", i+1),
			Name:       fmt.Sprintf("Name%d", i+1),
			Email:      fmt.Sprintf("user%d@example.com", i+1),
			Age:        rand.Intn(65) + 18, // Random age between 18 and 82
			Location:   fmt.Sprintf("Location%d", rand.Intn(10)+1),
		}
	}
	file, err := json.MarshalIndent(users, "", "  ")
	if err != nil {
		fmt.Println("Error marshalling JSON:", err)
		return
	}

	err = os.WriteFile("./user.json", file, 0644)
	if err != nil {
		fmt.Println("Error writing file:", err)
		return
	}

	fmt.Println("User Mock data generated successfully.")
}

package main

import (
	"fmt"
	"io/ioutil"
)

// ProcessData simulates processing input data and returns the processed data.
func ProcessData(input string) (string, error) {
	// Simulate data processing
	processedData := fmt.Sprintf("Processed: %s", input)
	return processedData, nil
}

// SaveData saves the processed data to a file.
func SaveData(data, filename string) error {
	// Write the data to a file
	return ioutil.WriteFile(filename, []byte(data), 0644)
}

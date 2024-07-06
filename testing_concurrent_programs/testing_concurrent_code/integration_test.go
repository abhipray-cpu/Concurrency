package main

/*
Test the integration points between concurrent components to ensure they interact correctly under concurrent execution.
Verify that shared resources are accessed and modified correctly.
*/

import (
	"os"
	"strings"
	"testing"
)

// TestProcessAndSaveDataIntegration tests the entire flow of processing and saving data.
func TestProcessAndSaveDataIntegration(t *testing.T) {
	input := "test data"
	expectedSubstring := "Processed: test data"
	filename := "test_output.txt"

	// Process the data
	processedData, err := ProcessData(input)
	if err != nil {
		t.Fatalf("ProcessData returned an error: %v", err)
	}

	// Save the processed data
	err = SaveData(processedData, filename)
	if err != nil {
		t.Fatalf("SaveData returned an error: %v", err)
	}

	// Clean up the file after the test
	defer os.Remove(filename)

	// Read back the data to verify
	data, err := os.ReadFile(filename)
	if err != nil {
		t.Fatalf("Error reading back the saved file: %v", err)
	}

	if !strings.Contains(string(data), expectedSubstring) {
		t.Errorf("The saved file does not contain the expected substring '%s'", expectedSubstring)
	}
}

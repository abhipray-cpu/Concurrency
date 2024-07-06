package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

func DatGenerator() {
	rand.Seed(time.Now().UnixNano())

	file_path := "large_dataset.txt"
	numLines := 10000000

	// create and open the file

	file, err := os.Create(file_path)
	if err != nil {
		fmt.Println("Error creating file", err)
		return
	}
	defer file.Close()

	// using  bufio for efficient writing
	writer := bufio.NewWriter(file)

	for i := 0; i < numLines; i++ {
		// generate random number
		num := rand.Intn(100000) + 1
		// write to file
		_, err := writer.WriteString(fmt.Sprintf("%d\n", num))
		if err != nil {
			fmt.Println("Error writing to file", err)
			return
		}
	}
	// Flush any buffered data to the file
	if err := writer.Flush(); err != nil {
		fmt.Printf("Error flushing to file: %v\n", err)
	}

	fmt.Println("Dataset generation complete.")
}

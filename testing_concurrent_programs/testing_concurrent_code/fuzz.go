package main

// ProcessInput takes a string input and performs some processing.
// For demonstration, it just checks if the input is "hello".
func ProcessInput(input string) bool {
	return input == "hello"
}

// command to test go test -fuzz=Fuzz

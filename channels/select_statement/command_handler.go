package main

import (
	"bufio"
	"os"
)

// listenForCommands listens for user input from the standard input and sends the input as commands to the commandChannel.
func listenForCommands(commandChannel chan<- string) {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		commandChannel <- scanner.Text()
	}
}

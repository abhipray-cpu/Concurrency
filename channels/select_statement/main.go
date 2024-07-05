package main

import (
	"fmt"
	"strings"
)

// main is the entry point of the program.
// It sets up channels for receiving price updates and commands,
// and maintains a watchlist of tickers.
// It listens for commands and performs actions based on the command received.
// If the command is "add", it adds a new ticker to the watchlist and starts receiving price updates for it.
// If the command is "remove", it removes the ticker from the watchlist.
// It also prints the new price for tickers in the watchlist whenever a price update is received.
func main() {
	priceUpdates := make(chan Ticker)
	commandChannel := make(chan string)
	watchlist := make(map[string]*Ticker)

	go listenForCommands(commandChannel)

	for {
		select {
		case command := <-commandChannel:
			args := strings.Split(command, " ")
			if len(args) != 2 {
				fmt.Println("Invalid command. Usage: add <symbol> or remove <symbol>")
				continue
			}
			action, symbol := args[0], args[1]

			switch action {
			case "add":
				if _, exists := watchlist[symbol]; !exists {
					ticker := NewTicker(symbol)
					watchlist[symbol] = ticker
					go ticker.Start(priceUpdates)
					fmt.Printf("Added %s to watchlist\n", symbol)
				}
			case "remove":
				if ticker, exists := watchlist[symbol]; exists {
					delete(watchlist, symbol)
					// In a real application, you would signal the ticker to stop.
					fmt.Printf("Removed %s from watchlist\n", ticker.Symbol)
				}
			default:
				fmt.Println("Unknown command:", action)
			}
		case update := <-priceUpdates:
			fmt.Printf("New price for %s: $%.2f\n", update.Symbol, update.Price)
		}
	}
}

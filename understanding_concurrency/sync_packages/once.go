package main

/*
The sync.Once type in Go, part of the sync package, provides a simple, thread-safe way to execute a piece of code
exactly once, regardless of how many times it is called. This is particularly useful for initialization code that
must run only once, even in the presence of multiple goroutines.

Key Features of sync.Once
1)Thread-Safety: Ensures that the function passed to Do is executed only once, even when called from multiple
goroutines concurrently.

2)Idempotency: Guarantees that the function will be executed no more than once, making it ideal for
initialization routines or singleton object creation.

3)Efficiency: After the function has been executed once, subsequent calls to Do with the same Once
instance have very little overhead, making it efficient.
*/

// an example would be config loader

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"sync"
)

// Config represents the application configuration
type Config struct {
	ServerAddress string `json:"serverAddress"`
	LogLevel      string `json:"logLevel"`
}

var (
	instance *Config
	once     sync.Once
)

// LoadConfig reads configuration from a file and returns a Config instance
func LoadConfig(filename string) *Config {
	once.Do(func() {
		bytes, err := ioutil.ReadFile(filename)
		if err != nil {
			panic(fmt.Sprintf("Unable to read config file: %v", err))
		}

		config := &Config{}
		err = json.Unmarshal(bytes, config)
		if err != nil {
			panic(fmt.Sprintf("Unable to parse config: %v", err))
		}

		instance = config
	})
	fmt.Println("Returning config instance")
	return instance
}

func ConfigIntializer() {
	// Example usage
	config := LoadConfig("config.json")
	fmt.Printf("Loaded config: %+v\n", config)

	// Attempt to load a different config to demonstrate that the original is retained
	anotherConfig := LoadConfig("another_config.json")
	fmt.Printf("Loaded config: %+v\n", anotherConfig)
}

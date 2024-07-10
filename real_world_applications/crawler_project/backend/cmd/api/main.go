package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"sync"
	"time"
	"web_crawler/models"
	"web_crawler/pkg"
	"web_crawler/utils"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Config struct {
	Model     *models.Models
	Logger    *logrus.Logger
	Scheduler *pkg.Scheduler
}

var seedUrls []string

func init() {
	file, err := os.Open("../../Seed.txt")
	if err != nil {
		fmt.Printf("error opening file: %s", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		seedUrls = append(seedUrls, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Printf("error reading file: %s", err)
	}
}

func main() {
	// Load the .env file
	if err := godotenv.Load("../../.env"); err != nil {
		fmt.Printf("Error loading .env file: %v", err)
		os.Exit(1)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Minute)
	defer cancel()

	esClient, err := initESClient()
	if err != nil {
		fmt.Printf("Error initializing ES client: %v", err)
		os.Exit(1)
	}
	mongClient, err := connectToMongo()
	if err != nil {
		fmt.Printf("Error initializing mongo client: %v", err)
		os.Exit(1)
	}

	models := models.NewModels(esClient, mongClient)
	logger := utils.NewLogger()
	e := echo.New()
	app := &Config{
		Model:  models,
		Logger: logger,
	}
	app.routes(e)

	// Initialize the scheduler with 10 workers
	sched := pkg.NewScheduler(10)
	app.Scheduler = sched

	// Initialize a map to keep track of URLs and a mutex for synchronization
	var urlsSeen = make(map[string]bool)
	var mutex sync.Mutex

	// Define the worker function
	workerFunc := func(job interface{}) {
		url, ok := job.(string)
		if !ok {
			fmt.Println("Invalid job type")
			return
		}

		// Fetch the content
		htmlContent, err := pkg.Fetch(url)
		if err != nil {
			fmt.Printf("Error fetching URL %s: %v\n", url, err)
			return
		}

		// Parse the content
		urls, contents, title, description, keywords, err := pkg.Parse(htmlContent)
		if err != nil {
			fmt.Printf("Error parsing content from %s: %v\n", url, err)
			return
		}
		// Store the parsed data
		_, err = pkg.InsertCombinedContent(ctx, url, 200, contents, title, description, keywords)
		if err != nil {
			fmt.Printf("Error storing content from %s: %v\n", url, err)
			return
		}

		// Optionally, submit new URLs to the scheduler
		mutex.Lock()
		for _, newURL := range urls {
			if _, seen := urlsSeen[newURL]; !seen {
				urlsSeen[newURL] = true
				sched.Submit(newURL)
			}
		}
		mutex.Unlock()
	}

	// Start the scheduler
	sched.Start(ctx, workerFunc)
	mutex.Lock()
	for _, url := range seedUrls {
		if _, seen := urlsSeen[url]; !seen {
			fmt.Printf("Submitting seed URL: %s\n", url)
			urlsSeen[url] = true
			sched.Submit(url)
		}
	}
	mutex.Unlock()

	go func() {
		<-ctx.Done()
		sched.Stop()
	}()

	initialWaitTime := 1 * time.Second
	maxRetries := 5
	factor := 2

	for attempt := 1; attempt <= maxRetries; attempt++ {
		err := e.Start(":8082")
		if err != nil {
			log.Printf("Attempt %d: server failed to start: %v", attempt, err)
			if attempt == maxRetries {
				log.Fatalf("Server failed to start after %d attempts", maxRetries)
			}
			time.Sleep(initialWaitTime)
			initialWaitTime *= time.Duration(factor)
		} else {
			break
		}
	}

}

func initESClient() (*elasticsearch.Client, error) {
	esHost := os.Getenv("ELASTICSEARCH_HOST")
	esPort := os.Getenv("ELASTICSEARCH_PORT")
	esAddress := fmt.Sprintf("http://%s:%s", esHost, esPort)

	esConfig := elasticsearch.Config{
		Addresses: []string{esAddress},
	}

	esClient, err := elasticsearch.NewClient(esConfig)
	if err != nil {
		return nil, fmt.Errorf("Failed to create ES client: %v", err)
	}
	fmt.Println("Connected to Elasticsearch")
	return esClient, nil
}

// connectToMongo connects to the MongoDB instance and returns the client.
func connectToMongo() (*mongo.Client, error) {
	mongoURI := os.Getenv("MONGOURL")
	mongoUser := os.Getenv("MONGOUSER")
	mongoPassword := os.Getenv("MONGOPASSWORD")

	// Check if the environment variables are loaded
	if mongoURI == "" {
		log.Panic("MONGOURL environment variable is not set")
		return nil, fmt.Errorf("MONGOURL environment variable is not set")
	}
	if mongoUser == "" {
		log.Panic("MONGOUSER environment variable is not set")
		return nil, fmt.Errorf("MONGOUSER environment variable is not set")
	}
	if mongoPassword == "" {
		log.Panic("MONGOPASSWORD environment variable is not set")
		return nil, fmt.Errorf("MONGOPASSWORD environment variable is not set")
	}

	clientOptions := options.Client().ApplyURI(mongoURI)
	clientOptions.SetAuth(options.Credential{
		Username: mongoUser,
		Password: mongoPassword,
	})

	c, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Panic(err)
		return nil, err
	}
	log.Println("Connected to MongoDB")
	return c, nil
}

package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"net/url"
	"os"
	"sync"
	"time"
	"web_crawler/models"
	"web_crawler/pkg"
	"web_crawler/utils"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/writeconcern"
)

type Config struct {
	Model     *models.Models
	Logger    *logrus.Logger
	Scheduler *pkg.Scheduler
}

var initialUrls []string

//Defining a channel for seed urls

var seedUrls = make(chan string, 5000)

func init() {
	file, err := os.Open("/app/Seed.txt")
	if err != nil {
		fmt.Printf("error opening file: %s", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		initialUrls = append(initialUrls, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Printf("error reading file: %s", err)
	}
	for _, url := range initialUrls { // Assuming initialUrls is a slice of your seed URLs
		seedUrls <- url
	}
}

func main() {
	// Load the .env file
	// if err := godotenv.Load("../../.env"); err != nil {
	// 	fmt.Printf("Error loading .env file: %v", err)
	// 	os.Exit(1)
	// }
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
	// Configure CORS middleware
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},                                        // Allows all origins
		AllowMethods: []string{echo.GET, echo.PUT, echo.POST, echo.DELETE}, // Specify allowed methods
	}))
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
		Url, ok := job.(string)
		if !ok {
			fmt.Println("Invalid job type")
			return
		}
		// Fetch the content
		htmlContent, err := pkg.Fetch(Url)
		if err != nil {
			app.Logger.Error(fmt.Sprintf("Error fetching content from %s: %v", Url, err))
			return
		}
		// Parse the content
		urls, contents, title, description, keywords, err := pkg.Parse(htmlContent)
		if err != nil {
			app.Logger.Error(fmt.Sprintf("Error parsing content from %s: %v", Url, err))
			return
		}
		// Submit new URLs to the scheduler
		for _, newURL := range urls {
			u, err := url.Parse(newURL)
			if err == nil && (u.Scheme == "http" || u.Scheme == "https") {
				seedUrls <- newURL
			}
		}
		// Store the parsed data
		_, err = pkg.InsertCombinedContent(ctx, Url, 200, contents, title, description, keywords)
		if err != nil {
			app.Logger.Error(fmt.Sprintf("Error inserting content from %s: %v", Url, err))
			return
		}
	}
	go func() {
		initialWaitTime := 1 * time.Second
		maxRetries := 5
		factor := 2

		for attempt := 1; attempt <= maxRetries; attempt++ {
			err := e.Start(":8081")
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
	}()
	// Start the scheduler
	sched.Start(ctx, workerFunc)
	mutex.Lock()
	for url := range seedUrls {

		if _, seen := urlsSeen[url]; !seen {
			urlsSeen[url] = true
			fmt.Println("Submitting URL: ", url)
		}
	}
	mutex.Unlock()

	go func() {
		<-ctx.Done()
		sched.Stop()
		close(seedUrls)
	}()

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

	clientOptions := options.Client().ApplyURI(mongoURI).SetAuth(options.Credential{
		Username: mongoUser,
		Password: mongoPassword,
	}).SetWriteConcern(writeconcern.New(writeconcern.WMajority()))

	c, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Panic(err)
		return nil, err
	}
	log.Println("Connected to MongoDB")

	// Ensure unique index on the email field of the users collection
	usersCollection := c.Database("users").Collection("users") // Replace "yourDatabaseName" with your actual database name
	indexModel := mongo.IndexModel{
		Keys:    bson.D{{Key: "email", Value: 1}}, // Index key
		Options: options.Index().SetUnique(true),  // Ensure uniqueness
	}
	_, err = usersCollection.Indexes().CreateOne(context.TODO(), indexModel)
	if err != nil {
		log.Panic("Failed to create unique index on email field:", err)
		return nil, err
	}
	log.Println("Unique index on email field created successfully")

	return c, nil
}

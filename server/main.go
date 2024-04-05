package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Registry struct {
	Slug string `bson:"slug"`
	URL  string `bson:"url"`
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

  	db_url := os.Getenv("DB_URL")
	// MongoDB connection options
	clientOptions := options.Client().ApplyURI(db_url)

	// Connect to MongoDB
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := client.Disconnect(context.Background()); err != nil {
			log.Fatalf("Error disconnecting from MongoDB: %v", err)
		}
	}()
	if err := client.Ping(context.Background(), nil); err != nil {
		log.Fatalf("Error pinging MongoDB: %v", err)
	}
	fmt.Println("Connected to MongoDB!")

	// Access the "registry" collection
	registryCollection := client.Database("xoly").Collection("registry")

	// Initialize Fiber app
	app := fiber.New()

	// Define route for handling shortened URLs
	app.Get("/:param", func(c *fiber.Ctx) error {
		param := c.Params("param")

		var result Registry
		filter := bson.M{"slug": param}
		err := registryCollection.FindOne(context.Background(), filter).Decode(&result)
		if err != nil {
			if err == mongo.ErrNoDocuments {
				// Handle not found error
				return c.Status(http.StatusNotFound).SendString("Not Found")
			}
			log.Fatal(err)
			return err
		}

		// Redirect to the URL
		return c.Redirect(result.URL)
	})

	// Start Fiber server
	go func() {
		if err := app.Listen(":3000"); err != nil {
			log.Fatalf("Error starting Fiber server: %v", err)
		}
	}()

	// Graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit

	// Give some time for existing connections to finish
	_, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	
	defer cancel()

	if err := app.Shutdown(); err != nil {
		log.Fatalf("Error shutting down Fiber server: %v", err)
	}
	fmt.Println("Server gracefully stopped")
}

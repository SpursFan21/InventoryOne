package main

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Initialize MongoDB connection
func init() {
	var err error
	dbClient, err = createMongoClient()
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}

	// Check connection
	if err = dbClient.Ping(context.Background(), nil); err != nil {
		log.Fatalf("Failed to ping MongoDB: %v", err)
	}
}

// Create and return a new MongoDB client
func createMongoClient() (*mongo.Client, error) {
	clientOptions := options.Client().ApplyURI("mongodb://mongo:27017")
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		return nil, err
	}
	return client, nil
}

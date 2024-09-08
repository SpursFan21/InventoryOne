package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type User struct {
	ID       string `json:"id,omitempty"`
	Username string `json:"username"`
	Role     string `json:"role"` // Can be "admin" or "user"
	Password string `json:"password,omitempty"`
}

// MongoDB collection for users
var userCollection *mongo.Collection

func init() {
	// Initialize MongoDB collection
	clientOptions := options.Client().ApplyURI("mongodb://mongo:27017")
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}

	userCollection = client.Database("inventoryone").Collection("users")
}

// Create user (admin only)
func createUser(w http.ResponseWriter, r *http.Request) {
	var user User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	// Check if user already exists
	filter := bson.M{"username": user.Username}
	err := userCollection.FindOne(context.Background(), filter).Decode(&user)
	if err == nil {
		http.Error(w, "User already exists", http.StatusConflict)
		return
	}

	// Insert user into database
	_, err = userCollection.InsertOne(context.Background(), user)
	if err != nil {
		http.Error(w, "Failed to create user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}

// Get all users
func getUsers(w http.ResponseWriter, r *http.Request) {
	var users []User
	cursor, err := userCollection.Find(context.Background(), bson.D{})
	if err != nil {
		http.Error(w, "Failed to get users", http.StatusInternalServerError)
		return
	}
	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		var user User
		if err := cursor.Decode(&user); err != nil {
			http.Error(w, "Error reading user", http.StatusInternalServerError)
			return
		}
		users = append(users, user)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

// Update a user (admin only)
func updateUser(w http.ResponseWriter, r *http.Request) {
	// Extract user ID from URL and handle update logic
}

// Delete a user (admin only)
func deleteUser(w http.ResponseWriter, r *http.Request) {
	// Extract user ID from URL and handle deletion logic
}

// Get a single user by ID
func getUser(w http.ResponseWriter, r *http.Request) {
	// Extract user ID from URL and return user info
}

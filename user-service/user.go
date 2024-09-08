package main

import (
	"context"
	"fmt"
	"log"

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
	clientOptions := options.Client().ApplyURI("mongodb://mongo:27017")
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}
	userCollection = client.Database("inventoryone").Collection("users")
}

func createUser(user User) error {
	filter := bson.M{"username": user.Username}
	err := userCollection.FindOne(context.Background(), filter).Decode(&user)
	if err == nil {
		return fmt.Errorf("User already exists")
	}

	_, err = userCollection.InsertOne(context.Background(), user)
	return err
}

func getUsers() ([]User, error) {
	var users []User
	cursor, err := userCollection.Find(context.Background(), bson.D{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		var user User
		if err := cursor.Decode(&user); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

// Other user-related functions (update, delete, get single user) go here

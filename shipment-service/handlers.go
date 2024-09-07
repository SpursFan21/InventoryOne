package main

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var shipmentsCollection *mongo.Collection

func init() {
	// Initialize MongoDB collection
	db := dbClient.Database("shipmentdb")
	shipmentsCollection = db.Collection("shipments")
}

// Create a new shipment
func shipmentsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		var shipment Shipment
		if err := json.NewDecoder(r.Body).Decode(&shipment); err != nil {
			http.Error(w, "Invalid request", http.StatusBadRequest)
			return
		}

		_, err := shipmentsCollection.InsertOne(context.Background(), shipment)
		if err != nil {
			http.Error(w, "Failed to save shipment", http.StatusInternalServerError)
			return
		}

		if err := sendMessage(producer, "shipment-status-updated", shipment); err != nil {
			http.Error(w, "Failed to send message", http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte("Shipment created"))
		return
	}
	http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
}

// Track shipment status
func shipmentHandler(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/shipments/")
	var shipment Shipment
	filter := bson.M{"id": id}
	err := shipmentsCollection.FindOne(context.Background(), filter).Decode(&shipment)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			http.Error(w, "Shipment not found", http.StatusNotFound)
		} else {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
		}
		return
	}
	if r.Method == http.MethodGet {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(shipment)
		return
	}
	http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
}

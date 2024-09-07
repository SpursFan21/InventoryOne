package main

// Shipment represents a shipment with an ID and status.
type Shipment struct {
	ID     string `json:"id" bson:"id"`
	Status string `json:"status" bson:"status"`
}

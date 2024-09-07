package main

import (
	"context"
	"log"
	"net/http"

	"github.com/IBM/sarama"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	producer sarama.SyncProducer
	dbClient *mongo.Client
)

func main() {
	var err error
	producer, err = createKafkaProducer()
	if err != nil {
		log.Fatalf("Failed to create Kafka producer: %v", err)
	}
	defer producer.Close()

	dbClient, err = createMongoClient()
	if err != nil {
		log.Fatalf("Failed to create MongoDB client: %v", err)
	}
	defer func() {
		if err := dbClient.Disconnect(context.Background()); err != nil {
			log.Fatalf("Failed to disconnect MongoDB client: %v", err)
		}
	}()

	http.HandleFunc("/shipments", shipmentsHandler)
	http.HandleFunc("/shipments/", shipmentHandler)

	log.Println("Shipment Service is running on port 8081")
	if err := http.ListenAndServe(":8081", nil); err != nil {
		log.Fatal(err)
	}
}

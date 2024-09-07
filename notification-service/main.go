package main

import (
	"log"
	"net/http"

	"github.com/IBM/sarama"
)

// Global Kafka producer
var producer sarama.SyncProducer

func main() {
	// Initialize Kafka producer
	var err error
	producer, err = createKafkaProducer()
	if err != nil {
		log.Fatalf("Failed to create Kafka producer: %v", err)
	}
	defer producer.Close()

	// Set up HTTP routes
	http.HandleFunc("/notify", sendNotificationHandler)

	log.Println("Notification Service is running on port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

func createKafkaProducer() (sarama.SyncProducer, error) {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Partitioner = sarama.NewRoundRobinPartitioner
	config.Producer.Return.Successes = true

	producer, err := sarama.NewSyncProducer([]string{"kafka:9092"}, config)
	if err != nil {
		return nil, err
	}

	return producer, nil
}

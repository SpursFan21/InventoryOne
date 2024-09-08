package main

import (
	"context"
	"log"
	"time"

	"github.com/segmentio/kafka-go"
)

var kafkaWriter *kafka.Writer

func init() {
	var err error
	kafkaWriter = &kafka.Writer{
		Addr:     kafka.TCP("kafka:9092"), // Update with the actual Kafka broker address
		Topic:    "user-events",
		Balancer: &kafka.LeastBytes{},
	}

	// Check connectivity to Kafka broker
	_, err = kafkaWriter.WriteMessages(context.Background(), kafka.Message{
		Value: []byte("test connection"),
	})
	if err != nil {
		log.Fatalf("Error connecting to Kafka: %v", err)
	}
}

// Publish user-created event to Kafka
func publishUserEvent(user User, eventType string) {
	msg := kafka.Message{
		Key:   []byte(user.ID), // Ensure user.ID is a string; convert if needed
		Value: []byte(eventType),
		Time:  time.Now(),
	}

	err := kafkaWriter.WriteMessages(context.Background(), msg)
	if err != nil {
		log.Printf("Error publishing user event: %v", err)
	}
}

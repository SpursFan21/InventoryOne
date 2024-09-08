package main

import (
	"context"
	"log"

	"github.com/segmentio/kafka-go"
)

func SendKafkaMessage(topic, message string) error {
	writer := kafka.NewWriter(kafka.WriterConfig{
		Brokers:  []string{"localhost:9092"},
		Topic:    topic,
		Balancer: &kafka.LeastBytes{},
	})

	err := writer.WriteMessages(context.Background(),
		kafka.Message{
			Value: []byte(message),
		},
	)

	if err != nil {
		log.Println("Error sending message to Kafka:", err)
		return err
	}

	log.Println("Message sent to Kafka:", message)
	return nil
}

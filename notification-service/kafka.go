package main

import (
	"github.com/IBM/sarama"
)

// Function to send a message to Kafka using the global producer
func sendMessage(message string) error {
	msg := &sarama.ProducerMessage{
		Topic: "notifications",
		Value: sarama.StringEncoder(message),
	}

	_, _, err := producer.SendMessage(msg)
	if err != nil {
		return err
	}
	return nil
}

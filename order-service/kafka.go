package main

import (
	"encoding/json"

	"github.com/IBM/sarama"
)

// Function to send a message to Kafka
func sendMessage(order OrderRequest) error {
	config := sarama.NewConfig()
	producer, err := sarama.NewSyncProducer([]string{"kafka:9092"}, config)
	if err != nil {
		return err
	}
	defer producer.Close()

	msgBytes, err := json.Marshal(order)
	if err != nil {
		return err
	}

	msg := &sarama.ProducerMessage{
		Topic: "orders",
		Value: sarama.StringEncoder(msgBytes),
	}

	_, _, err = producer.SendMessage(msg)
	if err != nil {
		return err
	}
	return nil
}

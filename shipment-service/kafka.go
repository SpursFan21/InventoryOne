package main

import (
	"github.com/IBM/sarama"
)

// Create a new Kafka producer
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

// Send a message to Kafka
func sendMessage(producer sarama.SyncProducer, topic string, shipment Shipment) error {
	msg := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.StringEncoder(shipment.ID + " " + shipment.Status),
	}

	_, _, err := producer.SendMessage(msg)
	if err != nil {
		return err
	}
	return nil
}

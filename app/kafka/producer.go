package kafka

import (
	"context"
	"encoding/json"
	"log"
	"os"

	"github.com/segmentio/kafka-go"
)

// getBrokerAddress retrieves the Kafka broker address from environment variables
func getBrokerAddress() string {
	broker := os.Getenv("KAFKA_BROKER_ADDRESS")
	if broker == "" {
		log.Fatal("KAFKA_BROKER_ADDRESS not set")
	}
	return broker
}

// ProduceEvent sends a message to a Kafka topic.
func ProduceEvent(topic string, message interface{}) error {
	writer := kafka.NewWriter(kafka.WriterConfig{
		Brokers:  []string{getBrokerAddress()},
		Topic:    topic,
		Balancer: &kafka.LeastBytes{},
	})
	defer writer.Close()

	data, err := json.Marshal(message)
	if err != nil {
		return err
	}

	err = writer.WriteMessages(context.Background(),
		kafka.Message{
			Key:   []byte("key"),
			Value: data,
		},
	)
	if err != nil {
		log.Println("Error producing event: ", err)
		return err
	}

	log.Println("Message produced successfully")
	return nil
}

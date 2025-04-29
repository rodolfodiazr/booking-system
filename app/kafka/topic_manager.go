package kafka

import (
	"context"
	"log"
	"time"

	kafka "github.com/segmentio/kafka-go"
)

// CreateTopic creates a Kafka topic if it doesn't exist.
func CreateTopic(topicName string) {
	maxAttempts := 10

	var conn *kafka.Conn
	var err error

	for i := 1; i <= maxAttempts; i++ {
		log.Printf("Attempt %d: connecting to Kafka broker at %s...", i, getBrokerAddress())
		conn, err = kafka.DialLeader(context.Background(), "tcp", getBrokerAddress(), topicName, 0)
		if err == nil {
			break
		}

		log.Printf("Kafka connection failed: %v", err)
		time.Sleep(5 * time.Second)
	}

	if err != nil {
		log.Fatalf("Failed to connect to Kafka broker after %d attempts: %v", maxAttempts, err)
	}

	defer conn.Close()
	log.Println("✅ Connected to Kafka broker")
}

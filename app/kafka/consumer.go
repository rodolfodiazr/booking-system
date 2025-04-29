package kafka

import (
	"time"

	"github.com/segmentio/kafka-go"
)

func NewKafkaReader(topic, groupID string) *kafka.Reader {
	return kafka.NewReader(kafka.ReaderConfig{
		Brokers:  []string{getBrokerAddress()},
		GroupID:  groupID,
		Topic:    topic,
		MinBytes: 10e3, // 10KB
		MaxBytes: 10e6, // 10MB
		MaxWait:  1 * time.Second,
	})
}

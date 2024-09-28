package kafka_service

import (
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

type KService struct {
	C *kafka.Consumer
}

func NewConsumer(bootstrapServers, groupID, topic string) (KService, error) {
	// Create a new consumer with the specified configurations
	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers":        bootstrapServers,
		"group.id":                 groupID,
		"broker.address.family":    "v4", // Avoid connecting to IPv6 brokers
		"session.timeout.ms":       6000,
		"auto.offset.reset":        "earliest",
		"enable.auto.offset.store": false,
	})
	if err != nil {
		return KService{}, err
	}
	return KService{C: c}, nil
}

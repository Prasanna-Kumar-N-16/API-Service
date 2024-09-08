package kafka_service

import (
	"log"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

func NewConsumer(bootstrapServers, groupID, topic string) (*kafka.Consumer, error) {
	// Create a new consumer with the specified configurations
	c, err := kafka.NewConsumer(map[string]interface{}{
		"bootstrap.servers":        bootstrapServers,
		"group.id":                 groupID,
		"broker.address.family":    "v4", // Avoid connecting to IPv6 brokers
		"session.timeout.ms":       6000,
		"auto.offset.reset":        "earliest",
		"enable.auto.offset.store": false,
	})
	if err != nil {
		return nil, err
	}

	// Subscribe to the topic
	err = c.SubscribeTopics([]string{topic}, nil)
	if err != nil {
		return nil, err
	}
	return c, nil
}

func ConsumeMessages(c *kafka.Consumer) {
	for {
		msg, err := c.ReadMessage(-1)
		if err == nil {
			log.Printf("Message on %s: %s\n", msg.TopicPartition, string(msg.Value))
		} else {
			log.Printf("Consumer error: %v (%v)\n", err, msg)
		}
	}
}

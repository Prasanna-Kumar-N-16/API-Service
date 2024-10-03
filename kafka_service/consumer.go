package kafka_service

import (
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

type KService struct {
	C *kafka.Consumer
	P *kafka.Producer
}

func NewKService() KService {
	return KService{}
}

// creates a new consumer service
func (k *KService) NewConsumer(bootstrapServers, groupID, topic string) error {
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
		return err
	}
	k.C = c
	return nil
}

func (k *KService) ConsumeMessages(topic string) ([]byte, error) {
	err := k.C.SubscribeTopics([]string{topic}, nil)
	if err != nil {
		return nil, err
	}

	msg, err := k.C.ReadMessage(2)
	if err != nil {
		return nil, err
	}
	return msg.Value, nil

}

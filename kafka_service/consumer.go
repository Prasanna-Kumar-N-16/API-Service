package kafka_service

import (
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

type KService struct {
	C *kafka.Consumer
	P *kafka.Producer
}

// New service for kafka
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

func (k *KService) NewProducer(brokers string) error {
	p, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": brokers})
	if err != nil {
		return nil
	}
	k.P = p
	return nil
}

func (k *KService) ProduceMessage(topic string, message []byte) error {
	err := k.P.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: 0},
		Value:          message,
	}, nil)
	if err != nil {
		return err
	}
	// Wait for message deliveries
	k.P.Flush(15 * 1000)
	return nil
}

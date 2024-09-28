package kafka_service

import (
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

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

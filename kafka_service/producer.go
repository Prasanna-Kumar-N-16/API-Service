package kafka_service

import (
	k "github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

func NewProducer(brokers string) (*k.Producer, error) {
	p, err := k.NewProducer(&k.ConfigMap{"bootstrap.servers": brokers})
	if err != nil {
		return nil, err
	}
	return p, nil
}

func ProduceMessage(p *k.Producer, topic string, message []byte) error {
	err := p.Produce(&k.Message{
		TopicPartition: k.TopicPartition{Topic: &topic, Partition: 0},
		Value:          message,
	}, nil)
	if err != nil {
		return err
	}
	// Wait for message deliveries
	p.Flush(15 * 1000)
	return nil
}

package go_kafka

import (
	k "github.com/confluentinc/confluent-kafka-go/kafka"
)

func NewProducer(brokers string) (*k.Producer, error) {
	p, err := k.NewProducer(&k.ConfigMap{"bootstrap.servers": brokers})
	if err != nil {
		return nil, err
	}
	return p, nil
}

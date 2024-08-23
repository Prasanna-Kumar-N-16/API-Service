package go_kafka

import (
	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func NewConsumer(brokers, groupID, topic string) (*kafka.Consumer, error) {
	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": brokers,
		"group.id":          groupID,
		"auto.offset.reset": "earliest",
	})
	if err != nil {
		return nil, err
	}

	err = c.SubscribeTopics([]string{topic}, nil)
	if err != nil {
		return nil, err
	}
	return c, nil
}

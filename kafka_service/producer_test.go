package kafka_service_test

import (
	"testing"

	k "github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"kafka_service" // replace with the actual import path of your kafka_service package
)

// MockProducer to mock the actual Kafka producer
type MockProducer struct {
	mock.Mock
}

func (m *MockProducer) Produce(msg *k.Message, deliveryChan chan k.Event) error {
	args := m.Called(msg, deliveryChan)
	return args.Error(0)
}

func (m *MockProducer) Flush(timeoutMs int) int {
	return 0
}

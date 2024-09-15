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

func TestNewProducer_Success(t *testing.T) {
	brokers := "localhost:9092"
	producer, err := kafka_service.NewProducer(brokers)
	assert.NoError(t, err)
	assert.NotNil(t, producer)
}

func TestNewProducer_Failure(t *testing.T) {
	brokers := "invalid-broker"
	producer, err := kafka_service.NewProducer(brokers)
	assert.Error(t, err)
	assert.Nil(t, producer)
}

func TestProduceMessage_Success(t *testing.T) {
	mockProducer := new(MockProducer)
	topic := "test-topic"
	message := []byte("test-message")

	// Set up expectations
	mockProducer.On("Produce", mock.Anything, nil).Return(nil)

	err := kafka_service.ProduceMessage(mockProducer, topic, message)
	assert.NoError(t, err)

	mockProducer.AssertExpectations(t)
}

func TestProduceMessage_Failure(t *testing.T) {
	mockProducer := new(MockProducer)
	topic := "test-topic"
	message := []byte("test-message")

	// Set up expectations for failure
	mockProducer.On("Produce", mock.Anything, nil).Return(assert.AnError)

	err := kafka_service.ProduceMessage(mockProducer, topic, message)
	assert.Error(t, err)

	mockProducer.AssertExpectations(t)
}

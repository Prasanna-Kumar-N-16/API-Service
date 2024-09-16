package kafka_service_test

import (
	"api-service/kafka_service"
	"errors"
	"testing"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	// replace with the actual path of your kafka_service package
)

// MockConsumer struct to mock the Kafka consumer
type MockConsumer struct {
	mock.Mock
}

func (m *MockConsumer) SubscribeTopics(topics []string, rebalanceCb kafka.RebalanceCb) error {
	args := m.Called(topics, rebalanceCb)
	return args.Error(0)
}

func (m *MockConsumer) ReadMessage(timeoutMs int) (*kafka.Message, error) {
	args := m.Called(timeoutMs)
	if msg, ok := args.Get(0).(*kafka.Message); ok {
		return msg, args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MockConsumer) Close() error {
	return nil
}
func TestNewConsumer_Success(t *testing.T) {
	bootstrapServers := "localhost:9092"
	groupID := "test-group"
	topic := "test-topic"

	consumer, err := kafka_service.NewConsumer(bootstrapServers, groupID, topic)
	assert.NoError(t, err)
	assert.NotNil(t, consumer)
}

func TestNewConsumer_SubscribeFailure(t *testing.T) {
	bootstrapServers := "localhost:9092"
	groupID := "test-group"
	topic := "test-topic"

	// Mock the subscribe failure
	consumerMock := new(MockConsumer)
	consumerMock.On("SubscribeTopics", mock.Anything, nil).Return(errors.New("subscribe error"))

	consumer, err := kafka_service.NewConsumer(bootstrapServers, groupID, topic)
	assert.Nil(t, consumer)
	assert.Error(t, err)
}

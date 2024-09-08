package login_test

import (
	"github.com/stretchr/testify/mock"
)

// Mocks
type MockPostgresService struct {
	mock.Mock
}

type MockEmailService struct {
	mock.Mock
}

// Implement mock methods
func (m *MockPostgresService) AutoMigrate(value interface{}) error {
	args := m.Called(value)
	return args.Error(0)
}

func (m *MockPostgresService) Create(value interface{}) *MockPostgresService {
	m.Called(value)
	return m
}

type MockLogger struct{}

func (m *MockLogger) Errorln(args ...interface{}) {}

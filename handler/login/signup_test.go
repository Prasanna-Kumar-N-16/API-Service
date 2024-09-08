package login_test

import (
	"api-service/config"
	"api-service/utils"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
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

// Test case for Signup handler
func TestSignup(t *testing.T) {
	// Set Gin to test mode
	gin.SetMode(gin.TestMode)

	// Setup mock dependencies
	mockPostgres := new(MockPostgresService)
	mockEmail := new(MockEmailService)

	// Setup configuration
	config := config.ConfigStruct{
		Domain:     "example.com",
		EncryptKey: "test-encryption-key",
		Email: utils.EmailConfig{
			Username: "abcd",
			Password: "efgh",
		},
	}

	// Mocking services
	mockPostgres.On("AutoMigrate", &Admin{}).Return(nil)
	mockPostgres.On("Create", mock.Anything).Return(nil)

	// Mocking OTP generation and email sending
	utils.GenerateOTP(5)

	utils.NewSMTPClient("host", nil)

	mockEmail.On("SendOTPEmail", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(nil)

	// Create the handler with mocked services and config
	handler := &Authenticationhandler{
		service: mockPostgres,
		c:       config,
	}

	// Create a test Gin router and register the handler
	router := gin.Default()
	router.POST("/signup", handler.Signup)

	// Create a test request
	reqBody := AdminSignupRequest{
		Email:    "admin@example.com",
		Password: "password123",
	}
	jsonBody, _ := json.Marshal(reqBody)
	req, _ := http.NewRequest(http.MethodPost, "/signup", bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")

	// Create a ResponseRecorder to capture the response
	w := httptest.NewRecorder()

	// Serve the request
	router.ServeHTTP(w, req)

	// Assert that the response status code is 200 OK
	assert.Equal(t, http.StatusOK, w.Code)

	// Assert the response body
	expectedBody := `{"message":"Admin registered successfully! Please check your email to verify your account."}`
	assert.JSONEq(t, expectedBody, w.Body.String())

	// Assert that all mocks were called
	mockPostgres.AssertExpectations(t)
	mockEmail.AssertExpectations(t)
}

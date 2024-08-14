package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSendOTPEmail(t *testing.T) {
	emailConfig := EmailConfig{
		Username: "testuser@gmail.com",
		Password: "password123",
	}

	mockClient := &mockSMTPClient{}

	err := emailConfig.SendOTPEmail(mockClient, "recipient@example.com", "TestPortal", "http://testportal.com", "123456")
	assert.NoError(t, err)

	assert.True(t, mockClient.authCalled, "expected auth to be called")
	assert.True(t, mockClient.mailCalled, "expected mail to be called")
	assert.True(t, mockClient.rcptCalled, "expected rcpt to be called")
	assert.True(t, mockClient.dataCalled, "expected data to be called")
	assert.True(t, mockClient.quitCalled, "expected quit to be called")

	expectedMsg := "From: testuser@gmail.com\nTo: recipient@example.com\nSubject: Welcome to TestPortal - Your OTP\n\nWelcome to TestPortal!\n\nYour OTP is: 123456\n\nPlease use this OTP to complete your registration.\n\nYou can access the portal here: http://testportal.com\n\nThank you!"
	assert.Equal(t, expectedMsg, mockClient.message, "expected message to be %q, but got %q", expectedMsg, mockClient.message)
}

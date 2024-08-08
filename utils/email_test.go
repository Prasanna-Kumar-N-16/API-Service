package utils

import (
	"testing"
)

func TestSendOTPEmail(t *testing.T) {
	emailConfig := EmailConfig{
		Username: "testuser@gmail.com",
		Password: "password123",
	}

	mockClient := &mockSMTPClient{}

	err := emailConfig.sendOTPEmail(mockClient, "recipient@example.com", "TestPortal", "http://testportal.com", "123456")
	if err != nil {
		t.Fatalf("expected no error, but got %v", err)
	}

}

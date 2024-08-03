package utils

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Test function for generateOTP
func TestGenerateOTP(t *testing.T) {
	otpLength := 6

	// Generate OTP
	otp, err := GenerateOTP(otpLength)

	fmt.Println(otp)

	// Use assert to verify the OTP length
	assert.NoError(t, err, "Expected no error while generating OTP")
	assert.Equal(t, otpLength, len(otp), "Expected OTP length to be %d", otpLength)

	// Verify that the OTP contains only alphanumeric characters
	for _, char := range otp {
		assert.True(t, isNumeric(char), "Expected OTP to contain only alphanumeric characters")
	}
}

// Helper function to check if a character is alphanumeric
func isNumeric(char rune) bool {
	return (char >= '0' && char <= '9')
}

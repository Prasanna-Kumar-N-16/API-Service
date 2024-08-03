package utils

import (
	"crypto/rand"
	"encoding/base64"
	"strings"
)

func IsAdminEmail(email, adminDomain string) bool {
	return strings.HasSuffix(email, "@"+adminDomain)
}

// Function to generate a random OTP
func generateOTP(length int) (string, error) {
	otpBytes := make([]byte, length)
	_, err := rand.Read(otpBytes)
	if err != nil {
		return "", err
	}
	otp := base64.StdEncoding.EncodeToString(otpBytes)
	return otp[:length], nil
}

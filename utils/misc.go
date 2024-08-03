package utils

import (
	"crypto/rand"
	"math/big"
	"strings"
)

func IsAdminEmail(email, adminDomain string) bool {
	return strings.HasSuffix(email, "@"+adminDomain)
}

// Function to generate a random OTP
func GenerateOTP(length int) (string, error) {
	otp := make([]byte, length)
	for i := 0; i < length; i++ {
		num, err := rand.Int(rand.Reader, big.NewInt(10))
		if err != nil {
			return "", err
		}
		otp[i] = '0' + byte(num.Int64())
	}
	return string(otp), nil
}

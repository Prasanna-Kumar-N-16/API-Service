package utils

import (
	"crypto/rand"
	"encoding/json"
	"fmt"
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

// Function to Marshal and Unmarshal
func MarshalUnmarshal(input interface{}, output interface{}) error {
	// Marshal the input interface into JSON
	jsonData, err := json.Marshal(input)
	if err != nil {
		return fmt.Errorf("error during marshal: %v", err)
	}

	// Unmarshal the JSON data into the output struct
	err = json.Unmarshal(jsonData, output)
	if err != nil {
		return fmt.Errorf("error during unmarshal: %v", err)
	}

	return nil
}

package encryption

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"io"
)

// Encrypt encrypts plain text string into cipher text string using AES256
func Encrypt(plainText, key string) (string, error) {
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return "", err
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	nonce := make([]byte, aesGCM.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return "", err
	}

	cipherText := aesGCM.Seal(nonce, nonce, []byte(plainText), nil)
	return base64.StdEncoding.EncodeToString(cipherText), nil
}

// Decrypt decrypts cipher text string into plain text string using AES256
func Decrypt(cipherText, key string) (string, error) {
	// Decode the base64 encoded cipherText
	decodedCipherText, err := base64.StdEncoding.DecodeString(cipherText)
	if err != nil {
		return "", err
	}

	// Create the AES cipher block
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return "", err
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	// Extract the nonce from the cipherText
	nonceSize := aesGCM.NonceSize()
	nonce, cipherTextBytes := decodedCipherText[:nonceSize], decodedCipherText[nonceSize:]

	// Decrypt the cipherText
	plainText, err := aesGCM.Open(nil, nonce, cipherTextBytes, nil)
	if err != nil {
		return "", err
	}

	return string(plainText), nil
}

func VerifyPassword(enteredPassword, storedEncryptedPassword, key string) (bool, error) {
	// Encrypt the entered password using the same encryption method
	encryptedPassword, err := Encrypt(enteredPassword, key)
	if err != nil {
		return false, err
	}

	// Compare the encrypted entered password with the stored encrypted password
	if encryptedPassword == storedEncryptedPassword {
		return true, nil
	}

	return false, nil
}

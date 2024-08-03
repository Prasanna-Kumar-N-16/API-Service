package utils

import (
	"crypto/tls"
	"fmt"
	"net/smtp"
)

const (
	smtpHost = "smtp.gmail.com"
	smtpPort = "587"
)

type EmailConfig struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// Function to send an email with OTP
func (u EmailConfig) sendOTPEmail(email, portalName, portalLink, otp string) error {

	// Create the email content
	subject := fmt.Sprintf("Welcome to %s - Your OTP", portalName)
	body := fmt.Sprintf("Welcome to %s!\n\nYour OTP is: %s\n\nPlease use this OTP to complete your registration.\n\nYou can access the portal here: %s\n\nThank you!", portalName, otp, portalLink)
	msg := fmt.Sprintf("From: %s\nTo: %s\nSubject: %s\n\n%s", u.Username, email, subject, body)

	// Connect to the SMTP server
	auth := smtp.PlainAuth("", u.Username, u.Password, smtpHost)
	tlsConfig := &tls.Config{
		InsecureSkipVerify: true,
		ServerName:         smtpHost,
	}
	conn, err := tls.Dial("tcp", smtpHost+":"+smtpPort, tlsConfig)
	if err != nil {
		return err
	}
	client, err := smtp.NewClient(conn, smtpHost)
	if err != nil {
		return err
	}

	// Authenticate
	if err = client.Auth(auth); err != nil {
		return err
	}

	// Set the sender and recipient
	if err = client.Mail(u.Username); err != nil {
		return err
	}
	if err = client.Rcpt(email); err != nil {
		return err
	}

	// Send the email body
	wc, err := client.Data()
	if err != nil {
		return err
	}
	_, err = wc.Write([]byte(msg))
	if err != nil {
		return err
	}
	err = wc.Close()
	if err != nil {
		return err
	}

	return client.Quit()
}

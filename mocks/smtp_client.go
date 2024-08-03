package mocks

import (
	"crypto/tls"
	"net/smtp"
)

// SMTPClient is an interface representing an SMTP client
type SMTPClient interface {
	Mail(from string) error
	Rcpt(to string) error
	Auth(a smtp.Auth) error
	Quit() error
}

// TLSDialer is an interface representing a TLS dialer
type TLSDialer interface {
	Dial(network, address string, config *tls.Config) (*tls.Conn, error)
}

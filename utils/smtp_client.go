package utils

import (
	"net/smtp"
)

type SMTPClient interface {
	Auth(auth smtp.Auth) error
	Mail(from string) error
	Rcpt(to string) error
	Data() (smtpClientWriter, error)
	Quit() error
}

type smtpClient struct {
	client *smtp.Client
}

type smtpClientWriter interface {
	Write([]byte) (int, error)
	Close() error
}

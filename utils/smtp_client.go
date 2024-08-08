package utils

import (
	"crypto/tls"
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

func (c *smtpClient) Auth(auth smtp.Auth) error {
	return c.client.Auth(auth)
}

func (c *smtpClient) Mail(from string) error {
	return c.client.Mail(from)
}

func (c *smtpClient) Rcpt(to string) error {
	return c.client.Rcpt(to)
}

func (c *smtpClient) Data() (smtpClientWriter, error) {
	return c.client.Data()
}

func (c *smtpClient) Quit() error {
	return c.client.Quit()
}

func newSMTPClient(host string, tlsConfig *tls.Config) (SMTPClient, error) {
	conn, err := tls.Dial("tcp", host+":"+smtpPort, tlsConfig)
	if err != nil {
		return nil, err
	}
	client, err := smtp.NewClient(conn, host)
	if err != nil {
		return nil, err
	}
	return &smtpClient{client: client}, nil
}

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

// Mock implementations for testing

type mockSMTPClient struct {
	authCalled bool
	mailCalled bool
	rcptCalled bool
	dataCalled bool
	quitCalled bool
	message    string
	authError  error
	mailError  error
	rcptError  error
	dataError  error
	writeError error
	closeError error
	quitError  error
}

func (m *mockSMTPClient) Auth(auth smtp.Auth) error {
	m.authCalled = true
	return m.authError
}

func (m *mockSMTPClient) Mail(from string) error {
	m.mailCalled = true
	return m.mailError
}

func (m *mockSMTPClient) Rcpt(to string) error {
	m.rcptCalled = true
	return m.rcptError
}

func (m *mockSMTPClient) Data() (smtpClientWriter, error) {
	m.dataCalled = true
	return &mockSMTPClientWriter{mock: m}, m.dataError
}

func (m *mockSMTPClient) Quit() error {
	m.quitCalled = true
	return m.quitError
}

type mockSMTPClientWriter struct {
	mock *mockSMTPClient
}

func (w *mockSMTPClientWriter) Write(p []byte) (int, error) {
	w.mock.message = string(p)
	return len(p), w.mock.writeError
}

func (w *mockSMTPClientWriter) Close() error {
	return w.mock.closeError
}

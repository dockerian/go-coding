// Package msg :: mail_mockClient.go
package msg

import (
	"crypto/tls"
	"errors"
	"io"
	"net/smtp"
)

// clientWriteCloser struct
type clientWriteCloser struct {
	closeErr string
	writeErr string
}

// Close implemnents io.WriteCloser
func (c *clientWriteCloser) Close() error {
	if c.closeErr != "" {
		return errors.New(c.closeErr)
	}
	return nil
}

// Write implemnents io.WriteCloser
func (c *clientWriteCloser) Write(data []byte) (int, error) {
	if c.writeErr != "" {
		return 0, errors.New(c.writeErr)
	}
	return 10, nil
}

// mockClient struct
type mockClient struct {
	*clientWriteCloser
	helloErr       string
	startTLSErr    string
	authErr        string
	mailErr        string
	rcptErr        string
	dataErr        string
	clientWriteErr string
	clientCloseErr string
	quitErr        string
	closeErr       string
}

func (m *mockClient) Auth(auth smtp.Auth) error {
	if m.authErr != "" {
		return errors.New(m.authErr)
	}
	return nil
}

func (m *mockClient) Close() error {
	if m.closeErr != "" {
		return errors.New(m.closeErr)
	}
	return nil

}

func (m *mockClient) Data() (io.WriteCloser, error) {
	if m.dataErr != "" {
		return m.clientWriteCloser, errors.New(m.dataErr)
	}
	return m.clientWriteCloser, nil
}

func (m *mockClient) Extension(string) (bool, string) {
	return true, ""
}

func (m *mockClient) Hello(string) error {
	if m.helloErr != "" {
		return errors.New(m.helloErr)
	}
	return nil
}

func (m *mockClient) Mail(string) error {
	if m.mailErr != "" {
		return errors.New(m.mailErr)
	}
	return nil
}

func (m *mockClient) Quit() error {
	if m.quitErr != "" {
		return errors.New(m.quitErr)
	}
	return nil
}

func (m *mockClient) Rcpt(string) error {
	if m.rcptErr != "" {
		return errors.New(m.rcptErr)
	}
	return nil
}

func (m *mockClient) Reset() error {
	return nil
}

func (m *mockClient) StartTLS(*tls.Config) error {
	if m.startTLSErr != "" {
		return errors.New(m.startTLSErr)
	}
	return nil
}

func (m *mockClient) TLSConnectionState() (tls.ConnectionState, bool) {
	return tls.ConnectionState{}, true
}

func (m *mockClient) Verify(string) error {
	return nil
}

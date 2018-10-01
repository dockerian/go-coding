// +build all common pkg msg mail

// Package msg :: mail_test.go
package msg

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"math/big"
	"mime/multipart"
	"net/smtp"
	"net/textproto"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	base64tests = map[string]string{
		"attached text":    "YXR0YWNoZWQgdGV4dA==",
		"attachment test1": "YXR0YWNobWVudCB0ZXN0MQ==",
		"more attachment":  "bW9yZSBhdHRhY2htZW50",
	}
	recipients = map[string][]string{
		"to":  {"to1@t.com", "to2@t.com"},
		"cc":  {"cc1@t1.com", "cc2@t2.com"},
		"bcc": {"bcc1@test.com", "bcc2@test.com"},
	}
	messages = []Message{
		{
			From:    "from@a.com",
			To:      recipients["to"],
			Cc:      recipients["cc"],
			Bcc:     recipients["bcc"],
			Subject: "mail test subject",
			Body: []byte(`<h1>This email header</h2>
                    and mail body content with a few more lines
                    and more and more`),
		},
		{
			To:      recipients["to"],
			Subject: "mail test subject 1",
			Body:    []byte(`mail body test 1`),
		},
		{
			From:  "from@a.com",
			To:    []string{},
			Body:  []byte("mail body"),
			Plain: true,
		},
		{
			From: "bad email address",
			To:   []string{"bad recipient address"},
		},
		{
			From: `"" <>`,
			To:   recipients["to"],
		},
		{},
	}
	mockSendMail = func(addr string, a smtp.Auth, from string, to []string, msg []byte) error {
		return nil
	}
	smtpAuth = smtp.PlainAuth("", "username", "password", "host")
)

// ErrWriter struct implements io.Writer
type ErrWriter struct {
}

func (errWriter *ErrWriter) Write(data []byte) (int, error) {
	return 0, errors.New("err to write")
}

// TestAttach
func TestAttach(t *testing.T) {
	message := messages[0]
	rdr := strings.NewReader("text")
	att, err := message.Attach(rdr, "unknown", "")
	assert.Nil(t, err)
	assert.NotNil(t, att)
	assert.Equal(t, "unknown", att.Filename)
	header := att.Header
	assert.Equal(t, "application/octet-stream", header.Get("Content-Type"))
}

// TestAttachError
func TestAttachError(t *testing.T) {
	ioCopyFunc := ioCopy
	ioCopy = func(w io.Writer, r io.Reader) (int64, error) {
		return 0, errors.New("io copy error")
	}
	message := messages[0]
	att := strings.NewReader("text")
	_, err := message.Attach(att, "a.txt", "text/plain")
	assert.Equal(t, "io copy error", err.Error())

	ioCopy = ioCopyFunc
}

// TestAttachFile
func TestAttachFile(t *testing.T) {
	message := messages[0]
	att0, err0 := message.AttachFile("doe not exist")
	assert.NotNil(t, err0)
	assert.Nil(t, att0)
	att1, err1 := message.AttachFile("mail.go")
	assert.NotNil(t, att1)
	assert.Nil(t, err1)
}

// TestGenerateID
func TestGenerateID(t *testing.T) {
	osHostnameFunc := osHostname
	osHostname = func() (name string, err error) {
		return "", errors.New("os.Hostname error")
	}
	msgID, _ := GenerateID()
	assert.Contains(t, msgID, "localhost")
	osHostname = osHostnameFunc

	randIntFunc := randInt
	randInt = func(r io.Reader, max *big.Int) (*big.Int, error) {
		return nil, errors.New("rand.Int error")
	}
	_, err := GenerateID()
	assert.Equal(t, "rand.Int error", err.Error())
	randInt = randIntFunc
}

// TestMessage tests Message's Attach and Bytes methods
func TestMessage(t *testing.T) {
	message := messages[0]
	content := string(message.Bytes())
	assert.True(t, len(content) > 0)

	to := fmt.Sprintf("To: %s", strings.Join(recipients["to"], ", "))
	assert.Contains(t, content, to)

	num := 1
	for text, encoded := range base64tests {
		var buffer bytes.Buffer
		writeBase64([]byte(text), &buffer)
		assert.Equal(t, buffer.String(), encoded)

		extension, contentType := "txt", "text/plain"
		if num%1 == 0 {
			extension, contentType = "bin", ""
		}
		att := strings.NewReader(text)
		name := fmt.Sprintf("%d.%s", num, extension)
		message.Attach(att, name, contentType)
		content = string(message.Bytes())
		assert.Contains(t, content, encoded)
		num += 1
	}

	var buffer = bytes.NewBuffer(nil)
	var str = strings.Repeat("...", 100)
	strReader := strings.NewReader(str)
	_, err := message.Attach(strReader, "dot.txt", "text/plain")
	assert.Nil(t, err)
	writeAttachedData([]byte(str), buffer)
	assert.True(t, len(buffer.Bytes()) > 100)
}

// TestNewDialClient
func TestNewDialClient(t *testing.T) {
	_, err := newDialClient("addr")
	assert.NotNil(t, err)
	expected := "dial tcp: address addr: missing port in address"
	assert.Equal(t, expected, err.Error())

	message := messages[0]
	sender := MessageSender{
		Message: message,
		WithTLS: true,
	}
	err = sender.sendWithTLS("", []string{}, "addr", smtpAuth, nil)
	assert.Equal(t, expected, err.Error())
}

// TestSend
func TestSend(t *testing.T) {
	for idx, test := range []struct {
		message Message
		err     error
	}{
		{
			message: messages[0], err: nil,
		},
		{
			message: messages[1], err: errors.New("missing message From address"),
		},
		{
			message: messages[2], err: errors.New("missing message To address(es)"),
		},
		{
			message: messages[3], err: errors.New("mail: no angle-addr"),
		},
		{
			message: messages[4], err: errors.New("mail: invalid string"),
		},
		{
			message: messages[5], err: errors.New("missing message From address"),
		},
	} {
		message := test.message
		sender := MessageSender{
			Message:  message,
			Password: "secret",
		}
		smtpSendMail = mockSendMail
		err := sender.Send()
		t.Logf("Test %2d: %+v\n", idx, err)
		assert.Equal(t, test.err, err)

		content := string(message.Bytes())
		if message.Plain {
			assert.Contains(t, content, "text/plain")
		} else {
			assert.Contains(t, content, "text/html")
		}
	}
}

// TestSendWithTLS
func TestSendWithTLS(t *testing.T) {
	var ccWriters = []*clientWriteCloser{
		{},
		{writeErr: "data write err"},
		{closeErr: "data close err"},
	}

	for num, test := range []struct {
		mock     *mockClient
		expected string
	}{
		{
			&mockClient{clientWriteCloser: ccWriters[0], helloErr: "hello err"},
			"hello err",
		},
		{
			&mockClient{clientWriteCloser: ccWriters[0], startTLSErr: "tls err"},
			"tls err",
		},
		{
			&mockClient{clientWriteCloser: ccWriters[0], startTLSErr: "tls err"},
			"tls err",
		},
		{
			&mockClient{clientWriteCloser: ccWriters[0], authErr: "auth err"},
			"auth err",
		},
		{
			&mockClient{clientWriteCloser: ccWriters[0], mailErr: "mail err"},
			"mail err",
		},
		{
			&mockClient{clientWriteCloser: ccWriters[0], rcptErr: "rcpt err"},
			"rcpt err",
		},
		{
			&mockClient{clientWriteCloser: ccWriters[0], dataErr: "data err"},
			"data err",
		},
		{
			&mockClient{clientWriteCloser: ccWriters[1]},
			"data write err",
		},
		{
			&mockClient{clientWriteCloser: ccWriters[2]},
			"data close err",
		},
		{
			&mockClient{clientWriteCloser: ccWriters[0], quitErr: "quit err"},
			"quit err",
		},
	} {
		smtpClient = func(addr string) (SMTPClient, error) {
			return test.mock, nil
		}

		message := messages[num%2]
		sender := MessageSender{
			Message:    message,
			DomainName: "test.com",
			UserName:   "foobar",
			Password:   "secret",
			WithTLS:    true,
		}
		t.Logf("Test %2d: %s\n", num+1, test.expected)
		sendErr := sender.Send()
		assert.Equal(t, test.expected, sendErr.Error())
	}
}

// TestWrite
func TestWrite(t *testing.T) {
	// testing writeBody with mocked writer to cause createPart error
	var errWriter = &ErrWriter{}
	var buffer = bytes.NewBuffer(nil)
	var mw = multipart.NewWriter(errWriter)
	err1 := writeBody([]byte("test"), buffer, "", mw)
	assert.NotNil(t, err1)
	_, err2 := mw.CreatePart(textproto.MIMEHeader{})
	assert.NotNil(t, err2)
}

// Package msg :: mail.go - A simple email constructor and sender
package msg

import (
	"bytes"
	"crypto/rand"
	"crypto/tls"
	"encoding/base64"
	"errors"
	"fmt"
	"io"
	"log"
	"math"
	"math/big"
	"mime"
	"mime/multipart"
	"mime/quotedprintable"
	"net/mail"
	"net/smtp"
	"net/textproto"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"time"
)

const (
	maxRaw = 57
	// maximum line length per RFC 2045
	maxLineLength = 76
)

var (
	// ioCopy sets function variable to io.Copy
	ioCopy = io.Copy
	// newPrintableWriter sets function variable to quotedprintable.NewWriter
	newPrintableWriter = quotedprintable.NewWriter
	// osHostname sets function variable to os.Hostname
	osHostname = os.Hostname
	// randInt sets function variable to rand.Int
	randInt = rand.Int
	// smtpClient sets function variable to newDialClient
	smtpClient = newDialClient
	// smtpSendMail sets function variable to smtp.SendMail
	smtpSendMail = smtp.SendMail
)

// Attachment represents an email attachment, with file name, MIMEHeader, and content
type Attachment struct {
	Filename string
	Header   textproto.MIMEHeader
	Content  []byte
}

// Message represents a simple e-mail message with attachments
type Message struct {
	From        string
	To          []string
	Cc          []string
	Bcc         []string
	Subject     string
	Body        []byte
	Attachments []*Attachment
	Headers     textproto.MIMEHeader
	Plain       bool
}

// MessageSender struct represents an email message sender
type MessageSender struct {
	// Message embeds Message struct
	Message
	// DomainName defines the domain name for mail sender
	DomainName string
	// ServerHost defines the mail server host
	ServerHost string
	// ServerPort defines the mail server port
	ServerPort uint32
	// UserName defines the login username for the mail server
	UserName string
	// Password defines the login password for the mail server
	Password string
	// PlainAuth specifies to use PlainAuth if WithTLS is not enabled
	PlainAuth bool
	// WithTLS specifies to use TLS
	WithTLS bool
}

// Sender interface represents a message sender
type Sender interface {
	Send() error
}

// SMTPClient interface
type SMTPClient interface {
	Auth(smtp.Auth) error
	Close() error
	Data() (io.WriteCloser, error)
	Extension(string) (bool, string)
	Hello(string) error
	Mail(string) error
	Quit() error
	Rcpt(string) error
	Reset() error
	StartTLS(*tls.Config) error
	TLSConnectionState() (tls.ConnectionState, bool)
	Verify(string) error
}

// Attach adds content from an io.Reader to email attachments
func (msg *Message) Attach(r io.Reader, filename string, contentType string) (*Attachment, error) {
	var content bytes.Buffer
	if _, err := ioCopy(&content, r); err != nil {
		log.Printf("[mail] ERROR: %s\n", err)
		return nil, err
	}

	part := &Attachment{
		Filename: filename,
		Header:   textproto.MIMEHeader{},
		Content:  content.Bytes(),
	}

	if contentType == "" {
		contentType = mime.TypeByExtension(filepath.Ext(filename))
	}

	if contentType == "" {
		contentType = "application/octet-stream"
	}

	part.Header.Set("Content-ID", fmt.Sprintf("<%s>", filename))
	part.Header.Set("Content-Disposition", fmt.Sprintf("attachment;\r\n filename=\"%s\"", filename))
	part.Header.Set("Content-Transfer-Encoding", "base64")
	part.Header.Set("Content-Type", contentType)

	siz := len(part.Content)
	log.Printf("[mail] adding attachment '%s' (%s) %d\n", filename, contentType, siz)
	msg.Attachments = append(msg.Attachments, part)

	return part, nil
}

// AttachFile adds content from a filename to email attachment
func (msg *Message) AttachFile(filename string) (*Attachment, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	basename := filepath.Base(filename)
	contentType := mime.TypeByExtension(filepath.Ext(filename))

	log.Printf("[mail] adding attachment file '%s'\n", filename)
	return msg.Attach(file, basename, contentType)
}

// Bytes converts an e-mail subject, body, and attachments to bytes
func (msg *Message) Bytes() []byte {
	buffer := bytes.NewBuffer(nil)

	msg.setHeaders()

	var boundary string
	var bodyType = "text/html"
	var hasMultiPart = len(msg.Attachments) > 0
	// Note: multiplart.Writer does not implement io.Writer but can create one
	var mw *multipart.Writer

	if msg.Plain {
		bodyType = "text/plain"
	}

	if hasMultiPart {
		mw = multipart.NewWriter(buffer)
		boundary = mw.Boundary()
		msg.Headers.Set("Content-Type", "multipart/mixed;\r\n\tboundary="+boundary)
	} else {
		msg.Headers.Set("Content-Type", bodyType+"; charset=UTF-8")
	}

	msg.writeHeader(buffer)

	writeBody(msg.Body, buffer, bodyType, mw)

	if len(msg.Attachments) > 0 {
		for _, att := range msg.Attachments {
			if partWriter, err := mw.CreatePart(att.Header); err == nil {
				writeAttachedData(att.Content, partWriter)
			}
		}
	}

	return buffer.Bytes()
}

// GenerateID returns an RFC 2822 compliant Message-ID, e.g.:
//
//     <1513099495572712211.62751.4513881499840151067@user.domain>
//
// by using
//   - current time nanoseconds since Epoch
//   - operating system process PID
//   - A cryptographically random int64
//   - sender's hostname
//
func GenerateID() (string, error) {
	sec := time.Now().UnixNano()
	pid := os.Getpid()
	var maxBigInt = big.NewInt(math.MaxInt64)
	hashInt, err := randInt(rand.Reader, maxBigInt)
	if err != nil {
		return "", err
	}
	hostname, err := osHostname()
	if err != nil {
		hostname = "localhost"
	}
	messageID := fmt.Sprintf("<%d.%d.%d@%s>", sec, pid, hashInt, hostname)

	return messageID, nil
}

// newDialClient returns an SMTPClient
func newDialClient(addr string) (SMTPClient, error) {
	return smtp.Dial(addr)
}

// Send is a sender pointer receiver to send email message
func (sender *MessageSender) Send() error {
	if sender.Message.From == "" {
		if sender.DomainName != "" && sender.UserName != "" {
			if strings.Contains(sender.UserName, "@") {
				sender.Message.From = sender.UserName
			} else {
				emailAddress := fmt.Sprintf("%s@%s", sender.UserName, sender.DomainName)
				sender.Message.From = emailAddress
			}
		}
	}

	toList, from, err := sender.Message.parseAddresses()
	if err != nil {
		return err
	}

	odump := fmt.Sprintf("%+v", sender)
	regex := regexp.MustCompile("Password:[^\\s]+")
	senderText := regex.ReplaceAllString(odump, "Password:********")

	log.Printf("[mail] sending %+s\n", senderText)

	addr := fmt.Sprintf("%s:%d", sender.ServerHost, sender.ServerPort)
	auth := LoginAuth(sender.UserName, sender.Password, sender.ServerHost)

	if sender.WithTLS {
		log.Println("[mail] TLS enabled")
		config := &tls.Config{ServerName: sender.ServerHost}
		return sender.sendWithTLS(from, toList, addr, auth, config)
	}

	if sender.PlainAuth {
		auth = smtp.PlainAuth("", sender.UserName, sender.Password, sender.ServerHost)
	}
	return sender.sendMail(from, toList, addr, auth)
}

func (sender *MessageSender) sendMail(
	from string, toList []string, addr string, auth smtp.Auth) error {

	data := sender.Message.Bytes()

	return smtpSendMail(addr, auth, from, toList, data)
}

// SendWithTLS sends an email with an optional TLS config
func (sender *MessageSender) sendWithTLS(
	from string, toList []string,
	addr string, auth smtp.Auth, config *tls.Config) error {

	// see https://github.com/golang/go/blob/master/src/net/smtp/smtp.go
	client, err := smtpClient(addr)
	if err != nil {
		log.Printf("[mail] smtp.Dial: %s\n", err)
		return err
	}
	defer client.Close()

	if err = client.Hello("localhost"); err != nil {
		log.Printf("[mail] smtp client.Hello: %s\n", err)
		return err
	}

	// start TLS
	if ok, _ := client.Extension("STARTTLS"); ok {
		if err = client.StartTLS(config); err != nil {
			log.Printf("[mail] smtp client.StartTLS: %s\n", err)
			return err
		}
	}

	if auth != nil {
		if ok, _ := client.Extension("AUTH"); ok {
			if err = client.Auth(auth); err != nil {
				log.Printf("[mail] smtp client.Auth: %s\n", err)
				return err
			}
		}
	}

	if err = client.Mail(from); err != nil {
		log.Printf("[mail] smtp client.Mail: %s\n", err)
		return err
	}

	for _, addr := range toList {
		if err = client.Rcpt(addr); err != nil {
			log.Printf("[mail] smtp client.Rcpt: %s\n", err)
			return err
		}
	}

	writeCloser, err := client.Data()
	if err != nil {
		log.Printf("[mail] smtp client.Data: %s\n", err)
		return err
	}

	bytes := sender.Message.Bytes()
	if _, err = writeCloser.Write(bytes); err != nil {
		writeCloser.Close()
		log.Printf("[mail] smtp client write: %s\n", err)
		return err
	}
	if err = writeCloser.Close(); err != nil {
		log.Printf("[mail] smtp client writer.Close: %s\n", err)
		return err
	}

	return client.Quit()
}

// setHeaders gets message headers
func (msg *Message) setHeaders() {
	headers := make(textproto.MIMEHeader)

	if msgID, err := GenerateID(); err == nil {
		headers.Set("Message-Id", msgID)
	}

	if len(msg.To) > 0 {
		headers.Set("To", strings.Join(msg.To, ", "))
	}
	if len(msg.Cc) > 0 {
		headers.Set("Cc", strings.Join(msg.Cc, ", "))
	}

	headers.Set("From", msg.From)
	headers.Set("Subject", msg.Subject)
	headers.Set("Date", time.Now().Format(time.RFC1123Z))
	headers.Set("MIME-Version", "1.0")

	msg.Headers = headers
}

// parseAddress parses an SMTP envelope address
func parseAddress(mailAddress string) (address *mail.Address, err error) {
	strimmed := strings.TrimSpace(mailAddress)

	if address, err = mail.ParseAddress(strimmed); err == nil {
		return address, nil
	}

	log.Printf("[mail] parseAddress: %s\n", err)
	return nil, err
}

// parseAddresses returns a list of validated recipients and from address
func (msg *Message) parseAddresses() (toList []string, from string, err error) {
	if msg == nil || msg.From == "" {
		err = errors.New("missing message From address")
		log.Printf("[mail] parseAddresses: %s\n", err)
		return
	}

	toList = make([]string, 0, len(msg.To)+len(msg.Cc)+len(msg.Bcc))
	toList = append(toList, msg.To...)
	toList = append(toList, msg.Cc...)
	toList = append(toList, msg.Bcc...)

	log.Printf("[mail] toList: %+v\n", toList)

	for i := 0; i < len(toList); i++ {
		var addr *mail.Address
		if addr, err = mail.ParseAddress(toList[i]); err != nil {
			log.Printf("[mail] parsing toList[%d] %s: %s\n", i, toList[i], err)
			return
		}
		toList[i] = addr.Address
	}
	if len(toList) == 0 {
		err = errors.New("missing message To address(es)")
		log.Printf("[mail] parseAddress: %s\n", err)
		return
	}

	from = msg.From
	fromAddr, err := parseAddress(msg.From)
	if err != nil {
		log.Printf("[mail] parseAddress: %s\n", err)
		return
	}
	from = fromAddr.Address

	return
}

// writeHeader writes message header to buffer
func (msg *Message) writeHeader(buffer *bytes.Buffer) {
	for key, values := range msg.Headers {
		for _, keyval := range values {
			io.WriteString(buffer, key)
			io.WriteString(buffer, ": ")
			switch {
			case key == "Content-Type" || key == "Content-Disposition":
				buffer.Write([]byte(keyval))
			default:
				buffer.Write([]byte(mime.QEncoding.Encode("UTF-8", keyval)))
			}
			io.WriteString(buffer, "\r\n")
		}
	}
	io.WriteString(buffer, "\r\n")
}

// writeAttachedData encodes the data, by RFC 2045 (76 chars per line)
// and writes output specified multipart writer
func writeAttachedData(data []byte, partWriter io.Writer) {
	// buffer line with trailing CRLF
	buffer := make([]byte, maxLineLength+len("\r\n"))
	copy(buffer[maxLineLength:], "\r\n")

	// for loop raw chunks until bytes shorter than a line
	for len(data) >= maxRaw {
		base64.StdEncoding.Encode(buffer, data[:maxRaw])
		partWriter.Write(buffer)
		data = data[maxRaw:]
	}
	// write the last chunk of data
	if len(data) > 0 {
		siz := base64.StdEncoding.EncodedLen(len(data))
		out := buffer[:siz]
		base64.StdEncoding.Encode(out, data)
		out = append(out, "\r\n"...)
		partWriter.Write(out)
	}
}

// writeBase64 writes encoded data to writer
func writeBase64(data []byte, partWriter io.Writer) error {
	bufsiz := base64.StdEncoding.EncodedLen(len(data))
	buffer := make([]byte, bufsiz)
	base64.StdEncoding.Encode(buffer, data)
	_, err := partWriter.Write(buffer)

	return err
}

// writeBody writes message body to buffer, with MIME part (if multipart writer is provided)
func writeBody(body []byte, buffer *bytes.Buffer, contentType string, mw *multipart.Writer) error {
	if mw != nil {
		encoding := "quoted-printable"
		partHeader := textproto.MIMEHeader{
			"Content-Type":              {contentType + "; charset=UTF-8"},
			"Content-Transfer-Encoding": {encoding},
		}
		if _, err := mw.CreatePart(partHeader); err != nil {
			return err
		}
	}

	data := append(body, "\r\n"...)
	qp := quotedprintable.NewWriter(buffer)
	if _, err := qp.Write(data); err != nil {
		return err
	}
	return qp.Close()
}

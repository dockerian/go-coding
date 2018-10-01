# msg
--
    import "github.com/dockerian/go-coding/pkg/msg"

Package msg :: mail.go - A simple email constructor and sender

Package msg :: mail_auth.go - smtp.Auth implementations see

    - http://www.samlogic.net/articles/smtp-commands-reference-auth.htm
    - https://golang.org/src/net/smtp/auth.go

## Usage

#### func  GenerateID

```go
func GenerateID() (string, error)
```
GenerateID returns an RFC 2822 compliant Message-ID, e.g.:

    <1513099495572712211.62751.4513881499840151067@user.domain>

by using

    - current time nanoseconds since Epoch
    - operating system process PID
    - A cryptographically random int64
    - sender's hostname

#### func  LoginAuth

```go
func LoginAuth(username, password, host string) smtp.Auth
```
LoginAuth returns an smtp.Auth implementation for LOGIN authentication as
defined in RFC 4616

#### type Attachment

```go
type Attachment struct {
	Filename string
	Header   textproto.MIMEHeader
	Content  []byte
}
```

Attachment represents an email attachment, with file name, MIMEHeader, and
content

#### type Message

```go
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
```

Message represents a simple e-mail message with attachments

#### func (*Message) Attach

```go
func (msg *Message) Attach(r io.Reader, filename string, contentType string) (*Attachment, error)
```
Attach adds content from an io.Reader to email attachments

#### func (*Message) AttachFile

```go
func (msg *Message) AttachFile(filename string) (*Attachment, error)
```
AttachFile adds content from a filename to email attachment

#### func (*Message) Bytes

```go
func (msg *Message) Bytes() []byte
```
Bytes converts an e-mail subject, body, and attachments to bytes

#### type MessageSender

```go
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
```

MessageSender struct represents an email message sender

#### func (*MessageSender) Send

```go
func (sender *MessageSender) Send() error
```
Send is a sender pointer receiver to send email message

#### type SMTPClient

```go
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
```

SMTPClient interface

#### type Sender

```go
type Sender interface {
	Send() error
}
```

Sender interface represents a message sender

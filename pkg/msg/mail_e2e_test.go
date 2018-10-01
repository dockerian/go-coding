// +build e2e mail

// Package msg :: mail_e2e_test.go
package msg

import (
	"fmt"
	"net/smtp"
	"strconv"
	"strings"
	"testing"

	"github.com/dockerian/go-coding/pkg/cfg"
	"github.com/dockerian/go-coding/pkg/str"
	"github.com/stretchr/testify/assert"
)

var (
	config = cfg.Config{}

	toTest  = config.Get("mail.to", "jason_zhuyx@hotmail.com")
	ccTest  = config.Get("mail.to", "jason_zhuyx@hotmail.com")
	bccTest = config.Get("mail.to", "jason_zhuyx@hotmail.com")
	host    = config.Get("mail.host", "smtp.office365.com") // "smtp.gmail.com"
	user    = config.Get("mail.user", "jason_zhuyx@hotmail.com")
	pass    = config.Get("mail.pass", "********")
	useTLS  = config.Get("mail.tls", "true")

	recipients = map[string][]string{
		"from": {user},
		"to":   {toTest},
		"cc":   {ccTest},
		"bcc":  {bccTest},
	}

	recipientFmt = "<br/>Recipients (in config):<br/>\n<pre>%s</pre>\n"
	recipientAll = fmt.Sprintf(recipientFmt, str.IndentJSON(recipients, "  "))

	mailBodyHtml = `<h1>This is an email header</h1>
                  <div style="font-weight:bold;font-size:larger;">
                    <br/>**** <b>Please IGNORE this e-mail</b> *****
                  </div>
                  <p> and mail body content with a few more lines
                      and more and more</p>`

	message = Message{
		From:    recipients["from"][0],
		To:      recipients["to"],
		Cc:      recipients["cc"],
		Bcc:     recipients["bcc"],
		Subject: "go-coding: MessageSender test",
		Body:    []byte(fmt.Sprintf("%s\n%s", mailBodyHtml, recipientAll)),
	}

	smtpAuth = smtp.PlainAuth("", "username", "password", "host")
)

// TestMessageSender
// DevNotes:
// - Use `MAIL_TO` environment variable to set recipient.
// - Optionally use `MAIL_USER` and `MAIL_PASSWORD` to set sender.
// - Default `MAIL_USER` is "threatintelligence@infoblox.com".
// - Default `MAIL_PASSWORD` is from var mailPassword.
func TestMessageSender(t *testing.T) {
	mailUser := user
	if host != "smtp.office365.com" {
		mailUser = strings.Split(user, "@")[0]
	}
	domainName := strings.Split(user, "@")[1]
	withTLS, _ := strconv.ParseBool(useTLS)
	sender := MessageSender{
		Message:    message,
		DomainName: domainName,
		ServerHost: host,
		ServerPort: uint32(587),
		UserName:   mailUser,
		Password:   pass,
		PlainAuth:  false,
		WithTLS:    withTLS,
	}

	t.Logf("[e2e] Testing MessageSender: %+v\n", sender)
	err := sender.Send()
	if err != nil {
		t.Logf("[e2e] Testing MessageSender err: %s\n", err.Error())
	}
	assert.Nil(t, err)
}

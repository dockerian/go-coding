// Package main :: x.go - examples and tests
// Note: any function in `main` package but not in `main.go` seems invisible by `main()`
package main

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/dockerian/go-coding/pkg/msg"
	"github.com/dockerian/go-coding/pkg/zip"
)

// ExamplesTest runs any functional test for this project
func ExamplesTest() {
	message := msg.Message{
		To:      []string{"to@mail.com"},
		Cc:      []string{"cc@mail.com"},
		Bcc:     []string{"bcc@mail.com"},
		Body:    []byte("test"),
		Subject: "go-coding/pkg/msg test 1",
	}

	sources := []*zip.Source{
		{
			Reader: strings.NewReader("attachment test 1"), Name: "test1.txt",
		},
		{
			Reader: strings.NewReader("attachment test 2"), Name: "test2.txt",
		},
	}

	var buffer bytes.Buffer
	if err := zip.CreateZip(sources, &buffer); err != nil {
		fmt.Println("zip.CreateZip error:", err)
	} else {
		message.Attach(&buffer, "test.zip", "application/zip")
	}

	// log.Println("----zip content----:\n", string(buffer.Bytes()))
	message.From = "threatintelligence@infoblox.com"
	sender := msg.MessageSender{
		Message:    message,
		DomainName: "gmail.com",
		ServerHost: "smtp.office365.com",
		ServerPort: 587,
		UserName:   "dockeria@gmail.com",
		Password:   "password",
		WithTLS:    true,
	}

	fmt.Printf("sending mail by %+v\n", sender)
	fmt.Println(sender.Send())
}

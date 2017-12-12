// Package msg :: mail_auth.go - smtp.Auth implementations
// see
//   - http://www.samlogic.net/articles/smtp-commands-reference-auth.htm
//   - https://golang.org/src/net/smtp/auth.go
package msg

import (
	"errors"
	"fmt"
	"log"
	"net/smtp"
	"strings"
)

type loginAuth struct {
	username, password, host string
}

// LoginAuth returns an smtp.Auth implementation for LOGIN authentication
// as defined in RFC 4616
func LoginAuth(username, password, host string) smtp.Auth {
	return &loginAuth{username, password, host}
}

// Start implements smtp.Auth
func (auth *loginAuth) Start(server *smtp.ServerInfo) (string, []byte, error) {
	if !server.TLS {
		advertised := false
		for _, mechanism := range server.Auth {
			if mechanism == "LOGIN" {
				advertised = true
				break
			}
		}
		if !advertised {
			return "", nil, errors.New("unencrypted connection")
		}
	}
	if auth.host != server.Name {
		err := fmt.Errorf("auth host '%s' mismatch server name '%s'", auth.host, server.Name)
		log.Println("[smtp]", err)
		return "", nil, err
	}
	return "LOGIN", nil, nil
}

// Next implements smtp.Auth
func (auth *loginAuth) Next(fromServer []byte, more bool) ([]byte, error) {
	var command = string(fromServer)
	command = strings.TrimSpace(command)
	command = strings.TrimSuffix(command, ":")
	command = strings.ToLower(command)

	if more {
		if command == "username" {
			log.Printf("[smtp] LOGIN username: %v\n", auth.username)
			return []byte(fmt.Sprintf("%s", auth.username)), nil
		} else if command == "password" {
			log.Printf("[smtp] LOGIN password: ********\n")
			return []byte(fmt.Sprintf("%s", auth.password)), nil
		} else {
			// had already sent everything
			err := fmt.Errorf("unexpected server command: %s", command)
			log.Printf("[smtp] %v\n", err)
			return nil, err
		}
	}
	return nil, nil
}

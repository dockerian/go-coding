// +build all common pkg msg mail auth

// Package msg :: mail_auth_test.go
package msg

import (
	"net/smtp"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestLoginAuthStart
func TestLoginAuthStart(t *testing.T) {
	host := "mailhost"
	for idx, test := range []struct {
		serverInfo *smtp.ServerInfo
		expected   string
	}{
		{
			serverInfo: &smtp.ServerInfo{Name: "s1", TLS: false, Auth: []string{}},
			expected:   "unencrypted connection",
		},
		{
			serverInfo: &smtp.ServerInfo{Name: "s1", TLS: false, Auth: []string{"LOGIN"}},
			expected:   "auth host 'mailhost' mismatch server name 's1'",
		},
		{
			serverInfo: &smtp.ServerInfo{Name: "mailhost", TLS: true, Auth: []string{}},
		},
	} {
		auth := LoginAuth("username", "password", host)
		t.Logf("Test %2d: %+v => %s\n", idx+1, test.serverInfo, test.expected)
		result, data, err := auth.Start(test.serverInfo)
		assert.Nil(t, data)
		if test.expected != "" {
			assert.NotNil(t, err)
			assert.Equal(t, test.expected, err.Error())
			assert.Equal(t, "", result)
		} else {
			assert.Equal(t, "LOGIN", result)
			assert.Nil(t, err)
		}
	}
}

// TestLoginAuthNext
func TestLoginAuthNext(t *testing.T) {
	host := "smtp.gmail.com"
	for idx, test := range []struct {
		more              bool
		command, expected string
		expectError       bool
	}{
		{
			true, "Username: ", "username", false,
		},
		{
			true, "  PASSWORD: ", "password", false,
		},
		{
			true, "Unknown ", "unexpected server command: unknown", true,
		},
		{
			false, "command: ", "", false,
		},
	} {
		auth := LoginAuth("username", "password", host)
		t.Logf("Test %2d: %+v => %s\n", idx+1, test.command, test.expected)
		result, err := auth.Next([]byte(test.command), test.more)
		if test.expectError {
			assert.Nil(t, result)
			assert.NotNil(t, err)
			assert.Equal(t, test.expected, err.Error())
		} else if test.more {
			assert.Equal(t, test.expected, string(result))
		} else {
			assert.Nil(t, result)
			assert.Nil(t, err)
		}
	}

}

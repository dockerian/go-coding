// +build all pkg net ipaddress

// Package net :: ipaddress_test.go
package net

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsPublic(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
		err      string
	}{
		{"", false, "invalid IPv4 format: "},
		{".", false, "invalid IPv4 format: ."},
		{"1.1", false, "invalid IPv4 format: 1.1"},
		{"1.1.333.3", false, "invalid IPv4 format: 1.1.333.3"},
		{"1.1.1.1", true, ""},
		{"10.11.12.13", false, ""},
		{"172.16.255.255", false, ""},
		{"172.31.255.255", false, ""},
		{"192.168.1.1", false, ""},
	}
	for index, test := range tests {
		result, err := IsPublic(test.input)
		msg := fmt.Sprintf("expecting '%v' => '%v', '%v'", test.input, test.expected, test.err)
		t.Logf("Test %v: %v\n", index+1, msg)
		assert.Equal(t, test.expected, result, msg)
		if test.err != "" {
			assert.Equal(t, test.err, err.Error(), msg)
		} else {
			assert.Nil(t, err, msg)
		}
	}
}

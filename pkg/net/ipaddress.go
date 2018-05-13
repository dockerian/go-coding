// Package net :: ipaddress.go
package net

import (
	"fmt"
	"strconv"
	"strings"
)

// IsPublic returns true if the ipaddress is a public IP address.
func IsPublic(ipaddress string) (bool, error) {
	addresses := strings.Split(ipaddress, ".")

	errFormat := fmt.Errorf("invalid IPv4 format: %s", ipaddress)
	if len(addresses) != 4 {
		return false, errFormat
	}
	var ipFields []int
	for _, v := range addresses {
		i, err := strconv.Atoi(v)
		if err != nil || i < 0 || i > 255 {
			return false, errFormat
		}
		ipFields = append(ipFields, i)
	}

	isPrivate := ipFields[0] == 172 && ipFields[1] >= 16 && ipFields[1] <= 31 ||
		strings.HasPrefix(ipaddress, "192.168.") ||
		strings.HasPrefix(ipaddress, "10.")

	return !isPrivate, nil
}

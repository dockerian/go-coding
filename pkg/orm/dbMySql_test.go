// +build all common pkg orm mysql

// Package orm :: orm_test.go

package orm

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestOpenMySQL tests orm.OpenMySQL function
func TestOpenMySQL(t *testing.T) {
	openDB = mockGormOpen
	db, err := OpenMySQL("host", "", "db", "user", "pass")
	assert.Nil(t, err)
	assert.Equal(t, map[string]string{
		"dialect": "mysql",
		"options": "user:pass@tcp(host)/db",
	}, db.Value)

	db, err = OpenMySQL("host", "port", "db", "user", "pass", "")
	assert.Nil(t, err)
	assert.Equal(t, map[string]string{
		"dialect": "mysql",
		"options": "user:pass@tcp(host:port)/db",
	}, db.Value)

	db, err = OpenMySQL("host", "port", "db", "user", "pass", "p1=v1", "p2=v2", "p3=v3")
	assert.Nil(t, err)
	assert.Equal(t, map[string]string{
		"dialect": "mysql",
		"options": "user:pass@tcp(host:port)/db?p1=v1&p2=v2&p3=v3",
	}, db.Value)
}

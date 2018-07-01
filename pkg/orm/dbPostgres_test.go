// +build all common pkg orm postgres

// Package orm :: orm_test.go

package orm

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestOpenPostgres tests orm.OpenPostgres function
func TestOpenPostgres(t *testing.T) {
	openDB = mockGormOpen
	tests := []struct {
		host     string
		port     string
		dbName   string
		user     string
		pass     string
		options  []string
		expected string
	}{
		{
			"host", "", "db", "user", "pass", []string{},
			"host=host port=5432 user=user password=pass dbname=db sslmode=disable",
		},
		{
			"host", "port", "db", "user", "pass", []string{""},
			"host=host port=5432 user=user password=pass dbname=db sslmode=disable",
		},
		{
			"host", "port", "db", "user", "pass", []string{"sslmode=require"},
			"host=host port=5432 user=user password=pass dbname=db sslmode=require",
		},
		{
			"host", "3456", "db", "user", "pass", []string{"p1=v1", "p2=v2", "p3=v3"},
			"host=host port=3456 user=user password=pass dbname=db p1=v1 p2=v2 p3=v3 sslmode=disable",
		},
	}

	for idx, test := range tests {
		t.Logf("Test %2d: %+v\n", idx+1, test)
		db, err := OpenPostgres(
			test.host, test.port, test.dbName, test.user, test.pass, test.options...)
		assert.Nil(t, err)
		assert.Equal(t, map[string]string{
			"dialect": "postgres",
			"options": test.expected,
		}, db.Value)
	}
}

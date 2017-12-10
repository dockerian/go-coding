// +build all common pkg orm

// Package orm :: orm_test.go

package orm

import (
	"testing"

	"github.com/dockerian/go-coding/pkg/api"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"
)

var (
	testDB = &gorm.DB{}
	params = &api.Params{
		Form: map[string][]string{
			"date": []string{
				"2017-11-11 11:01:01",
			},
			"dates": []string{
				"2017-11-11 11:01:01",
				"2009-11-22T11:22:02",
				"2009-11-30T23:33:03",
			},
			"datetime-updated": []string{
				"one",
				"2017-33-33 33:44:55,yyyy-mm-dd",
				"2017-11-30,",
			},
			"dates-range": []string{
				"some invalid date",
				"2017-11-11 11:01:01,2009-11-22T11:22:02",
				"2017-11-30,2009-11-30T23:33:03",
			},
			"debug": []string{
				"1",
			},
			"key": []string{
				"value%",
			},
			"pgNeg": []string{
				"-5",
			},
			"pgOffset": []string{
				"5",
			},
			"pgSize": []string{
				"20",
			},
			"order": []string{
				"field desc,name",
			},
			"search": []string{
				"value",
			},
			"name": []string{
				"value",
			},
			"names": []string{
				"value1",
				"value2",
				"value3",
			},
		},
	}
)

// mockOpenMySQL mocks gorm.Open
func mockOpenMySQL(dialect string, args ...interface{}) (*gorm.DB, error) {
	options := ""
	if len(args) > 0 {
		options = args[0].(string)
	}
	value := map[string]string{"dialect": dialect, "options": options}
	db := &gorm.DB{
		Value: value,
	}
	return db, nil
}

// TestOpenMySQL tests orm.OpenMySQL function
func TestOpenMySQL(t *testing.T) {
	openDB = mockOpenMySQL
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

// TestGetClauseByParams tests orm.GetClauseByParams
func TestGetClauseByParams(t *testing.T) {
	db0 := GetClauseByParams(testDB, params, "foo", "bar")
	assert.Equal(t, testDB, db0)
	db1 := GetClauseByParams(testDB, params, "name", "name")
	assert.NotEqual(t, testDB, db1)
	db2 := GetClauseByParams(testDB, params, "names", "name")
	assert.NotEqual(t, testDB, db2)
}

// TestGetDateClause tests orm.GetDateClause
func TestGetGetDateClause(t *testing.T) {

}

// TestGetDateClauseByParams tests orm.GetDateClauseByParams
func TestGetDateClauseByParams(t *testing.T) {
	db0 := GetDateClauseByParams(testDB, params, "foo", "bar")
	assert.Equal(t, testDB, db0)
	db1 := GetDateClauseByParams(testDB, params, "date", "field")
	assert.NotEqual(t, testDB, db1)
	db2 := GetDateClauseByParams(testDB, params, "dates", "field")
	assert.NotEqual(t, testDB, db2)
	db3 := GetDateClauseByParams(testDB, params, "datetime-updated", "field")
	assert.NotEqual(t, testDB, db3)
	db4 := GetDateClauseByParams(testDB, params, "dates-range", "field")
	assert.NotEqual(t, testDB, db4)
}

// TestGetLikeClauseByParams tests orm.GetLikeClauseByParams
func TestGetLikeClauseByParams(t *testing.T) {
	db0 := GetLikeClauseByParams(testDB, params, "foo", "bar")
	assert.Equal(t, testDB, db0)
	db1 := GetLikeClauseByParams(testDB, params, "debug", "field")
	assert.NotEqual(t, testDB, db1)
	db2 := GetLikeClauseByParams(testDB, params, "name", "field")
	assert.NotEqual(t, testDB, db2)
	db3 := GetLikeClauseByParams(testDB, params, "search", "field")
	assert.NotEqual(t, testDB, db3)
	db4 := GetLikeClauseByParams(testDB, params, "key", "field")
	assert.NotEqual(t, testDB, db4)
}

// TestGetOrderClauseByParams tests orm.GetOrderClauseByParams
func TestGetOrderClauseByParams(t *testing.T) {
	db0 := GetOrderClauseByParams(testDB, params, "foo")
	assert.Equal(t, testDB, db0)
	db1 := GetOrderClauseByParams(testDB, params, "order")
	assert.NotEqual(t, testDB, db1)
	db2 := GetOrderClauseByParams(testDB, params, "names")
	assert.NotEqual(t, testDB, db2)
}

// TestGetPageClauseByParams tests orm.GetPageClauseByParams
func TestGetPageClauseByParams(t *testing.T) {
	db0, sz, pg := GetPageClauseByParams(testDB, params, "foo", "pgOffset")
	assert.Equal(t, testDB, db0)
	assert.Zero(t, pg)
	assert.Zero(t, sz)
	db1, sz, pg := GetPageClauseByParams(testDB, params, "pgSize", "pgNeg")
	assert.NotEqual(t, testDB, db1)
	assert.Equal(t, 20, sz)
	assert.Zero(t, pg)
	db2, sz, pg := GetPageClauseByParams(testDB, params, "pgSize", "pgOffset")
	assert.NotEqual(t, testDB, db2)
	assert.Equal(t, 20, sz)
	assert.Equal(t, 5, pg)
}

// Package orm :: orm.go - extended ORM wrapper functions
package orm

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/dockerian/go-coding/api/api"
	"github.com/jinzhu/gorm"
	// golint should ignore the following blank import
	_ "github.com/jinzhu/gorm/dialects/mysql" // required by gorm
)

const (
	// MaximumPageSize defines the maximum page size for LIMIT clause
	MaximumPageSize = 500
	// MinimumPageSize defines the minimum page size for LIMIT clause
	MinimumPageSize = 5
	// MinimumSearchLength defines the minimum length of search string
	MinimumSearchLength = 3
)

var (
	openDB = gorm.Open
)

// gormDB provides interface for gorm.DB
type gormDB interface {
	Where(interface{}, ...interface{}) *gorm.DB
}

// GetClauseByParams returns string comparison clause from params
func GetClauseByParams(db *gorm.DB, params *api.Params, key, field string) *gorm.DB {
	strValues := params.GetValues(key)

	if siz := len(strValues); siz > 0 && strValues[0] != "" {
		if siz > 1 {
			return db.Where(fmt.Sprintf("%s in (?)", field), strValues)
		}
		return db.Where(fmt.Sprintf("%s = ?", field), strValues[0])
	}
	return db
}

// GetDateClause returns database where clause from a sorted time slice
func GetDateClause(db *gorm.DB, field string, dateValues ...time.Time) *gorm.DB {
	siz := len(dateValues)
	if siz > 0 {
		strValues := make([]string, siz)
		for i, dateValue := range dateValues {
			strValues[i] = dateValue.Format("2006-01-02")
		}
		if siz == 1 {
			db = db.Where(fmt.Sprintf("date(%s) = ?", field), strValues[0])
		} else {
			db = db.Where(fmt.Sprintf("date(%s) in (?)", field), strValues)
		}
	}
	return db
}

// GetDateRangeClause returns database with range clause from a sorted time slice
func GetDateRangeClause(db *gorm.DB, field string, dateValues ...time.Time) *gorm.DB {
	siz := len(dateValues)
	if siz > 0 {
		begin := dateValues[0].Format("2006-01-02")
		if siz > 1 {
			end := dateValues[siz-1].Format("2006-01-02")
			db = db.Where(fmt.Sprintf("date(%s) >= ? AND date(%s) <= ?", field, field), begin, end)
		} else {
			db = db.Where(fmt.Sprintf("date(%s) = ?", field), begin)
		}
	}
	return db
}

// GetDateClauseByParams returns database where clause from params
//
// Allowing 2 types of date queries for any key=value pairs in params
// - date range: ?key=2017-11-11,2017-11-30
// - date selections: ?key=2017-11-11&key=2017-11-20&key=2017-11-30
//
// Note: The date selections (IN clause) takes preference;
//       Otherwise, for date range format, all comma-delimited values will be
//       parsed and sorted, so that the first and last dates define the range.
func GetDateClauseByParams(db *gorm.DB, params *api.Params, key, field string) *gorm.DB {
	// checking date selection params
	if dateValues, err := params.GetDateValues(key); err == nil {
		return GetDateClause(db, field, dateValues...)
	}

	// checking date range params
	if dateValues, err := params.GetDateRange(key); err == nil {
		return GetDateRangeClause(db, field, dateValues...)
	}

	return db
}

// GetLikeClauseByParams returns LIKE comparison clause from params
func GetLikeClauseByParams(db *gorm.DB, params *api.Params, key, field string) *gorm.DB {
	if strValue := params.GetValue(key); strValue != "" {
		if len(strValue) > MinimumSearchLength {
			if !strings.Contains(strValue, "%") {
				strValue = "%" + strValue + "%"
			}
			return db.Where(fmt.Sprintf("%s LIKE ?", field), strValue)
		}
		return db.Where(fmt.Sprintf("%s = ?", field), strValue)
	}
	return db
}

// GetOrderClauseByParams returns ORDER BY clause by params
func GetOrderClauseByParams(db *gorm.DB, params *api.Params, orderKey string) *gorm.DB {
	orders := params.GetValues(orderKey)
	for _, order := range orders {
		if order != "" {
			db = db.Order(order)
		}
	}
	return db
}

// GetPageClauseByParams returns LIMIT and OFFSET clause by params
func GetPageClauseByParams(db *gorm.DB, params *api.Params, pgSizeKey, pgOffsetKey string) (*gorm.DB, int, int) {
	var pgSize, pgOffset int
	if pgSize = params.GetIntByRange(pgSizeKey, 0, MaximumPageSize); pgSize >= MinimumPageSize {
		db = db.Limit(pgSize)
		if pgOffset = params.GetIntByRange(pgOffsetKey, 0); pgOffset > 0 {
			db = db.Offset(pgOffset * pgSize)
		}
	}
	return db, pgSize, pgOffset
}

// OpenMySQL wraps gorm.Open function to open a mysql db connection with
// host, port, database name, user name, password, and options
// Note: options allow to use following case-sensitive parameters:
//    charset (e.g. charset=utf8, default: none)
//    collation (default: utf8_general_ci)
//    columnsWithAlias (default: false)
//    loc (default: UTC)
//    maxAllowedPacket (default: 4194304)
//    multiStatements (default: false)
//    parseTime (default: false, changing DATE or DATETIME values to time.Time)
//    readTimeout (default: 0) - a unit suffix ("ms", "s", "m", "h"), such as "30s", "0.5m" or "1m30s"
//    timeout (default: OS default)
// Parameters can be joined by amphersand (e.g. "charset=utf8&parseTime=true")
// See https://github.com/go-sql-driver/mysql#parameters
func OpenMySQL(host, port, db, user, pass string, options ...string) (*gorm.DB, error) {
	address := host
	if port != "" {
		address += ":" + port
	}
	parameters := ""
	for _, opt := range options {
		if opt == "" {
			continue
		}
		if parameters == "" {
			parameters += "?" + opt
		} else {
			parameters += "&" + opt
		}
	}
	connLog := fmt.Sprintf("%s:%s@tcp(%s)/%s%s", user, "********", address, db, parameters)
	conn := fmt.Sprintf("%s:%s@tcp(%s)/%s%s", user, pass, address, db, parameters)
	log.Println("[mysql] connect to", connLog)
	return openDB("mysql", conn)
}

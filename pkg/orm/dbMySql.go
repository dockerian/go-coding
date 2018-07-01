// Package orm :: orm.go - extended ORM wrapper functions
package orm

import (
	// golint should ignore the following blank import
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql" // required by gorm
)

// OpenMySQL wraps gorm.Open function to open a mysql db connection with
// host, port, database name, user name, password, and options.
//
// Note: options allow to use following case-sensitive parameters:
//
// * `charset` (e.g. `charset=utf8`, default: none)
//
// * `collation` (default: `utf8_general_ci`)
//
// * `columnsWithAlias` (default: `false`)
//
// * `loc` (default: `UTC`)
//
// * `maxAllowedPacket` (default: `4194304`)
//
// * `multiStatements` (default: `false`)
//
// * `parseTime` (default: `false`, changing DATE or DATETIME values to time.Time)
//
// * `readTimeout` (default: 0) - a unit suffix ("ms", "s", "m", "h"), such as "30s", "0.5m" or "1m30s"
//
// * `timeout` (default: OS default)
//
// Parameters are joined by amphersand (e.g. "charset=utf8&parseTime=true")
//
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

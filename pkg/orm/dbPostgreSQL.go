// Package orm :: orm.go - extended ORM wrapper functions
package orm

import (
	// golint should ignore the following blank import
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" // required by gorm
)

// OpenPostgres wraps gorm.Open function to open a postgreSQL db connection
// with host, port, database name, user name, password, and options.
//
// Valid connection options:
//
// * `sslmode` - Whether or not to use SSL (default is `require`); valid values:
//     * `disable` - No SSL
//     * `require` - Always SSL (skip verification)
//     * `verify-ca` - Always SSL (verify that the certificate
//       presented by the server was signed by a trusted CA)
//     * `verify-full` - Always SSL (verify that the certification
//       presented by the server was signed by a trusted CA and the
//       server host name matches the one in the certificate)
//
// * `fallback_application_name` - An application_name to fall back if not provided.
//
// * `connect_timeout` - Maximum wait for connection, in seconds.
//                       Zero or not specified means wait indefinitely.
//
// * `sslcert` - Cert file location. The file must contain PEM encoded data.
//
// * `sslkey` - Key file location. The file must contain PEM encoded data.
//
// * `sslrootcert` - The location of the root certificate file.
//                   The file must contain PEM encoded data.
//
// Parameters are joined by space (e.g. "sslmode=disable connect_timeout=30")
//
// See https://godoc.org/github.com/lib/pq#hdr-Connection_String_Parameters
func OpenPostgres(host, port, db, user, pass string, options ...string) (*gorm.DB, error) {
	parameters := ""
	for _, opt := range options {
		if opt == "" {
			continue
		}
		parameters += " " + opt
	}
	if !strings.Contains(parameters, "sslmode=") {
		parameters += " sslmode=disable"
	}
	validPort, err := strconv.Atoi(port)
	if err != nil {
		validPort = 5432
	}

	connLog := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s%s",
		host, validPort, user, "********", db, parameters)
	conn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s%s",
		host, validPort, user, pass, db, parameters)
	log.Println("[postgres] connect to", connLog)
	return openDB("postgres", conn)
}

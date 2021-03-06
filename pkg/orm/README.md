# orm
--
    import "github.com/dockerian/go-coding/pkg/orm"

Package orm :: orm.go - extended ORM wrapper functions

Package orm :: orm.go - extended ORM wrapper functions

Package orm :: orm.go - extended ORM wrapper functions

## Usage

```go
const (
	// MaximumPageSize defines the maximum page size for LIMIT clause
	MaximumPageSize = 500
	// MinimumPageSize defines the minimum page size for LIMIT clause
	MinimumPageSize = 5
	// MinimumSearchLength defines the minimum length of search string
	MinimumSearchLength = 3
)
```

#### func  GetClauseByParams

```go
func GetClauseByParams(db *gorm.DB, params *api.Params, key, field string) *gorm.DB
```
GetClauseByParams returns string comparison clause from params

#### func  GetDateClause

```go
func GetDateClause(db *gorm.DB, field string, dateValues ...time.Time) *gorm.DB
```
GetDateClause returns database where clause from a sorted time slice

#### func  GetDateClauseByParams

```go
func GetDateClauseByParams(db *gorm.DB, params *api.Params, key, field string) *gorm.DB
```
GetDateClauseByParams returns database where clause from params.

Allowing 2 types of date queries for any key=value pairs in params.

    - date range: ?key=2017-11-11,2017-11-30
    - date selections: ?key=2017-11-11&key=2017-11-20&key=2017-11-30

Note: The date selections (IN clause) takes preference; otherwise, for date
range format, all comma-delimited values will be parsed and sorted, so that the
first and last dates define the range.

#### func  GetDateRangeClause

```go
func GetDateRangeClause(db *gorm.DB, field string, dateValues ...time.Time) *gorm.DB
```
GetDateRangeClause returns database with range clause from a sorted time slice

#### func  GetLikeClauseByParams

```go
func GetLikeClauseByParams(db *gorm.DB, params *api.Params, key, field string) *gorm.DB
```
GetLikeClauseByParams returns LIKE comparison clause from params

#### func  GetNumberClauseByParams

```go
func GetNumberClauseByParams(db *gorm.DB, params *api.Params, key, field string) *gorm.DB
```
GetNumberClauseByParams returns number comparison clause from params

#### func  GetOrderClauseByParams

```go
func GetOrderClauseByParams(db *gorm.DB, params *api.Params, orderKey string) *gorm.DB
```
GetOrderClauseByParams returns ORDER BY clause by params

#### func  GetPageClauseByParams

```go
func GetPageClauseByParams(db *gorm.DB, params *api.Params, pgSizeKey, pgOffsetKey string) (*gorm.DB, int, int)
```
GetPageClauseByParams returns LIMIT and OFFSET clause by params

#### func  OpenMySQL

```go
func OpenMySQL(host, port, db, user, pass string, options ...string) (*gorm.DB, error)
```
OpenMySQL wraps gorm.Open function to open a mysql db connection with host,
port, database name, user name, password, and options.

Note: options allow to use following case-sensitive parameters:

* `charset` (e.g. `charset=utf8`, default: none)

* `collation` (default: `utf8_general_ci`)

* `columnsWithAlias` (default: `false`)

* `loc` (default: `UTC`)

* `maxAllowedPacket` (default: `4194304`)

* `multiStatements` (default: `false`)

* `parseTime` (default: `false`, changing DATE or DATETIME values to time.Time)

* `readTimeout` (default: 0) - a unit suffix ("ms", "s", "m", "h"), such as
"30s", "0.5m" or "1m30s"

* `timeout` (default: OS default)

Parameters are joined by amphersand (e.g. "charset=utf8&parseTime=true")

See https://github.com/go-sql-driver/mysql#parameters

#### func  OpenPostgres

```go
func OpenPostgres(host, port, db, user, pass string, options ...string) (*gorm.DB, error)
```
OpenPostgres wraps gorm.Open function to open a postgreSQL db connection with
host, port, database name, user name, password, and options.

Valid connection options:

* `sslmode` - Whether or not to use SSL (default is `require`); valid values:

    * `disable` - No SSL
    * `require` - Always SSL (skip verification)
    * `verify-ca` - Always SSL (verify that the certificate
      presented by the server was signed by a trusted CA)
    * `verify-full` - Always SSL (verify that the certification
      presented by the server was signed by a trusted CA and the
      server host name matches the one in the certificate)

* `fallback_application_name` - An application_name to fall back if not
provided.

* `connect_timeout` - Maximum wait for connection, in seconds.

    Zero or not specified means wait indefinitely.

* `sslcert` - Cert file location. The file must contain PEM encoded data.

* `sslkey` - Key file location. The file must contain PEM encoded data.

* `sslrootcert` - The location of the root certificate file.

    The file must contain PEM encoded data.

Parameters are joined by space (e.g. "sslmode=disable connect_timeout=30")

See https://godoc.org/github.com/lib/pq#hdr-Connection_String_Parameters

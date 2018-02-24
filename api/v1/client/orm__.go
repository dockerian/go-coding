// Package client :: orm__.go - extended definitions and relationships
package client

// TableName sets DbSchema's table name to be `db_schema`
func (*DbSchema) TableName() string {
	return "db_schema"
}

/*
 * Go-coding API
 *
 * Go-coding API Gateway.
 *
 * OpenAPI spec version: 0.0.1
 * Contact: jason.zhuy@gmail.com
 * Generated by: https://github.com/swagger-api/swagger-codegen.git
 */

package client

// DbSchema provides detailed info for database releases.
type DbSchema struct {

	// ID represents a schema identifier.
	Id int32 `json:"id,omitempty"`

	// DbVersion is the database schema version.
	DbVersion string `json:"dbVersion,omitempty"`

	// ChangeDate is a date/time string of the schema change.
	ChangeDate string `json:"changeDate,omitempty"`

	// Author (in format of 'Name <email.address>') of database schema change.
	Author string `json:"author,omitempty"`

	// Description of the schema change.
	Description string `json:"description,omitempty"`

	// DeployNotes is a special notes for database deployment.
	DeployNotes string `json:"deployNotes,omitempty"`

	// Scriptis the SQL script file name for database deployment.
	Script string `json:"script,omitempty"`
}

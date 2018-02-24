// Package db :: db.go - MySql database functions
package db

import (
	"errors"
	"fmt"
	"log"

	"github.com/dockerian/go-coding/api/v1/client"
	"github.com/dockerian/go-coding/pkg/cfg"
	"github.com/dockerian/go-coding/pkg/orm"
	"github.com/jinzhu/gorm"
)

var (
	// errDBOpen set to db connection error
	errDBOpen = errors.New("cannot connect to database")
	// ormOpenMySQL set to orm.OpenMySQL
	ormOpenMySQL = orm.OpenMySQL
)

func openMySQLDB(ctx cfg.Context) *gorm.DB {
	db, err := OpenMySQLDB(ctx)
	if err != nil {
		log.Printf("[ERROR] %v\n", err)
		return nil
	}
	if ctx.Env.Get("debug") == "true" {
		log.Println("[db] DEBUG is ON.")
		return db.Debug()
	}
	return db
}

// OpenMySQLDB returns gorm.DB connection for MySql database
func OpenMySQLDB(ctx cfg.Context) (*gorm.DB, error) {
	env := ctx.Env
	if env == nil {
		return nil, fmt.Errorf("ctx does not have Env")
	}

	name := env.Get("mysql.db.name")
	host := env.Get("mysql.db.host")
	port := env.Get("mysql.db.port")
	user := env.Get("mysql.db.user")
	pass := env.Get("mysql.db.pass")
	opts := "charset=utf8&parseTime=True&loc=Local"

	log.Printf("[db] opening db: '%s:%s/%s'\n", host, port, name)
	log.Printf("        by user: '%s' (options: %s)\n", user, opts)

	return ormOpenMySQL(host, port, name, user, pass, opts)
}

// SchemaInfo returns the latest db schema info
func SchemaInfo(ctx cfg.Context) *client.DbSchema {
	dbInfo := client.DbSchema{}

	if db := openMySQLDB(ctx); db != nil {
		defer db.Close()
		log.Println("[db_schema] reading the latest db schema ...")
		db.Last(&dbInfo)
	}

	return &dbInfo
}

// SchemaInfoAll returns all db schema info
func SchemaInfoAll(ctx cfg.Context) []client.DbSchema {
	dbInfoAll := []client.DbSchema{}

	if db := openMySQLDB(ctx); db != nil {
		defer db.Close()
		log.Println("[db_schema] reading all db schema history ...")
		db.Find(&dbInfoAll)
	}

	return dbInfoAll
}

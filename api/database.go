package main

import (
	"database/sql"
	"fmt"

	_ "github.com/denisenkom/go-mssqldb"
)

func dbInit() {
	var server = configuration.Db.Server
	var user = configuration.Db.Username
	var password = configuration.Db.Password
	var database = configuration.Db.Database
	configuration.Db.Dsn = fmt.Sprintf("server=%s;user id=%s;password=%s;database=%s", server, user, password, database)
}

var dbQuery = func(query string) (*sql.Rows, error) {
	db, err := sql.Open("sqlserver", configuration.Db.Dsn)
	checkErr(err)
	defer db.Close()
	return db.Query(query)
}

package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func dbInit() {
	var user = configuration.Db.Username
	var password = configuration.Db.Password
	var dbname = configuration.Db.Name
	configuration.Db.DbInfo = fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", user, password, dbname)
}

var dbQuery = func(query string) (*sql.Rows, error) {
	db, err := sql.Open("postgres", configuration.Db.DbInfo)
	checkErr(err)
	defer db.Close()
	return db.Query(query)
}

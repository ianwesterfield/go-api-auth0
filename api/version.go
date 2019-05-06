package main

import (
	"encoding/json"
	"errors"
	"net/http"

	models "./models"
)

var versionHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

	var getVersion = func() (int, error) {
		rows, err := dbQuery("select version from version limit 1")
		checkErr(err)

		var version int

		if rows.Next() {
			rows.Scan(&version)
		} else {
			return -1, errors.New("no results")
		}

		return version, nil
	}

	version, err := getVersion()
	checkErr(err)

	s := &models.Status{
		Authorized: true,
		Version:    version,
	}

	payload, err := json.Marshal(s)
	checkErr(err)

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(payload))
})

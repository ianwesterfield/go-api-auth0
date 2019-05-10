package controllers

import (
	"encoding/json"
	"net/http"

	models "../models"
)

// SendMessage -
var SendMessage = func(w http.ResponseWriter, r *http.Request) {
	s := &models.APIResponse{
		Message: "message sent",
		Body:    `{"totalItem":0, "items": []}`,
	}

	payload, err := json.Marshal(s)
	checkErr(err)

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(payload))
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

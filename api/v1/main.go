package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

var configuration Configuration

func main() {
	// get configuration
	file, err := os.Open("./main.config.json")
	checkErr(err)

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&configuration)
	checkErr(err)

	// setup database
	dbInit()

	// setup router
	r := mux.NewRouter()
	r.Handle("/api/v1/version", versionHandler).Methods("GET")

	// secure all routes
	securedRoutes := auth0(r)

	// handle CORS
	c := cors.New(cors.Options{
		AllowedOrigins:   configuration.Cors.AllowedOrigins,
		AllowCredentials: configuration.Cors.AllowCredentials,
		Debug:            configuration.Cors.Debug,
		AllowedHeaders:   configuration.Cors.AllowedHeaders,
	})
	corsHandler := c.Handler(securedRoutes)

	// log all requests
	loggingHandler := handlers.LoggingHandler(os.Stdout, corsHandler)

	// start listening
	http.ListenAndServe(fmt.Sprintf(":%s", configuration.Server.Port), loggingHandler)
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

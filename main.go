package main

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "test"
	dbname   = "testdb"

	handlerPort = "3000"
)

var someError = errors.New("some error")

func main() {
	fmt.Printf("started\n")
	dbClient, err := dbSetup()
	if err != nil {
		panic(err)
	}

	// Set up mux router
	router := mux.NewRouter()
	// Set up routes
	router.HandleFunc("/dbconfig", dbconfig(dbClient)).Methods(http.MethodGet)

	log.Printf("Starting service on http://localhost:%s\n", handlerPort)
	log.Fatal(http.ListenAndServe(":"+handlerPort, router))
}

func dbconfig(dbClient *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		httpStatusCode := http.StatusOK
		var err error

		resp := map[string]interface{}{
			"status": "Ok",
			"data": Response{
				Data: dbConnectionString(),
			},
		}

		writeResponse(w, resp, httpStatusCode, err)
	}
}

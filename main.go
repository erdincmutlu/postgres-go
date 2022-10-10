package main

import (
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

func main() {
	fmt.Printf("started\n")
	dbClient, err := dbSetup()
	defer dbClient.Close()
	if err != nil {
		panic(err)
	}

	// Set up mux router
	router := mux.NewRouter()
	// Set up routes
	router.HandleFunc("/dbconfig", dbconfig(dbClient)).Methods(http.MethodGet)
	router.HandleFunc("/dbsize", dbsize(dbClient)).Methods(http.MethodGet)
	router.HandleFunc("/select", dbselect(dbClient)).Methods(http.MethodGet)

	log.Printf("Starting service on http://localhost:%s\n", handlerPort)
	log.Fatal(http.ListenAndServe(":"+handlerPort, router))
}

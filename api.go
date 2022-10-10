package main

import (
	"database/sql"
	"net/http"

	_ "github.com/lib/pq"
)

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

func dbsize(dbClient *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		httpStatusCode := http.StatusOK

		size, err := dbSize(dbClient)
		resp := map[string]interface{}{
			"status": "Ok",
			"data": Response{
				Data: size,
			},
		}

		writeResponse(w, resp, httpStatusCode, err)
	}
}

func dbselect(dbClient *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		httpStatusCode := http.StatusOK

		acc, err := dbSelect(dbClient)
		resp := map[string]interface{}{
			"status": "Ok",
			"data": Response{
				Data: acc.email,
			},
		}

		writeResponse(w, resp, httpStatusCode, err)
	}
}

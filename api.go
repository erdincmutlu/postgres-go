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

		resp := Response{
			Status:   "Ok",
			Endpoint: "/dbconfig",
			Data:     dbConnectionString(),
		}
		if err != nil {
			resp.Error = err.Error()
		}

		writeResponse(w, resp, httpStatusCode, err)
	}
}

func dbsize(dbClient *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		httpStatusCode := http.StatusOK

		size, err := dbSize(dbClient)
		resp := Response{
			Status:   "Ok",
			Endpoint: "/dbsize",
			Data:     size,
		}
		if err != nil {
			resp.Error = err.Error()
		}

		writeResponse(w, resp, httpStatusCode, err)
	}
}

func dbselect(dbClient *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		httpStatusCode := http.StatusOK

		acc, err := dbSelect(dbClient)
		resp := Response{
			Status:   "Ok",
			Endpoint: "/select",
			Data:     acc.email,
		}
		if err != nil {
			resp.Error = err.Error()
		}

		writeResponse(w, resp, httpStatusCode, err)
	}
}

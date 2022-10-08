package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func dbSetup() (*sql.DB, error) {
	psqlconn := dbConnectionString()
	dbClient, err := sql.Open("postgres", psqlconn)
	if err != nil {
		return nil, err
	}
	defer dbClient.Close()

	err = dbClient.Ping()
	if err != nil {
		return nil, err
	}

	fmt.Println("DB Connected!")
	return dbClient, nil
}

func dbConnectionString() string {
	// connection string
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	return psqlconn
}

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

func dbSize(db *sql.DB) (int, error) {
	query := "SELECT pg_database_size('testdb')"

	rows, err := db.Query(query)
	if err != nil {
		return 0, err
	}
	defer rows.Close()

	var size int
	for rows.Next() {
		err := rows.Scan(&size)
		if err != nil {
			panic(err)
		}
		fmt.Printf("%+v\n", size)
	}
	err = rows.Err()
	if err != nil {
		return 0, err
	}

	return size, nil
}

type account struct {
	id    int
	email string
	token string
}

func dbSelect(db *sql.DB) (*account, error) {
	query := "SELECT * from account"
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var acc account
	for rows.Next() {
		err := rows.Scan(&acc.id, &acc.email, &acc.token)
		if err != nil {
			panic(err)
		}
		fmt.Printf("%+v\n", acc)
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return &acc, nil
}

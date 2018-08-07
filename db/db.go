package db

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

var (
	DB *sql.DB
)

// Connect to postgres database
func Connect() (db *sql.DB) {
	var err error

	connStr := "password='postgres' dbname=lost sslmode=disable"
	DB, err = sql.Open("postgres", connStr)

	if err != nil {
		panic(err)
	}
	log.Println("Connect to postgress success")

	return DB
}


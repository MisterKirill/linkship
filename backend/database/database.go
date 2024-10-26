package database

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func ConnectDatabase() {
	var err error
	DB, err = sql.Open("postgres", "postgres://postgres:postgres@localhost/linksh?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
}

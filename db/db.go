package db

import (
	"database/sql"
	"log"
)

var db *sql.DB

// Setup PostgreSQL connection 
func Setup(connStr string) {
	db, _ = sql.Open("postgres", connStr)

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}
}

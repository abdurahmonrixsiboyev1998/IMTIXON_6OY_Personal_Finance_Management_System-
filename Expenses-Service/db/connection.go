package db

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func Connect(dbURI string) *sql.DB {
	db, err := sql.Open("postgres", dbURI)
	if err != nil {
		log.Fatal("Error opening database connection:", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal("Error connecting to the database:", err)
	}

	log.Println("Successfully connected to the database")
	return db
}
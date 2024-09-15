package db

import (
	"database/sql"
	"log"
)

func Migrate(db *sql.DB) {
	query := `
	CREATE TABLE IF NOT EXISTS transactions (
		transaction_id SERIAL PRIMARY KEY,
		amount FLOAT NOT NULL,
		currency VARCHAR(3) NOT NULL,
		category VARCHAR(50) NOT NULL,
		date DATE NOT NULL,
		type VARCHAR(10) NOT NULL
	);
	`

	if _, err := db.Exec(query); err != nil {
		log.Fatal("Error creating table:", err)
	}
}

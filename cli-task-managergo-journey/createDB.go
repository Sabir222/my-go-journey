package main

import (
	"database/sql"
	"log"
)

func CreateTaskDb(db *sql.DB) {
	query := `
	CREATE TABLE IF NOT EXISTS tasks (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	description TEXT,
	status CHAR(1) DEFAULT '✗' CHECK (status IN ('✗','✓')),
	date TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);`

	_, err := db.Exec(query)
	if err != nil {
		log.Fatal(err)
	}
}

package main

import (
	"database/sql"
	"fmt"
	"log"
)

func InsertTask(db *sql.DB, desc string) {
	query := `INSERT INTO tasks (description,status) VALUES (?,?)`
	_, err := db.Exec(query, desc, "✗")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("\nInserted task: %s\n", desc)
}

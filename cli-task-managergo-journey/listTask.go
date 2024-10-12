package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"
)

func ListTasks(db *sql.DB) {
	query := `SELECT * FROM tasks;`

	rows, err := db.Query(query)

	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	for rows.Next() {

		var id uint
		var description string
		var status string
		var date time.Time

		err = rows.Scan(&id, &description, &status, &date)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("\nid: %d | date: %s |description: %s | status: %s", id, date, description, status)
	}

	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}
}

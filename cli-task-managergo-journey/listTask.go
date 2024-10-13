package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"
)

func ListTasks(db *sql.DB) {
	const (
		Reset     = "\033[0m"
		Red       = "\033[31m"
		Green     = "\033[32m"
		Yellow    = "\033[33m"
		LightBlue = "\033[34;1m"
		Cyan      = "\033[36m"
	)

	query := `SELECT * FROM tasks;`

	rows, err := db.Query(query)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	fmt.Println()
	for rows.Next() {
		var id uint
		var description string
		var status string
		var date time.Time

		err = rows.Scan(&id, &description, &status, &date)
		if err != nil {
			log.Fatal(err)
		}

		// Determine color based on status
		var statusColor string
		switch status {
		case "✓":
			statusColor = Green
		case "✗":
			statusColor = Red
		default:
			statusColor = Yellow // Fallback color for unknown status
		}

		// Print colored output
		fmt.Printf("%sID: %d%s | %sDescription: %s%s | %sStatus: %s%s | %sDate: %s%s\n",
			LightBlue, id, Reset,
			Yellow, description, Reset,
			statusColor, status, Reset,
			Cyan, date.Format("2006-01-02 15:04:05"), Reset)
	}

	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}
}

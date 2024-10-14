package main

import (
	"database/sql"
	"fmt"
	"log"
)

func MarkAsDoneUnDone(db *sql.DB, id uint) {
	query := `
	UPDATE tasks
	SET status = CASE
		WHEN status = '✗' THEN '✓'
		WHEN status = '✓' THEN '✗'
		ELSE '✗'
	END
	WHERE id = ?;
	`
	result, err := db.Exec(query, id)
	if err != nil {
		log.Fatal(err)
	}

	rows, err := result.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}

	if rows == 0 {
		fmt.Printf("Task with id %d not found", id)
	} else {
		fmt.Printf("Task with id: %d modified successfully", id)
	}
}

func ModifyTask(db *sql.DB, newDesc string, id uint) {
	query := `
	UPDATE tasks
	SET description = ?
	WHERE id = ?;
	`
	result, err := db.Exec(query, newDesc, id)
	if err != nil {
		log.Fatal(err)
	}

	rows, err := result.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}

	if rows == 0 {
		fmt.Printf("Task with id %d not found", id)
	} else {
		fmt.Printf("Task with id: %d modified successfully", id)
	}

}

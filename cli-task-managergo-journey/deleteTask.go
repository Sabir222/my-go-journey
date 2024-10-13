package main

import (
	"database/sql"
	"fmt"
	"log"
)

func ResetDataBase(db *sql.DB) {
	query := "DELETE FROM tasks;"
	_, err := db.Exec(query)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("All tasks were deleted successfully")
}

func RemoveTask(db *sql.DB, id uint) {
	query := `DELETE FROM tasks WHERE id = ?`

	result, err := db.Exec(query, id)
	if err != nil {
		log.Fatal(err)
	}

	rows, err := result.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}

	if rows == 0 {
		fmt.Printf("Task with id %d not found\n", id)
	} else {
		fmt.Printf("Task with id: %d removed successfully\n", id)
	}
}

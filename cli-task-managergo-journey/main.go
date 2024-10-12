package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "./tasks.db")

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	// testing db

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	createTaskDb(db)

	reader := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("\nChoose Option : ")
		fmt.Println("1: Add task")
		fmt.Println("2: List task")
		fmt.Println("3: Update task")
		fmt.Println("4: Delete task")
		fmt.Println("5: Exit")

		reader.Scan()
		option := reader.Text()

		switch option {
		case "1":
			fmt.Println("\nEnter the task description")
			reader.Scan()
			description := reader.Text()
			insertTask(db, description)
		case "2":
			listTasks(db)
		case "3":
			fmt.Println("\nWhat do you want to update: ")
			fmt.Println("1: Mark as Complete/Incomplete")
			fmt.Println("2: Modify task")
			fmt.Println("")
		case "4":
			fmt.Println("\nEnter the id of the task or type 0 if you want to remove them all: ")
			reader.Scan()

			cmd := reader.Text()
			intCmd, err := strconv.Atoi(cmd)
			if err != nil {
				log.Fatal("Invalid input: ", err)
			}

			if intCmd == 0 {
				resetDataBase(db)
			}

			removeTask(db, uint(intCmd))
		case "5":
			return
		}
	}

}

func createTaskDb(db *sql.DB) {
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

	fmt.Println("Task table created successfully")
}

func insertTask(db *sql.DB, desc string) {
	query := `INSERT INTO tasks (description,status) VALUES (?,?)`
	_, err := db.Exec(query, desc, "✗")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("\nInserted task: %s\n", desc)
}

func resetDataBase(db *sql.DB) {
	query := "DELETE FROM tasks;"
	_, err := db.Exec(query)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("All tasks were deleted successfully")
}

func listTasks(db *sql.DB) {
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

		err = rows.Scan(&id, &description, &status)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("\nid: %d | description: %s | status: %s", id, description, status)
	}

	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}
}

func removeTask(db *sql.DB, id uint) {
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
		fmt.Printf("Task with id %d not found", id)
	} else {
		fmt.Printf("Task with id: %d removed successfully", id)
	}
}

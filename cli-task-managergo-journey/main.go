package main

import (
	"bufio"
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"os"
	"strconv"
)

func main() {
	db, err := sql.Open("sqlite3", "~/tasks.db")

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	// testing db

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	CreateTaskDb(db)
	ListTasks(db)
	reader := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("\n ")
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
			ClearCli()
			fmt.Println("\nEnter the task description")
			reader.Scan()
			description := reader.Text()
			ClearCli()
			InsertTask(db, description)
		case "2":
			ClearCli()
			ListTasks(db)
		case "3":
			ClearCli()
			fmt.Println("\nWhat do you want to update: ")
			fmt.Println("1: Modify task")
			fmt.Println("2: Mark as Complete/Incomplete")
			reader.Scan()
			option := reader.Text()
			if option == "1" {
				ClearCli()
				ListTasks(db)
				fmt.Println("\nPlease choose the id of task you want to modify: ")
				reader.Scan()
				id := reader.Text()
				intId, err := strconv.Atoi(id)
				if err != nil {
					log.Fatal("Invalid input: ", err)
				}
				ClearCli()
				fmt.Printf("\nYou choose id number %d", intId)
				fmt.Println("\nWhat's the new description ? ")
				reader.Scan()
				newDesc := reader.Text()
				ClearCli()
				ModifyTask(db, newDesc, uint(intId))

			} else if option == "2" {
				ClearCli()
				ListTasks(db)
				fmt.Println("What id of the task you want to mark as complete/incomplete ?")
				reader.Scan()
				id := reader.Text()
				intId, err := strconv.Atoi(id)
				if err != nil {
					log.Fatal("Invalid input: ", err)
				}
				ClearCli()
				MarkAsDoneUnDone(db, uint(intId))
			} else {
				fmt.Println("invalid option")
				return
			}

		case "4":
			ClearCli()
			ListTasks(db)
			fmt.Println("\nEnter the id of the task or type 0 if you want to remove them all: ")
			reader.Scan()

			cmd := reader.Text()
			intCmd, err := strconv.Atoi(cmd)
			if err != nil {
				log.Fatal("Invalid input: ", err)
			}

			if intCmd == 0 {
				ClearCli()
				ResetDataBase(db)
			}
			ClearCli()
			RemoveTask(db, uint(intCmd))
		case "5":
			ClearCli()
			return
		}
	}

}

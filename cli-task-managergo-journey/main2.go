//
// import (
// 	"bufio"
// 	"database/sql"
// 	"fmt"
// 	"log"
// 	"os"
// 	"path/filepath"
// 	"strconv"
//
// 	_ "github.com/mattn/go-sqlite3"
// )
//
// func main() {
//
// 	args := os.Args
//
// 	homeDir, err := os.UserHomeDir()
// 	if err != nil {
// 		log.Fatal(err)
// 	}
//
// 	dbPath := filepath.Join(homeDir, "tasks.db")
// 	db, err := sql.Open("sqlite3", dbPath)
//
// 	if err != nil {
// 		log.Fatal(err)
// 	}
//
// 	defer db.Close()
//
// 	// testing db
//
// 	err = db.Ping()
// 	if err != nil {
// 		log.Fatal(err)
// 	}
//
// 	CreateTaskDb(db)
//
// 	if len(args) > 1 {
// 		switch args[1] {
// 		case "--list", "-l":
// 			ListTasks(db)
// 			return
// 		case "-a", "--add":
// 			InsertTask(db, args[2])
// 			return
// 		case "-d", "--delete":
// 			id, err := strconv.Atoi(args[2])
// 			if err != nil {
// 				log.Fatal(err)
// 			}
// 			if id == 0 {
// 				ResetDataBase(db)
// 			}
// 			RemoveTask(db, uint(id))
// 			return
// 		case "-h", "--help":
//
// 			fmt.Println("Usage : tasks [option] [arguments]")
// 			fmt.Println("Options:")
// 			fmt.Println("-l, --list\t\tList all tasks")
// 			fmt.Println("-a, --add [task]\tAdd a new task. Provide the task description")
// 			fmt.Println("-d, --delete [id]\tDelete a task by its ID. Use 0 to reset the database")
// 			fmt.Println("-h, --help\t\tDisplay this help message")
//
// 			return
// 		default:
// 			log.Fatal("Invalid arguments run tasks -h or --help for more informations")
// 		}
// 	}
// 	ListTasks(db)
// 	reader := bufio.NewScanner(os.Stdin)
//
// 	for {
// 		fmt.Println("\n ")
// 		fmt.Println("\nChoose Option : ")
// 		fmt.Println("1: Add task")
// 		fmt.Println("2: List task")
// 		fmt.Println("3: Update task")
// 		fmt.Println("4: Delete task")
// 		fmt.Println("5: Exit")
//
// 		reader.Scan()
// 		option := reader.Text()
//
// 		switch option {
// 		case "1":
// 			ClearCli()
// 			fmt.Println("\nEnter the task description")
// 			reader.Scan()
// 			description := reader.Text()
// 			ClearCli()
// 			InsertTask(db, description)
// 		case "2":
// 			ClearCli()
// 			ListTasks(db)
// 		case "3":
// 			ClearCli()
// 			fmt.Println("\nWhat do you want to update: ")
// 			fmt.Println("1: Modify task")
// 			fmt.Println("2: Mark as Complete/Incomplete")
// 			reader.Scan()
// 			option := reader.Text()
// 			if option == "1" {
// 				ClearCli()
// 				ListTasks(db)
// 				fmt.Println("\nPlease choose the id of task you want to modify: ")
// 				reader.Scan()
// 				id := reader.Text()
// 				intId, err := strconv.Atoi(id)
// 				if err != nil {
// 					log.Fatal("Invalid input: ", err)
// 				}
// 				ClearCli()
// 				fmt.Printf("\nYou choose id number %d", intId)
// 				fmt.Println("\nWhat's the new description ? ")
// 				reader.Scan()
// 				newDesc := reader.Text()
// 				ClearCli()
// 				ModifyTask(db, newDesc, uint(intId))
//
// 			} else if option == "2" {
// 				ClearCli()
// 				ListTasks(db)
// 				fmt.Println("What id of the task you want to mark as complete/incomplete ?")
// 				reader.Scan()
// 				id := reader.Text()
// 				intId, err := strconv.Atoi(id)
// 				if err != nil {
// 					log.Fatal("Invalid input: ", err)
// 				}
// 				ClearCli()
// 				MarkAsDoneUnDone(db, uint(intId))
// 			} else {
// 				fmt.Println("invalid option")
// 				return
// 			}
//
// 		case "4":
// 			ClearCli()
// 			ListTasks(db)
// 			fmt.Println("\nEnter the id of the task or type 0 if you want to remove them all: ")
// 			reader.Scan()
//
// 			cmd := reader.Text()
// 			intCmd, err := strconv.Atoi(cmd)
// 			if err != nil {
// 				log.Fatal("Invalid input: ", err)
// 			}
//
// 			if intCmd == 0 {
// 				ClearCli()
// 				ResetDataBase(db)
// 			}
// 			ClearCli()
// 			RemoveTask(db, uint(intCmd))
// 		case "5":
// 			ClearCli()
// 			return
// 		}
// 	}
//
// }
//
// func CreateTaskDb(db *sql.DB) {
// 	query := `
// 	CREATE TABLE IF NOT EXISTS tasks (
// 	id INTEGER PRIMARY KEY AUTOINCREMENT,
// 	description TEXT,
// 	status CHAR(1) DEFAULT '✗' CHECK (status IN ('✗','✓')),
// 	date TIMESTAMP DEFAULT CURRENT_TIMESTAMP
// 	);`
//
// 	_, err := db.Exec(query)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// }
//
// func InsertTask(db *sql.DB, desc string) {
// 	query := `INSERT INTO tasks (description,status) VALUES (?,?)`
// 	_, err := db.Exec(query, desc, "✗")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Printf("\nInserted task: %s\n", desc)
// }
//
// func ResetDataBase(db *sql.DB) {
// 	query := "DELETE FROM tasks;"
// 	_, err := db.Exec(query)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Println("All tasks were deleted successfully")
// }
//
// func RemoveTask(db *sql.DB, id uint) {
// 	query := `DELETE FROM tasks WHERE id = ?`
//
// 	result, err := db.Exec(query, id)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
//
// 	rows, err := result.RowsAffected()
// 	if err != nil {
// 		log.Fatal(err)
// 	}
//
// 	if rows == 0 {
// 		fmt.Printf("Task with id %d not found", id)
// 	} else {
// 		fmt.Printf("Task with id: %d removed successfully", id)
// 	}
// }
//
// func ListTasks(db *sql.DB) {
// 	const (
// 		Reset     = "\033[0m"
// 		Red       = "\033[31m"
// 		Green     = "\033[32m"
// 		Yellow    = "\033[33m"
// 		LightBlue = "\033[34;1m"
// 		Cyan      = "\033[36m"
// 	)
//
// 	query := `SELECT * FROM tasks;`
//
// 	rows, err := db.Query(query)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer rows.Close()
//
// 	fmt.Println()
// 	for rows.Next() {
// 		var id uint
// 		var description string
// 		var status string
// 		var date time.Time
//
// 		err = rows.Scan(&id, &description, &status, &date)
// 		if err != nil {
// 			log.Fatal(err)
// 		}
//
// 		// Determine color based on status
// 		var statusColor string
// 		switch status {
// 		case "✓":
// 			statusColor = Green
// 		case "✗":
// 			statusColor = Red
// 		default:
// 			statusColor = Yellow // Fallback color for unknown status
// 		}
//
// 		// Print colored output
// 		fmt.Printf("%sID: %d%s | %sDescription: %s%s | %sStatus: %s%s | %sDate: %s%s\n",
// 			LightBlue, id, Reset,
// 			Yellow, description, Reset,
// 			statusColor, status, Reset,
// 			Cyan, date.Format("2006-01-02 15:04:05"), Reset)
// 	}
//
// 	if err = rows.Err(); err != nil {
// 		log.Fatal(err)
// 	}
// }
//
// func MarkAsDoneUnDone(db *sql.DB, id uint) {
//
// 	query := `
// 	UPDATE tasks
// 	SET status = CASE
// 		WHEN status = '✗' THEN '✓'
// 		ELSE ''
// 	END
// 	WHERE id = ?;
// 	`
// 	result, err := db.Exec(query, id)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
//
// 	rows, err := result.RowsAffected()
// 	if err != nil {
// 		log.Fatal(err)
// 	}
//
// 	if rows == 0 {
// 		fmt.Printf("Task with id %d not found", id)
// 	} else {
// 		fmt.Printf("Task with id: %d modified successfully", id)
// 	}
// }
//
// func ModifyTask(db *sql.DB, newDesc string, id uint) {
// 	query := `
// 	UPDATE tasks
// 	SET description = ?
// 	WHERE id = ?;
// 	`
// 	result, err := db.Exec(query, newDesc, id)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
//
// 	rows, err := result.RowsAffected()
// 	if err != nil {
// 		log.Fatal(err)
// 	}
//
// 	if rows == 0 {
// 		fmt.Printf("Task with id %d not found", id)
// 	} else {
// 		fmt.Printf("Task with id: %d modified successfully", id)
// 	}
//
// }

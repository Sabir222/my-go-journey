# Task CLI

Task CLI is a simple command-line interface application for managing your tasks. It allows you to add, list, update, and delete tasks, as well as mark them as complete or incomplete.

## Features

- Add new tasks
- List all tasks
- Update task descriptions
- Mark tasks as complete or incomplete
- Delete individual tasks
- Reset the entire task database
- Command-line arguments for quick actions
- Interactive menu for more detailed operations

## Installation

1. Ensure you have Go installed on your system.
3. Install the required dependencies:

```
go get github.com/mattn/go-sqlite3
```

4. Build the application:

```
go build -o tasks .
```

5. (Optional) Move the `tasks` executable to a directory in your system PATH for easy access.

## Usage

### Command-line Arguments

You can use the following command-line arguments for quick actions:

- List all tasks: `tasks -l` or `tasks --list`
- Add a new task: `tasks -a "Task description"` or `tasks --add "Task description"`
- Delete a task: `tasks -d <task_id>` or `tasks --delete <task_id>`
  - Use `tasks -d 0` to reset the entire database
- Display help: `tasks -h` or `tasks --help`

### Interactive Menu

Run `tasks` without any arguments to enter the interactive menu:

1. Add task
2. List tasks
3. Update task
4. Delete task
5. Exit

#### Adding a Task

Choose option 1 and enter the task description when prompted.

#### Listing Tasks

Choose option 2 to display all tasks with their IDs, descriptions, statuses, and creation dates.

#### Updating a Task

Choose option 3, then select whether you want to:
1. Modify the task description
2. Mark the task as complete/incomplete

Follow the prompts to select the task by ID and make the desired changes.

#### Deleting a Task

Choose option 4, then enter the ID of the task you want to delete. Enter 0 to delete all tasks.

## Database

Tasks are stored in an SQLite database located at `~/tasks.db`. The database is created automatically when you first run the application.

## Color Coding

When listing tasks, the output is color-coded for better readability:
- Task ID: Light Blue
- Description: Yellow
- Status: Green (complete) or Red (incomplete)
- Date: Cyan

## Contributing

Feel free to fork this repository and submit pull requests for any improvements or bug fixes.

## License

[MIT]

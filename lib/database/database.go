package database

import (
	"database/sql"
	"log"
	"time"

	"github.com/beadex/espresso/lib/models"
	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

// Init initializes the database connection and schema.
func Init() {
	var err error
	db, err = sql.Open("sqlite3", "./espresso.db")
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	log.Println("Database connected successfully.")
	createSchema()
}

// createSchema creates the database schema.
func createSchema() {
	schema := `
	CREATE TABLE IF NOT EXISTS tasks (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		due_date TEXT NOT NULL,
		is_recurring BOOLEAN NOT NULL
	);
	`
	_, err := db.Exec(schema)
	if err != nil {
		log.Fatalf("Failed to create schema: %v", err)
	}
	log.Println("Database schema initialized.")
}

// GetAllTasks retrieves all tasks from the database.
func GetAllTasks() ([]models.Task, error) {
	rows, err := db.Query("SELECT id, name, due_date, is_recurring FROM tasks")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []models.Task
	for rows.Next() {
		var task models.Task
		var dueDateStr string
		if err := rows.Scan(&task.ID, &task.Name, &dueDateStr, &task.IsRecurring); err != nil {
			return nil, err
		}

		dueDate, err := time.Parse(time.RFC3339, dueDateStr)
		if err != nil {
			return nil, err
		}
		task.DueDate = dueDate
		tasks = append(tasks, task)
	}
	return tasks, nil
}

// InsertTask inserts a new task into the database.
func InsertTask(task models.Task) (int, error) {
	stmt, err := db.Prepare("INSERT INTO tasks (name, due_date, is_recurring) VALUES (?, ?, ?)")
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	result, err := stmt.Exec(task.Name, task.DueDate.Format(time.RFC3339), task.IsRecurring)
	if err != nil {
		return 0, err
	}
	id, err := result.LastInsertId()
	return int(id), err
}

// UpdateTask updates an existing task in the database.
func UpdateTask(task models.Task) error {
	stmt, err := db.Prepare("UPDATE tasks SET name = ?, due_date = ?, is_recurring = ? WHERE id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(task.Name, task.DueDate.Format(time.RFC3339), task.IsRecurring, task.ID)
	return err
}

// DeleteTask deletes a task from the database.
func DeleteTask(id int) error {
	stmt, err := db.Prepare("DELETE FROM tasks WHERE id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(id)
	return err
}

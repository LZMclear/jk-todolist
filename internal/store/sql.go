package store

import (
	"database/sql"
	"jk-todolist/internal/model"
	"time"
)

func InitDB(db *sql.DB) error {
	// create table if not exists (MySQL compatible)
	schema := `
	CREATE TABLE IF NOT EXISTS tasks (
		id BIGINT AUTO_INCREMENT PRIMARY KEY,
		title VARCHAR(255) NOT NULL,
		description TEXT,
		category VARCHAR(100) DEFAULT '',
		completed TINYINT(1) NOT NULL DEFAULT 0,
		due_date DATETIME NULL,
		created_at DATETIME NOT NULL,
		updated_at DATETIME NOT NULL
	) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
	`
	if _, err := db.Exec(schema); err != nil {
		return err
	}
	return nil
}

func CreateTask(db *sql.DB, title, description, category string, dueDate *time.Time) (model.Task, error) {
	now := time.Now().UTC()
	// dueDate can be nil
	res, err := db.Exec("INSERT INTO tasks (title, description, category, completed, due_date, created_at, updated_at) VALUES (?, ?, ?, 0, ?, ?, ?)", title, description, category, dueDate, now, now)
	if err != nil {
		return model.Task{}, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return model.Task{}, err
	}
	return model.Task{ID: id, Title: title, Description: description, Category: category, Completed: false, DueDate: dueDate, CreatedAt: now, UpdatedAt: now}, nil
}

func ListTasks(db *sql.DB) ([]model.Task, error) {
	rows, err := db.Query("SELECT id, title, description, category, completed, due_date, created_at, updated_at FROM tasks ORDER BY id DESC")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []model.Task
	for rows.Next() {
		var t model.Task
		var completedInt int
		var due sql.NullTime
		if err := rows.Scan(&t.ID, &t.Title, &t.Description, &t.Category, &completedInt, &due, &t.CreatedAt, &t.UpdatedAt); err != nil {
			return nil, err
		}
		t.Completed = completedInt != 0
		if due.Valid {
			d := due.Time.UTC()
			t.DueDate = &d
		} else {
			t.DueDate = nil
		}
		tasks = append(tasks, t)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return tasks, nil
}

func GetTask(db *sql.DB, id int64) (model.Task, error) {
	var t model.Task
	var completedInt int
	var due sql.NullTime
	err := db.QueryRow("SELECT id, title, description, category, completed, due_date, created_at, updated_at FROM tasks WHERE id = ?", id).
		Scan(&t.ID, &t.Title, &t.Description, &t.Category, &completedInt, &due, &t.CreatedAt, &t.UpdatedAt)
	if err != nil {
		return model.Task{}, err
	}
	t.Completed = completedInt != 0
	if due.Valid {
		d := due.Time.UTC()
		t.DueDate = &d
	} else {
		t.DueDate = nil
	}
	return t, nil
}

func UpdateTask(db *sql.DB, id int64, title, description, category string, dueDate *time.Time, completed bool) (model.Task, error) {
	now := time.Now().UTC()
	completedInt := 0
	if completed {
		completedInt = 1
	}
	_, err := db.Exec("UPDATE tasks SET title = ?, description = ?, category = ?, due_date = ?, completed = ?, updated_at = ? WHERE id = ?", title, description, category, dueDate, completedInt, now, id)
	if err != nil {
		return model.Task{}, err
	}
	return GetTask(db, id)
}

func DeleteTask(db *sql.DB, id int64) error {
	_, err := db.Exec("DELETE FROM tasks WHERE id = ?", id)
	return err
}

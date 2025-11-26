package model

import "time"

type Task struct {
	ID          int64      `db:"id" json:"id"`
	Title       string     `db:"title" json:"title"`
	Description string     `db:"description" json:"description"`
	Category    string     `db:"category" json:"category"`
	Completed   bool       `db:"completed" json:"completed"`
	DueDate     *time.Time `db:"due_date" json:"due_date,omitempty"`
	CreatedAt   time.Time  `db:"created_at" json:"created_at"`
	UpdatedAt   time.Time  `db:"updated_at" json:"updated_at"`
}
